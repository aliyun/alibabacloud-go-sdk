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
	// E062C482-1A4B-469E-938C-96D28CFAE42E
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
