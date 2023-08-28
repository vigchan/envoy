package main

import (
	"fmt"
        "math/rand"
	"strconv"

	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
)

var UpdateUpstreamBody = "upstream response body updated by the simple plugin"

type filter struct {
	api.PassThroughStreamFilter

	callbacks api.FilterCallbackHandler
	path      string
	config    *config
}

func (f *filter) sendLocalReplyInternal() api.StatusType {
	
	// return api.Continue
        body := fmt.Sprintf("%s, path: %s\r\n", f.config.echoBody, f.path)
	f.callbacks.SendLocalReply(200, body, nil, 0, "")
	return api.LocalReply
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

// Callbacks which are called in request path
func (f *filter) DecodeHeaders(header api.RequestHeaderMap, endStream bool) api.StatusType {
	header.Set("user-agent-1", "test-sample-response")
	f.path, _ = header.Get(":path")
	//if f.path == "/localreply_by_config" {
	//	return f.sendLocalReplyInternal()
	//}
        for i := 1; i < 90; i++ {
	   header.Set("user-agent-"+strconv.Itoa(i), "test-sample-response-benchmark"+strconv.Itoa(i))
        }
        
	header.Set("user-agent-101", "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLSjFbcXoEFfRsWxPLDnJObCsNVlgTeMaPEZQleQYhYzRyWJjPjzpfRFEgmotaFetHsbZRjxAwnwekrBEmfdzdcEkXBAkjQZLCtTMtTCoaNatyyiNKAReKJyiXJrscctNswYNsGRussVmaozFZBsbOJiFQGZsnwTKSmVoiGLOpbUOpEdKupdOMeRVjaRzLNTXYeUCWKsXbGyRAOmBTvKSJfjzaLbtZsyMGeuDtRzQMDQiYCOhgHOvgSeycJPJHYNufNjJhhjUVRuSqfgqVMkPYVkURUpiFvIZRgBmyArKCtzkjkZIvaBjMkXVbWGvbqzgexyALBsdjSGpngCwFkDifIBuufFMoWdiTskZoQJMqrTICTojIYxyeSxZyfroRODMbNDRZnPNRWCJPMHDtJmHAYORsUfUMApsVgzHblmYYtEjVgwfFbbGGcnqbaEREunUZjQXmZOtaRLUtmYgmSVYBADDvoxIfsfgPyCKmxIubeYTNDtjAyRRDedMiyLprucjiOgjhYeVwBTCMLfrDGXqwpzwVGqMZcLVCxaSJlDSYEofkkEYeqkKHqgBpnbPbgHMLUIDjUMmpBHCSjMJjxzuaiIsNBakqSwQpOQgNczgaczAInLqLIbAatLYHdaopovFOkqIexsFzXzrlcztxcdJJFuyZHRCovgpVvlGsXalGqARmneBZBFelhXkzzfNaVtAyyqWzKqQFbucqNJYWRncGKKLdTkNyoCSfkFohsVVxSAZWEXejhAquXdaaaZlRHoNXvpayoSsqcnCTuGZamCToZvPynaEphIdXaKUaqmBdtZtcOfFSPqKXSLEfZAPaJzldaUEdhITGHvBrQPqWARPXPtPVGNpdGERwVhGCMdfLitTqwLUecgOczXTbRMGxqPexOUAbUdQrIPjyQyQFStFubVVdHtAknjEQxCqkDIfTGXeJtuncbfqQUsXTOdPORvAUkAwww"+strconv.Itoa(101))

	header.Set("user-agent-102", "TndUJHiQecbxzvqzlPWyqOsUbLnAIPRYxdkhCBcgRVdrSfOWLQrCNPANWvkkOdejVcugMVTnMUqBcVFohTtUTuWJRGiQjOTcVwHjKVtPoVVTYNXfTjPcDeMixVduzfhinMkGuotVuZJUuZkJwQsrzoVWaYedKZpEKrSOKAaKzKaWlfVcjYMcwVxDrnsfUrCFsVySzWtbpwvkbeRKvUdUyRSgDknQXeAJnkxLkNKWWUEqSfzvxvjTsnzIMKcSiStqDUZsoLEziJiWmAPKvATdjZyEvvmsIDSJeKovPdQSnyNSLTvmwZLbKtziYpTpzlFPhUhSWZdgpyWsRCkctWdcUosnpMVHZdiyeoUqzMJGHOxYTsbfqAqtXjkVItIBhgIvarcSbblmmHEfnwcNjQpkSlxkkUvYxTocDoTiUUJeVakGzAALtLhTUZHJOxKmRNwKDOUnXzmDikdJPukQLaxAHzFSnsZRqGNADbebYsIIwcKGFxUvaLskaPVQqUQnzXSfBigfpkmlDQoLjksSjFLLBLbuEhZFxQuvApLIPkPzXDPizdhpgJsjyBqaYFyqjyZdUShIIWpUYaqBhqHzoCIkjKQbHjIwqPZdSYRMmFeYRrkuEqCuPDMNeuiKTxRINKSzjAIvxBRDRLmciVEWlUwPHdPiHZuUiUFXrgXbHMbjeBvZwcJXyXLoetEAKetWXUDXJDUtrnryhmSlaOSSUVsZsmHKaNgeqzCZeWGRrDuCuWPtQAqMTyVfpvpkHdgWhxbhMSAcmaAvASRDtXBDGnMCDMOfzVWicMsnOkZlIUNSsCraKQBrSfkMAIStKMupnCMvADvezKymxCWrwlUxbdMAyhEcgCTauBGlunxyyOMAmrfxygzWgLcDnnMAWeJSYZBjGdAeQxmvLHqeeizRbUesPXmQumDoPOnnQVNsqKkejfUlxsWAwWjtpZBrWZKvdEANNUPkaNBORKcOQYawyaytedjZtrpZTljvjyNubgyU"+strconv.Itoa(102)) 
        
        header.Set("user-agent-103", "rETWdtOoPuIaFONFwrxmPGDUaMiraBzMlulhxkKvSvnXEUmQuNfIqPjHTcdmcQRyRegnCPCePNnimQoCNrtHHZSHVyIifKqRwXVBMhXNhuqXwvAqhnQXRQSXFJLDKmsJvvoFlpVssIAYkbKrImXeoveCluuuNdNJGMnLVgpITBPZGJZlDkAcDeOtKifMofRksnOKExQBQzIWdJnunSelbYYOzQJYZuYsIIVQniApTAIwFjxvlksQbJajXgvWVWGCniuyxyhXphnvyQQHwnzDRITguovpoGAtItTvRaEDtMbmKEEksgqibDiuAMPEkWUVcnYIYLIFCmbXzmAumYLwoHOHbYhEMHUIaCnkiALLfQEADWUAUeUdAOQmFjxfAzsDjAsjwttyMUOopasuhTESsDemerTKSzUdvZOjUrNyHXgwcGZyYCDAokdrOQioJHWwqbeWPZBPYYlgTvJYNOnKOMIkroYgnQkALVSDTcUfAdtzlLwwZUBUQPrhVthbTtgjFepZBAdbGDyMThoDrRwkOKoYvEcEIkUZpqyyWpLpdgInRQXYUsryZTqryjyDaOtvUvVJhEZGuZITIVyonUOuTKhtdSGddzgtAHxOrdVEDlYNUzOTCyrcQnuMZnigcFJHUgNTnlLJkTgDCopMSpiJBFCPdPhsGJzkaoIJIxIUZrAIchbkWPjwjpZlWqgmJVyXXwFqkmXNMGpFlkYFlhYVpIyJaJBWtOfGtDtzZJfCTtmKiAJuMlZgAzFIbwZIDZXEiOGZCHZMIntqKThKtfIudpvNkMpptmrHQfeakFOeZddIQVFiRudbwUeMgoeHHrLiHQjgGbLjugFFMZMgFPNXVKpyMQyfrRaoxIsmSbkYdrnGTUJfiCkQZkLSuZBPWvvYnvTKcWRNaacuLGQcqNpPxvGOXZaTcxkUZRzJXLkfbeNGEgdnLLmXSYGluJVOpzrcMjBUWVaGbdRmifvVDtAgGbbfyLMreEphgESvQlJl"+strconv.Itoa(103))
	
        //header.Set("user-agent-101", RandStringBytes(1000)+strconv.Itoa(101))
	//header.Set("user-agent-102", RandStringBytes(1000)+strconv.Itoa(102))
	//header.Set("user-agent-103", RandStringBytes(1000)+strconv.Itoa(103))
	return api.Continue
}

/*
The callbacks can be implemented on demand

func (f *filter) DecodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	return api.Continue
}

func (f *filter) DecodeTrailers(trailers api.RequestTrailerMap) api.StatusType {
	return api.Continue
}
*/

func (f *filter) EncodeHeaders(header api.ResponseHeaderMap, endStream bool) api.StatusType {
	//if f.path == "/update_upstream_response" {
	//	header.Set("Content-Length", strconv.Itoa(len(UpdateUpstreamBody)))
	//}
	//header.Set("Rsp-Header-From-Go", "bar-test")
	return api.Continue
}

// Callbacks which are called in response path
func (f *filter) EncodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	/*if f.path == "/update_upstream_response" {
		if endStream {
			buffer.SetString(UpdateUpstreamBody)
		} else {
			// TODO implement buffer->Drain, buffer.SetString means buffer->Drain(buffer.Len())
			buffer.SetString("")
		}
	}*/
	return api.Continue
}

/*
The callbacks can be implemented on demand

func (f *filter) EncodeTrailers(trailers api.ResponseTrailerMap) api.StatusType {
	return api.Continue
}

func (f *filter) OnDestroy(reason api.DestroyReason) {
}
*/
