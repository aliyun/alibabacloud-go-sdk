// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListProjectLabelsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListProjectLabelsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListProjectLabelsResponseBody
	GetRequestId() *string
	SetResult(v []*ListProjectLabelsResponseBodyResult) *ListProjectLabelsResponseBody
	GetResult() []*ListProjectLabelsResponseBodyResult
	SetSuccess(v bool) *ListProjectLabelsResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListProjectLabelsResponseBody
	GetTotal() *int64
}

type ListProjectLabelsResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListProjectLabelsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 30
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListProjectLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectLabelsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListProjectLabelsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListProjectLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectLabelsResponseBody) GetResult() []*ListProjectLabelsResponseBodyResult {
	return s.Result
}

func (s *ListProjectLabelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectLabelsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListProjectLabelsResponseBody) SetErrorCode(v string) *ListProjectLabelsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectLabelsResponseBody) SetErrorMessage(v string) *ListProjectLabelsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListProjectLabelsResponseBody) SetRequestId(v string) *ListProjectLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectLabelsResponseBody) SetResult(v []*ListProjectLabelsResponseBodyResult) *ListProjectLabelsResponseBody {
	s.Result = v
	return s
}

func (s *ListProjectLabelsResponseBody) SetSuccess(v bool) *ListProjectLabelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectLabelsResponseBody) SetTotal(v int64) *ListProjectLabelsResponseBody {
	s.Total = &v
	return s
}

func (s *ListProjectLabelsResponseBody) Validate() error {
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

type ListProjectLabelsResponseBodyResult struct {
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
	// example:
	//
	// 1
	OpenMergeRequestsCount *int64 `json:"openMergeRequestsCount,omitempty" xml:"openMergeRequestsCount,omitempty"`
}

func (s ListProjectLabelsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListProjectLabelsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListProjectLabelsResponseBodyResult) GetColor() *string {
	return s.Color
}

func (s *ListProjectLabelsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListProjectLabelsResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListProjectLabelsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListProjectLabelsResponseBodyResult) GetOpenMergeRequestsCount() *int64 {
	return s.OpenMergeRequestsCount
}

func (s *ListProjectLabelsResponseBodyResult) SetColor(v string) *ListProjectLabelsResponseBodyResult {
	s.Color = &v
	return s
}

func (s *ListProjectLabelsResponseBodyResult) SetDescription(v string) *ListProjectLabelsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListProjectLabelsResponseBodyResult) SetId(v string) *ListProjectLabelsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListProjectLabelsResponseBodyResult) SetName(v string) *ListProjectLabelsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListProjectLabelsResponseBodyResult) SetOpenMergeRequestsCount(v int64) *ListProjectLabelsResponseBodyResult {
	s.OpenMergeRequestsCount = &v
	return s
}

func (s *ListProjectLabelsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
