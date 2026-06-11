// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCmsServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *GetCmsServiceResponseBody
	GetEnabled() *bool
	SetErrorCode(v string) *GetCmsServiceResponseBody
	GetErrorCode() *string
	SetRequestId(v string) *GetCmsServiceResponseBody
	GetRequestId() *string
}

type GetCmsServiceResponseBody struct {
	// Indicates whether the service or product is activated.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// PROM_NOT_OPEN: Prometheus is not activated.SLS_NOT_OPEN: Simple Log Service (SLS), a dependency of Prometheus, is not activated.CMS_SLR_NOT_EXIST: The service-linked role (SLR) for CloudMonitor does not exist.SLS_SLR_NOT_EXIST: The SLR for SLS does not exist.
	//
	// example:
	//
	// PROM_NOT_OPEN
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCmsServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCmsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetCmsServiceResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetCmsServiceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCmsServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCmsServiceResponseBody) SetEnabled(v bool) *GetCmsServiceResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetCmsServiceResponseBody) SetErrorCode(v string) *GetCmsServiceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCmsServiceResponseBody) SetRequestId(v string) *GetCmsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCmsServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
