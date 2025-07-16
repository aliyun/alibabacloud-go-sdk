// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCdnServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenCdnServiceResponseBody
	GetRequestId() *string
}

type OpenCdnServiceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 97C68796-EB7F-4D41-9D5B-12B909D76508
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenCdnServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenCdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCdnServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenCdnServiceResponseBody) SetRequestId(v string) *OpenCdnServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenCdnServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
