// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmInstanceGlobalConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateGtmInstanceGlobalConfigResponseBody
	GetRequestId() *string
}

type UpdateGtmInstanceGlobalConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGtmInstanceGlobalConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmInstanceGlobalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGtmInstanceGlobalConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGtmInstanceGlobalConfigResponseBody) SetRequestId(v string) *UpdateGtmInstanceGlobalConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGtmInstanceGlobalConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
