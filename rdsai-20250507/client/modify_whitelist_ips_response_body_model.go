// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWhitelistIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifyWhitelistIpsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyWhitelistIpsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyWhitelistIpsResponseBody
	GetSuccess() *bool
}

type ModifyWhitelistIpsResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyWhitelistIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWhitelistIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWhitelistIpsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyWhitelistIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWhitelistIpsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyWhitelistIpsResponseBody) SetMessage(v string) *ModifyWhitelistIpsResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyWhitelistIpsResponseBody) SetRequestId(v string) *ModifyWhitelistIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWhitelistIpsResponseBody) SetSuccess(v bool) *ModifyWhitelistIpsResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyWhitelistIpsResponseBody) Validate() error {
	return dara.Validate(s)
}
