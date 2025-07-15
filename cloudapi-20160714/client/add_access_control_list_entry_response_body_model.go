// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAccessControlListEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddAccessControlListEntryResponseBody
	GetRequestId() *string
}

type AddAccessControlListEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEB6EC62-B6C7-5082-A45A-45A204724AC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAccessControlListEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAccessControlListEntryResponseBody) GoString() string {
	return s.String()
}

func (s *AddAccessControlListEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAccessControlListEntryResponseBody) SetRequestId(v string) *AddAccessControlListEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAccessControlListEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
