// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaskingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]*string) *ModifyMaskingRulesResponseBody
	GetData() map[string]*string
	SetMessage(v string) *ModifyMaskingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyMaskingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyMaskingRulesResponseBody
	GetSuccess() *string
}

type ModifyMaskingRulesResponseBody struct {
	Data map[string]*string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 8B1434A1-08A7-3E8C-A237-076A********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMaskingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaskingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMaskingRulesResponseBody) GetData() map[string]*string {
	return s.Data
}

func (s *ModifyMaskingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyMaskingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMaskingRulesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyMaskingRulesResponseBody) SetData(v map[string]*string) *ModifyMaskingRulesResponseBody {
	s.Data = v
	return s
}

func (s *ModifyMaskingRulesResponseBody) SetMessage(v string) *ModifyMaskingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMaskingRulesResponseBody) SetRequestId(v string) *ModifyMaskingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMaskingRulesResponseBody) SetSuccess(v string) *ModifyMaskingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMaskingRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
