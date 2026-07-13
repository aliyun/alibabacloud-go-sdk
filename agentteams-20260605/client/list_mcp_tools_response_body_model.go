// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpToolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMcpToolsResponseBody
	GetCode() *string
	SetItems(v []*ListMcpToolsResponseBodyItems) *ListMcpToolsResponseBody
	GetItems() []*ListMcpToolsResponseBodyItems
	SetMaxResults(v int32) *ListMcpToolsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMcpToolsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMcpToolsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMcpToolsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMcpToolsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListMcpToolsResponseBody
	GetTotalCount() *int32
}

type ListMcpToolsResponseBody struct {
	Code       *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Items      []*ListMcpToolsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults *int32                           `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message    *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *string                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcpToolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpToolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpToolsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMcpToolsResponseBody) GetItems() []*ListMcpToolsResponseBodyItems {
	return s.Items
}

func (s *ListMcpToolsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpToolsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMcpToolsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpToolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpToolsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcpToolsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMcpToolsResponseBody) SetCode(v string) *ListMcpToolsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetItems(v []*ListMcpToolsResponseBodyItems) *ListMcpToolsResponseBody {
	s.Items = v
	return s
}

func (s *ListMcpToolsResponseBody) SetMaxResults(v int32) *ListMcpToolsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetMessage(v string) *ListMcpToolsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetNextToken(v string) *ListMcpToolsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetRequestId(v string) *ListMcpToolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetSuccess(v bool) *ListMcpToolsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetTotalCount(v int32) *ListMcpToolsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMcpToolsResponseBody) Validate() error {
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

type ListMcpToolsResponseBodyItems struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InputSchema *string `json:"InputSchema,omitempty" xml:"InputSchema,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListMcpToolsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMcpToolsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMcpToolsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListMcpToolsResponseBodyItems) GetInputSchema() *string {
	return s.InputSchema
}

func (s *ListMcpToolsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListMcpToolsResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *ListMcpToolsResponseBodyItems) SetDescription(v string) *ListMcpToolsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListMcpToolsResponseBodyItems) SetInputSchema(v string) *ListMcpToolsResponseBodyItems {
	s.InputSchema = &v
	return s
}

func (s *ListMcpToolsResponseBodyItems) SetName(v string) *ListMcpToolsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListMcpToolsResponseBodyItems) SetTitle(v string) *ListMcpToolsResponseBodyItems {
	s.Title = &v
	return s
}

func (s *ListMcpToolsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
