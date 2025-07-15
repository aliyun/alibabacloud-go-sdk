// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiGroupVpcWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApiGroupVpcWhitelistResponseBody
	GetRequestId() *string
}

type ModifyApiGroupVpcWhitelistResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F8B9DC8C-D6E2-5065-BD1F-0401866E7F10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApiGroupVpcWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupVpcWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupVpcWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApiGroupVpcWhitelistResponseBody) SetRequestId(v string) *ModifyApiGroupVpcWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApiGroupVpcWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
