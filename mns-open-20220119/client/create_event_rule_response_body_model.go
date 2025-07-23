// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateEventRuleResponseBody
	GetCode() *int64
	SetData(v string) *CreateEventRuleResponseBody
	GetData() *string
	SetMessage(v string) *CreateEventRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEventRuleResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateEventRuleResponseBody
	GetStatus() *string
	SetSuccess(v bool) *CreateEventRuleResponseBody
	GetSuccess() *bool
}

type CreateEventRuleResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// rule-xsXDW
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEventRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateEventRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateEventRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEventRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEventRuleResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateEventRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateEventRuleResponseBody) SetCode(v int64) *CreateEventRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventRuleResponseBody) SetData(v string) *CreateEventRuleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateEventRuleResponseBody) SetMessage(v string) *CreateEventRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEventRuleResponseBody) SetRequestId(v string) *CreateEventRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventRuleResponseBody) SetStatus(v string) *CreateEventRuleResponseBody {
	s.Status = &v
	return s
}

func (s *CreateEventRuleResponseBody) SetSuccess(v bool) *CreateEventRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateEventRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
