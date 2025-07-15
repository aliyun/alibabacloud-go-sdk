// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApiConfigurationResponseBody
	GetRequestId() *string
}

type ModifyApiConfigurationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6C87A26A-6A18-4B8E-8099-705278381A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApiConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApiConfigurationResponseBody) SetRequestId(v string) *ModifyApiConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApiConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
