// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifyAccountDescriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAccountDescriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyAccountDescriptionResponseBody
	GetSuccess() *bool
}

type ModifyAccountDescriptionResponseBody struct {
	// example:
	//
	// msg
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAccountDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountDescriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAccountDescriptionResponseBody) SetMessage(v string) *ModifyAccountDescriptionResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountDescriptionResponseBody) SetSuccess(v bool) *ModifyAccountDescriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAccountDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
