// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteEventRuleResponseBody
	GetCode() *int64
	SetMessage(v string) *DeleteEventRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEventRuleResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteEventRuleResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DeleteEventRuleResponseBody
	GetSuccess() *bool
}

type DeleteEventRuleResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
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

func (s DeleteEventRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteEventRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEventRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventRuleResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteEventRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEventRuleResponseBody) SetCode(v int64) *DeleteEventRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventRuleResponseBody) SetMessage(v string) *DeleteEventRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventRuleResponseBody) SetRequestId(v string) *DeleteEventRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventRuleResponseBody) SetStatus(v string) *DeleteEventRuleResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteEventRuleResponseBody) SetSuccess(v bool) *DeleteEventRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEventRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
