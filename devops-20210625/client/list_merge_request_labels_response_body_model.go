// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListMergeRequestLabelsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListMergeRequestLabelsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListMergeRequestLabelsResponseBody
	GetRequestId() *string
	SetResult(v []*ListMergeRequestLabelsResponseBodyResult) *ListMergeRequestLabelsResponseBody
	GetResult() []*ListMergeRequestLabelsResponseBodyResult
	SetSuccess(v bool) *ListMergeRequestLabelsResponseBody
	GetSuccess() *bool
}

type ListMergeRequestLabelsResponseBody struct {
	// example:
	//
	// Invalid.IdNotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 4D6AF7CC-B43B-5454-86AB-023D25E44868
	RequestId *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListMergeRequestLabelsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListMergeRequestLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMergeRequestLabelsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMergeRequestLabelsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMergeRequestLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMergeRequestLabelsResponseBody) GetResult() []*ListMergeRequestLabelsResponseBodyResult {
	return s.Result
}

func (s *ListMergeRequestLabelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMergeRequestLabelsResponseBody) SetErrorCode(v string) *ListMergeRequestLabelsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMergeRequestLabelsResponseBody) SetErrorMessage(v string) *ListMergeRequestLabelsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMergeRequestLabelsResponseBody) SetRequestId(v string) *ListMergeRequestLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMergeRequestLabelsResponseBody) SetResult(v []*ListMergeRequestLabelsResponseBodyResult) *ListMergeRequestLabelsResponseBody {
	s.Result = v
	return s
}

func (s *ListMergeRequestLabelsResponseBody) SetSuccess(v bool) *ListMergeRequestLabelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMergeRequestLabelsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMergeRequestLabelsResponseBodyResult struct {
	// example:
	//
	// #A16AD7
	Color       *string `json:"color,omitempty" xml:"color,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// f2bf0e0b4ce34a348b2d971c69a1d11f
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListMergeRequestLabelsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestLabelsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMergeRequestLabelsResponseBodyResult) GetColor() *string {
	return s.Color
}

func (s *ListMergeRequestLabelsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListMergeRequestLabelsResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListMergeRequestLabelsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListMergeRequestLabelsResponseBodyResult) SetColor(v string) *ListMergeRequestLabelsResponseBodyResult {
	s.Color = &v
	return s
}

func (s *ListMergeRequestLabelsResponseBodyResult) SetDescription(v string) *ListMergeRequestLabelsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListMergeRequestLabelsResponseBodyResult) SetId(v string) *ListMergeRequestLabelsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListMergeRequestLabelsResponseBodyResult) SetName(v string) *ListMergeRequestLabelsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListMergeRequestLabelsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
