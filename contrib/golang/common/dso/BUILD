load(
    "//bazel:envoy_build_system.bzl",
    "envoy_cc_library",
    "envoy_extension_package",
)

licenses(["notice"])  # Apache 2

envoy_extension_package()

envoy_cc_library(
    name = "dso_lib",
    srcs = ["dso.cc"],
    hdrs = [
        "api.h",
        "dso.h",
        "libgolang.h",
    ],
    deps = [
        "//source/common/common:minimal_logger_lib",
        "//source/common/common:utility_lib",
    ],
)
