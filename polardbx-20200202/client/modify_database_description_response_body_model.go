// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifyDatabaseDescriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyDatabaseDescriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDatabaseDescriptionResponseBody
	GetSuccess() *bool
}

type ModifyDatabaseDescriptionResponseBody struct {
	// example:
	//
	// success
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

func (s ModifyDatabaseDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyDatabaseDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDatabaseDescriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDatabaseDescriptionResponseBody) SetMessage(v string) *ModifyDatabaseDescriptionResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDatabaseDescriptionResponseBody) SetRequestId(v string) *ModifyDatabaseDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabaseDescriptionResponseBody) SetSuccess(v bool) *ModifyDatabaseDescriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDatabaseDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
