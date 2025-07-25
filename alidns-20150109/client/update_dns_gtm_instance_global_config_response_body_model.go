// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmInstanceGlobalConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDnsGtmInstanceGlobalConfigResponseBody
	GetRequestId() *string
}

type UpdateDnsGtmInstanceGlobalConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDnsGtmInstanceGlobalConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmInstanceGlobalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponseBody) SetRequestId(v string) *UpdateDnsGtmInstanceGlobalConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
