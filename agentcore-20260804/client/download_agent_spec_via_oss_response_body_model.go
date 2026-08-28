// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadAgentSpecViaOssResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DownloadAgentSpecViaOssResponseBody
	GetData() *string
	SetRequestId(v string) *DownloadAgentSpecViaOssResponseBody
	GetRequestId() *string
}

type DownloadAgentSpecViaOssResponseBody struct {
	// The response data.
	//
	// example:
	//
	// https://example.com/artifacts/example.zip
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DownloadAgentSpecViaOssResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadAgentSpecViaOssResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadAgentSpecViaOssResponseBody) GetData() *string {
	return s.Data
}

func (s *DownloadAgentSpecViaOssResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadAgentSpecViaOssResponseBody) SetData(v string) *DownloadAgentSpecViaOssResponseBody {
	s.Data = &v
	return s
}

func (s *DownloadAgentSpecViaOssResponseBody) SetRequestId(v string) *DownloadAgentSpecViaOssResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadAgentSpecViaOssResponseBody) Validate() error {
	return dara.Validate(s)
}
