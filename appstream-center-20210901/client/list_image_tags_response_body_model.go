// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListImageTagsResponseBody
	GetCode() *string
	SetData(v []*ListImageTagsResponseBodyData) *ListImageTagsResponseBody
	GetData() []*ListImageTagsResponseBodyData
	SetMaxResults(v int32) *ListImageTagsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListImageTagsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListImageTagsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListImageTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListImageTagsResponseBody
	GetSuccess() *bool
}

type ListImageTagsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidAppInstanceGroup.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of built-in image tags.
	Data []*ListImageTagsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The error message returned when an error occurs.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// ptnJAAAAAAAxNzE5OTEwNQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListImageTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageTagsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListImageTagsResponseBody) GetData() []*ListImageTagsResponseBodyData {
	return s.Data
}

func (s *ListImageTagsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListImageTagsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListImageTagsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImageTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListImageTagsResponseBody) SetCode(v string) *ListImageTagsResponseBody {
	s.Code = &v
	return s
}

func (s *ListImageTagsResponseBody) SetData(v []*ListImageTagsResponseBodyData) *ListImageTagsResponseBody {
	s.Data = v
	return s
}

func (s *ListImageTagsResponseBody) SetMaxResults(v int32) *ListImageTagsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListImageTagsResponseBody) SetMessage(v string) *ListImageTagsResponseBody {
	s.Message = &v
	return s
}

func (s *ListImageTagsResponseBody) SetNextToken(v string) *ListImageTagsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListImageTagsResponseBody) SetRequestId(v string) *ListImageTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageTagsResponseBody) SetSuccess(v bool) *ListImageTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListImageTagsResponseBody) Validate() error {
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

type ListImageTagsResponseBodyData struct {
	// The label categorization.
	//
	// example:
	//
	// Solution
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The tag value.
	//
	// example:
	//
	// AIGC
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListImageTagsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListImageTagsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListImageTagsResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *ListImageTagsResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *ListImageTagsResponseBodyData) SetCategory(v string) *ListImageTagsResponseBodyData {
	s.Category = &v
	return s
}

func (s *ListImageTagsResponseBodyData) SetTag(v string) *ListImageTagsResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ListImageTagsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
