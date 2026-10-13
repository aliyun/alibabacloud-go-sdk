// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDevelopersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDevelopersResponseBody
	GetCode() *string
	SetData(v *ListDevelopersResponseBodyData) *ListDevelopersResponseBody
	GetData() *ListDevelopersResponseBodyData
	SetMaxResults(v int32) *ListDevelopersResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListDevelopersResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListDevelopersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDevelopersResponseBody
	GetRequestId() *string
}

type ListDevelopersResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListDevelopersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The maximum number of entries on the current page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The cursor for the next page. An empty value indicates that the last page has been reached.
	//
	// example:
	//
	// token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListDevelopersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDevelopersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevelopersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDevelopersResponseBody) GetData() *ListDevelopersResponseBodyData {
	return s.Data
}

func (s *ListDevelopersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDevelopersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDevelopersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDevelopersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDevelopersResponseBody) SetCode(v string) *ListDevelopersResponseBody {
	s.Code = &v
	return s
}

func (s *ListDevelopersResponseBody) SetData(v *ListDevelopersResponseBodyData) *ListDevelopersResponseBody {
	s.Data = v
	return s
}

func (s *ListDevelopersResponseBody) SetMaxResults(v int32) *ListDevelopersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDevelopersResponseBody) SetMessage(v string) *ListDevelopersResponseBody {
	s.Message = &v
	return s
}

func (s *ListDevelopersResponseBody) SetNextToken(v string) *ListDevelopersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDevelopersResponseBody) SetRequestId(v string) *ListDevelopersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevelopersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDevelopersResponseBodyData struct {
	// The list of developers.
	Content []*ListDevelopersResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDevelopersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDevelopersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDevelopersResponseBodyData) GetContent() []*ListDevelopersResponseBodyDataContent {
	return s.Content
}

func (s *ListDevelopersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDevelopersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDevelopersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDevelopersResponseBodyData) SetContent(v []*ListDevelopersResponseBodyDataContent) *ListDevelopersResponseBodyData {
	s.Content = v
	return s
}

func (s *ListDevelopersResponseBodyData) SetPageNumber(v int32) *ListDevelopersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDevelopersResponseBodyData) SetPageSize(v int32) *ListDevelopersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDevelopersResponseBodyData) SetTotalCount(v int32) *ListDevelopersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDevelopersResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDevelopersResponseBodyDataContent struct {
	// The developer ID. You can use this ID to locate the developer in API operations such as SetHiMarketDeveloperStatus and DeleteHiMarketDeveloper.
	//
	// example:
	//
	// dev-xxxxxxxx
	DeveloperId *string `json:"developerId,omitempty" xml:"developerId,omitempty"`
	// The ID of the portal to which the developer belongs. You can use this ID to locate the portal in API operations such as GetPortal, UpdatePortal, and DeletePortal.
	//
	// example:
	//
	// portal-xxxxxxxx
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
	// The status of the developer. Valid values: PENDING, APPROVED, and REJECTED.
	//
	// example:
	//
	// APPROVED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The logon username of the developer.
	//
	// example:
	//
	// developer-demo
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListDevelopersResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListDevelopersResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListDevelopersResponseBodyDataContent) GetDeveloperId() *string {
	return s.DeveloperId
}

func (s *ListDevelopersResponseBodyDataContent) GetPortalId() *string {
	return s.PortalId
}

func (s *ListDevelopersResponseBodyDataContent) GetStatus() *string {
	return s.Status
}

func (s *ListDevelopersResponseBodyDataContent) GetUsername() *string {
	return s.Username
}

func (s *ListDevelopersResponseBodyDataContent) SetDeveloperId(v string) *ListDevelopersResponseBodyDataContent {
	s.DeveloperId = &v
	return s
}

func (s *ListDevelopersResponseBodyDataContent) SetPortalId(v string) *ListDevelopersResponseBodyDataContent {
	s.PortalId = &v
	return s
}

func (s *ListDevelopersResponseBodyDataContent) SetStatus(v string) *ListDevelopersResponseBodyDataContent {
	s.Status = &v
	return s
}

func (s *ListDevelopersResponseBodyDataContent) SetUsername(v string) *ListDevelopersResponseBodyDataContent {
	s.Username = &v
	return s
}

func (s *ListDevelopersResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
