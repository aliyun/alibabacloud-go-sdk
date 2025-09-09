// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddServiceSharedAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddServiceSharedAccountsResponseBody
	GetRequestId() *string
}

type AddServiceSharedAccountsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E2815213-EA4F-5759-8EA1-56DD051BB3FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddServiceSharedAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddServiceSharedAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddServiceSharedAccountsResponseBody) SetRequestId(v string) *AddServiceSharedAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddServiceSharedAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}
