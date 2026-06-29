// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListTemplatesResponseBody
	GetCode() *int32
	SetDetails(v string) *ListTemplatesResponseBody
	GetDetails() *string
	SetErrorCode(v string) *ListTemplatesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListTemplatesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTemplatesResponseBody
	GetSuccess() *bool
	SetTemplates(v []*SimpleTemplate) *ListTemplatesResponseBody
	GetTemplates() []*SimpleTemplate
	SetTotalCount(v int32) *ListTemplatesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListTemplatesResponseBody
	GetTotalPage() *int32
}

type ListTemplatesResponseBody struct {
	// Return code. The default value is 0, indicating Normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Return message of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// is succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Template list
	Templates []*SimpleTemplate `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// Total count
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListTemplatesResponseBody) GetDetails() *string {
	return s.Details
}

func (s *ListTemplatesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTemplatesResponseBody) GetTemplates() []*SimpleTemplate {
	return s.Templates
}

func (s *ListTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTemplatesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListTemplatesResponseBody) SetCode(v int32) *ListTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTemplatesResponseBody) SetDetails(v string) *ListTemplatesResponseBody {
	s.Details = &v
	return s
}

func (s *ListTemplatesResponseBody) SetErrorCode(v string) *ListTemplatesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTemplatesResponseBody) SetMessage(v string) *ListTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTemplatesResponseBody) SetPageNumber(v int32) *ListTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesResponseBody) SetPageSize(v int32) *ListTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetSuccess(v bool) *ListTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*SimpleTemplate) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalPage(v int32) *ListTemplatesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
