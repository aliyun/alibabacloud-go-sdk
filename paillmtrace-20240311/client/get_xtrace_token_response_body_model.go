// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetXtraceTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetXtraceTokenResponseBody
	GetCode() *string
	SetGrpcEndpoint(v string) *GetXtraceTokenResponseBody
	GetGrpcEndpoint() *string
	SetGrpcInternalEndpoint(v string) *GetXtraceTokenResponseBody
	GetGrpcInternalEndpoint() *string
	SetHttpEndpoint(v string) *GetXtraceTokenResponseBody
	GetHttpEndpoint() *string
	SetHttpInternalEndpoint(v string) *GetXtraceTokenResponseBody
	GetHttpInternalEndpoint() *string
	SetMessage(v string) *GetXtraceTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetXtraceTokenResponseBody
	GetRequestId() *string
	SetToken(v string) *GetXtraceTokenResponseBody
	GetToken() *string
}

type GetXtraceTokenResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The gRPC endpoint used for uploading ARM traces.
	//
	// example:
	//
	// http://tracing-analysis-dc-hz.aliyuncs.com:8090
	GrpcEndpoint *string `json:"GrpcEndpoint,omitempty" xml:"GrpcEndpoint,omitempty"`
	// The internal gRPC endpoint used for uploading ARMS traces used by Alibaba Cloud.
	//
	// example:
	//
	// http://tracing-analysis-dc-hz-internal.aliyuncs.com:8090
	GrpcInternalEndpoint *string `json:"GrpcInternalEndpoint,omitempty" xml:"GrpcInternalEndpoint,omitempty"`
	// The endpoint used for uploading ARMS traces.
	//
	// example:
	//
	// http://tracing-analysis-dc-hz.aliyuncs.com/aaa@bbb@ccc/api/otlp/traces
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// The internal endpoint used for uploading ARMS traces used by Alibaba Cloud.
	//
	// example:
	//
	// http://tracing-analysis-dc-hz-internal.aliyuncs.com/aaa@bbb@ccc/api/otlp/traces
	HttpInternalEndpoint *string `json:"HttpInternalEndpoint,omitempty" xml:"HttpInternalEndpoint,omitempty"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// get_xtrace_token: failed, ERROR: NoPermission
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The token used for uploading ARMS traces.
	//
	// example:
	//
	// h1abcw7@5abcb_h1abcw7@5abc01
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetXtraceTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetXtraceTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetXtraceTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetXtraceTokenResponseBody) GetGrpcEndpoint() *string {
	return s.GrpcEndpoint
}

func (s *GetXtraceTokenResponseBody) GetGrpcInternalEndpoint() *string {
	return s.GrpcInternalEndpoint
}

func (s *GetXtraceTokenResponseBody) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *GetXtraceTokenResponseBody) GetHttpInternalEndpoint() *string {
	return s.HttpInternalEndpoint
}

func (s *GetXtraceTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetXtraceTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetXtraceTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetXtraceTokenResponseBody) SetCode(v string) *GetXtraceTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetGrpcEndpoint(v string) *GetXtraceTokenResponseBody {
	s.GrpcEndpoint = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetGrpcInternalEndpoint(v string) *GetXtraceTokenResponseBody {
	s.GrpcInternalEndpoint = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetHttpEndpoint(v string) *GetXtraceTokenResponseBody {
	s.HttpEndpoint = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetHttpInternalEndpoint(v string) *GetXtraceTokenResponseBody {
	s.HttpInternalEndpoint = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetMessage(v string) *GetXtraceTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetRequestId(v string) *GetXtraceTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetToken(v string) *GetXtraceTokenResponseBody {
	s.Token = &v
	return s
}

func (s *GetXtraceTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
