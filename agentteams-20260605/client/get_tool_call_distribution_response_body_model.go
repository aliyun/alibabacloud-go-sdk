// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetToolCallDistributionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetToolCallDistributionResponseBody
	GetCode() *string
	SetData(v *GetToolCallDistributionResponseBodyData) *GetToolCallDistributionResponseBody
	GetData() *GetToolCallDistributionResponseBodyData
	SetHttpStatusCode(v int32) *GetToolCallDistributionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetToolCallDistributionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetToolCallDistributionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetToolCallDistributionResponseBody
	GetSuccess() *bool
}

type GetToolCallDistributionResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetToolCallDistributionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetToolCallDistributionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetToolCallDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetToolCallDistributionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetToolCallDistributionResponseBody) GetData() *GetToolCallDistributionResponseBodyData {
	return s.Data
}

func (s *GetToolCallDistributionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetToolCallDistributionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetToolCallDistributionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetToolCallDistributionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetToolCallDistributionResponseBody) SetCode(v string) *GetToolCallDistributionResponseBody {
	s.Code = &v
	return s
}

func (s *GetToolCallDistributionResponseBody) SetData(v *GetToolCallDistributionResponseBodyData) *GetToolCallDistributionResponseBody {
	s.Data = v
	return s
}

func (s *GetToolCallDistributionResponseBody) SetHttpStatusCode(v int32) *GetToolCallDistributionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetToolCallDistributionResponseBody) SetMessage(v string) *GetToolCallDistributionResponseBody {
	s.Message = &v
	return s
}

func (s *GetToolCallDistributionResponseBody) SetRequestId(v string) *GetToolCallDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetToolCallDistributionResponseBody) SetSuccess(v bool) *GetToolCallDistributionResponseBody {
	s.Success = &v
	return s
}

func (s *GetToolCallDistributionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetToolCallDistributionResponseBodyData struct {
	Items      []*GetToolCallDistributionResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalCalls *int32                                          `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
}

func (s GetToolCallDistributionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetToolCallDistributionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetToolCallDistributionResponseBodyData) GetItems() []*GetToolCallDistributionResponseBodyDataItems {
	return s.Items
}

func (s *GetToolCallDistributionResponseBodyData) GetTotalCalls() *int32 {
	return s.TotalCalls
}

func (s *GetToolCallDistributionResponseBodyData) SetItems(v []*GetToolCallDistributionResponseBodyDataItems) *GetToolCallDistributionResponseBodyData {
	s.Items = v
	return s
}

func (s *GetToolCallDistributionResponseBodyData) SetTotalCalls(v int32) *GetToolCallDistributionResponseBodyData {
	s.TotalCalls = &v
	return s
}

func (s *GetToolCallDistributionResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetToolCallDistributionResponseBodyDataItems struct {
	CallCount *int32  `json:"CallCount,omitempty" xml:"CallCount,omitempty"`
	ToolName  *string `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
}

func (s GetToolCallDistributionResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetToolCallDistributionResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetToolCallDistributionResponseBodyDataItems) GetCallCount() *int32 {
	return s.CallCount
}

func (s *GetToolCallDistributionResponseBodyDataItems) GetToolName() *string {
	return s.ToolName
}

func (s *GetToolCallDistributionResponseBodyDataItems) SetCallCount(v int32) *GetToolCallDistributionResponseBodyDataItems {
	s.CallCount = &v
	return s
}

func (s *GetToolCallDistributionResponseBodyDataItems) SetToolName(v string) *GetToolCallDistributionResponseBodyDataItems {
	s.ToolName = &v
	return s
}

func (s *GetToolCallDistributionResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
