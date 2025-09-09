// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceTwoFactorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceTwoFactorResponseBody
	GetRequestId() *string
}

type ModifyInstanceTwoFactorResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9CE1A352-15E9-5EB4-B589-87A8DEECB20D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceTwoFactorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceTwoFactorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTwoFactorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceTwoFactorResponseBody) SetRequestId(v string) *ModifyInstanceTwoFactorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceTwoFactorResponseBody) Validate() error {
	return dara.Validate(s)
}
