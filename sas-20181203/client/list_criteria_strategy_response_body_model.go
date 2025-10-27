// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCriteriaStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCriteriaStrategyResponseBody
	GetCode() *string
	SetData(v []*ListCriteriaStrategyResponseBodyData) *ListCriteriaStrategyResponseBody
	GetData() []*ListCriteriaStrategyResponseBodyData
	SetMessage(v string) *ListCriteriaStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCriteriaStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCriteriaStrategyResponseBody
	GetSuccess() *bool
}

type ListCriteriaStrategyResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs and names of the rules.
	Data []*ListCriteriaStrategyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 835851E3-AFA2-5EA7-93E9-4FC9BCF3F973
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCriteriaStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCriteriaStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ListCriteriaStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCriteriaStrategyResponseBody) GetData() []*ListCriteriaStrategyResponseBodyData {
	return s.Data
}

func (s *ListCriteriaStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCriteriaStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCriteriaStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCriteriaStrategyResponseBody) SetCode(v string) *ListCriteriaStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *ListCriteriaStrategyResponseBody) SetData(v []*ListCriteriaStrategyResponseBodyData) *ListCriteriaStrategyResponseBody {
	s.Data = v
	return s
}

func (s *ListCriteriaStrategyResponseBody) SetMessage(v string) *ListCriteriaStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *ListCriteriaStrategyResponseBody) SetRequestId(v string) *ListCriteriaStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCriteriaStrategyResponseBody) SetSuccess(v bool) *ListCriteriaStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ListCriteriaStrategyResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCriteriaStrategyResponseBodyData struct {
	// The unique identifier of the rule.
	//
	// example:
	//
	// test
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCriteriaStrategyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCriteriaStrategyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCriteriaStrategyResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListCriteriaStrategyResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *ListCriteriaStrategyResponseBodyData) SetId(v int64) *ListCriteriaStrategyResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCriteriaStrategyResponseBodyData) SetValue(v string) *ListCriteriaStrategyResponseBodyData {
	s.Value = &v
	return s
}

func (s *ListCriteriaStrategyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
