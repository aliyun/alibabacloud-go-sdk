// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaskingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]*string) *CreateMaskingRulesResponseBody
	GetData() map[string]*string
	SetMessage(v string) *CreateMaskingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMaskingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateMaskingRulesResponseBody
	GetSuccess() *string
}

type CreateMaskingRulesResponseBody struct {
	Data map[string]*string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful create
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 792233B1-76B8-5A01-92B4-**********864
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMaskingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMaskingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMaskingRulesResponseBody) GetData() map[string]*string {
	return s.Data
}

func (s *CreateMaskingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMaskingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMaskingRulesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateMaskingRulesResponseBody) SetData(v map[string]*string) *CreateMaskingRulesResponseBody {
	s.Data = v
	return s
}

func (s *CreateMaskingRulesResponseBody) SetMessage(v string) *CreateMaskingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMaskingRulesResponseBody) SetRequestId(v string) *CreateMaskingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMaskingRulesResponseBody) SetSuccess(v string) *CreateMaskingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMaskingRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
