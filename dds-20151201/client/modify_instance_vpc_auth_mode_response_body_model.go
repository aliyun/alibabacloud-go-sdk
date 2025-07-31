// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVpcAuthModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceVpcAuthModeResponseBody
	GetRequestId() *string
}

type ModifyInstanceVpcAuthModeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BA51E9D9-B14A-4542-B6E6-7DE00BECCB8C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceVpcAuthModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVpcAuthModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAuthModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceVpcAuthModeResponseBody) SetRequestId(v string) *ModifyInstanceVpcAuthModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeResponseBody) Validate() error {
	return dara.Validate(s)
}
