// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewServiceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenewServiceInstanceResponseBody
	GetRequestId() *string
}

type RenewServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BB58FE53-ED8F-5D12-9746-CD3A5F463D95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewServiceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewServiceInstanceResponseBody) SetRequestId(v string) *RenewServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewServiceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
