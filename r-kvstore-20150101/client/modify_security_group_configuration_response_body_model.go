// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySecurityGroupConfigurationResponseBody
	GetRequestId() *string
}

type ModifySecurityGroupConfigurationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 147CAC09-E8C6-43F8-9599-982A43D6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityGroupConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityGroupConfigurationResponseBody) SetRequestId(v string) *ModifySecurityGroupConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityGroupConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
