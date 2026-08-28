// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAgentSpecViaOssResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UploadAgentSpecViaOssResponseBody
	GetData() *string
	SetRequestId(v string) *UploadAgentSpecViaOssResponseBody
	GetRequestId() *string
}

type UploadAgentSpecViaOssResponseBody struct {
	// The response data.
	//
	// example:
	//
	// agentspec-1234567890abcdef
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UploadAgentSpecViaOssResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSpecViaOssResponseBody) GoString() string {
	return s.String()
}

func (s *UploadAgentSpecViaOssResponseBody) GetData() *string {
	return s.Data
}

func (s *UploadAgentSpecViaOssResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadAgentSpecViaOssResponseBody) SetData(v string) *UploadAgentSpecViaOssResponseBody {
	s.Data = &v
	return s
}

func (s *UploadAgentSpecViaOssResponseBody) SetRequestId(v string) *UploadAgentSpecViaOssResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadAgentSpecViaOssResponseBody) Validate() error {
	return dara.Validate(s)
}
