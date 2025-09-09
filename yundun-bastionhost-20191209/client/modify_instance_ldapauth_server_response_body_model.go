// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceLDAPAuthServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceLDAPAuthServerResponseBody
	GetRequestId() *string
}

type ModifyInstanceLDAPAuthServerResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8F1085E3-F048-5F34-B650-F145216E4AA4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceLDAPAuthServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceLDAPAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceLDAPAuthServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceLDAPAuthServerResponseBody) SetRequestId(v string) *ModifyInstanceLDAPAuthServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerResponseBody) Validate() error {
	return dara.Validate(s)
}
