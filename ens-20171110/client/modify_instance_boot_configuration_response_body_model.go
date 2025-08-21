// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceBootConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceBootConfigurationResponseBody
	GetRequestId() *string
}

type ModifyInstanceBootConfigurationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6666C5A5-75ED-422E-A022-7121FA18C968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceBootConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceBootConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceBootConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceBootConfigurationResponseBody) SetRequestId(v string) *ModifyInstanceBootConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceBootConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
