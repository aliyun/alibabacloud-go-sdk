// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterprisePptTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEnterprisePptTemplatesResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListEnterprisePptTemplatesResponseBody
	GetCurrent() *int32
	SetData(v []*ListEnterprisePptTemplatesResponseBodyData) *ListEnterprisePptTemplatesResponseBody
	GetData() []*ListEnterprisePptTemplatesResponseBodyData
	SetHttpStatusCode(v int32) *ListEnterprisePptTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListEnterprisePptTemplatesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListEnterprisePptTemplatesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListEnterprisePptTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEnterprisePptTemplatesResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListEnterprisePptTemplatesResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListEnterprisePptTemplatesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListEnterprisePptTemplatesResponseBody
	GetTotalCount() *int32
}

type ListEnterprisePptTemplatesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// An array of enterprise-specific PPT template objects.
	Data []*ListEnterprisePptTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The maximum number of results per page, as specified in the request. Note: This parameter is not yet in effect.
	//
	// example:
	//
	// 4
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token for the next page of results. An empty value indicates that no more data is available. Note: This parameter is not yet in effect.
	//
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJANEQ4AACjMDLgAAADFTNzMyZDMwMzAzMDM4NzA3MjZjN2E2NDYyNzUzODMxMzY3ODM0NmIzNTZkNjc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique request ID. Provide this ID when you contact technical support.
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned on the current page (the \\"page size\\").
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// Indicates whether the request succeeded. `true` indicates success, and `false` indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries that match the query.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEnterprisePptTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnterprisePptTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterprisePptTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEnterprisePptTemplatesResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListEnterprisePptTemplatesResponseBody) GetData() []*ListEnterprisePptTemplatesResponseBodyData {
	return s.Data
}

func (s *ListEnterprisePptTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListEnterprisePptTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEnterprisePptTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnterprisePptTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEnterprisePptTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnterprisePptTemplatesResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListEnterprisePptTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEnterprisePptTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEnterprisePptTemplatesResponseBody) SetCode(v string) *ListEnterprisePptTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBody) SetCurrent(v int32) *ListEnterprisePptTemplatesResponseBody {
	s.Current = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBody) SetData(v []*ListEnterprisePptTemplatesResponseBodyData) *ListEnterprisePptTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBody) SetHttpStatusCode(v int32) *ListEnterprisePptTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBody) SetMaxResults(v int32) *ListEnterprisePptTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBody) SetMessage(v string) *ListEnterprisePptTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBody) SetNextToken(v string) *ListEnterprisePptTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBody) SetRequestId(v string) *ListEnterprisePptTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBody) SetSize(v int32) *ListEnterprisePptTemplatesResponseBody {
	s.Size = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBody) SetSuccess(v bool) *ListEnterprisePptTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBody) SetTotalCount(v int32) *ListEnterprisePptTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBody) Validate() error {
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

type ListEnterprisePptTemplatesResponseBodyData struct {
	// The URL of the cover image.
	//
	// example:
	//
	// http://xxx.com/a.png
	CoverImg *string `json:"CoverImg,omitempty" xml:"CoverImg,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 10
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListEnterprisePptTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnterprisePptTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnterprisePptTemplatesResponseBodyData) GetCoverImg() *string {
	return s.CoverImg
}

func (s *ListEnterprisePptTemplatesResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListEnterprisePptTemplatesResponseBodyData) SetCoverImg(v string) *ListEnterprisePptTemplatesResponseBodyData {
	s.CoverImg = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBodyData) SetId(v int64) *ListEnterprisePptTemplatesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListEnterprisePptTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
