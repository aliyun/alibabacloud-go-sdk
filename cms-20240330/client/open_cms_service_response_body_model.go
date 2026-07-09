// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCmsServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *OpenCmsServiceResponseBody
	GetEnabled() *bool
	SetRequestId(v string) *OpenCmsServiceResponseBody
	GetRequestId() *string
}

type OpenCmsServiceResponseBody struct {
	// Indicates whether the service or commodity is activated.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OpenCmsServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenCmsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCmsServiceResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *OpenCmsServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenCmsServiceResponseBody) SetEnabled(v bool) *OpenCmsServiceResponseBody {
	s.Enabled = &v
	return s
}

func (s *OpenCmsServiceResponseBody) SetRequestId(v string) *OpenCmsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenCmsServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
