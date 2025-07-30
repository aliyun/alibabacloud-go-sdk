// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyGlobalSecurityIPGroupNameResponseBody
	GetRequestId() *string
}

type ModifyGlobalSecurityIPGroupNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2B17D708-1D6D-49F3-B6D7-478371DD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) SetRequestId(v string) *ModifyGlobalSecurityIPGroupNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) Validate() error {
	return dara.Validate(s)
}
