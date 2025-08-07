// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaskingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifyMaskingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyMaskingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyMaskingRulesResponseBody
	GetSuccess() *bool
}

type ModifyMaskingRulesResponseBody struct {
	// The message that is returned for the request.
	//
	// > If the request is successful, Successful is returned. If the request fails, an error message such as an error code is returned.
	//
	// example:
	//
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 99B355CE-526C-478B-B730-AD9D7C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid value:
	//
	// 	- **true**:
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMaskingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaskingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMaskingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyMaskingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMaskingRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMaskingRulesResponseBody) SetMessage(v string) *ModifyMaskingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMaskingRulesResponseBody) SetRequestId(v string) *ModifyMaskingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMaskingRulesResponseBody) SetSuccess(v bool) *ModifyMaskingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMaskingRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
