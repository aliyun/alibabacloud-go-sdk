// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaCatalogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLumaCatalogsResponseBody
	GetCode() *string
	SetData(v *ListLumaCatalogsResponseBodyData) *ListLumaCatalogsResponseBody
	GetData() *ListLumaCatalogsResponseBodyData
	SetMessage(v string) *ListLumaCatalogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLumaCatalogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLumaCatalogsResponseBody
	GetSuccess() *bool
}

type ListLumaCatalogsResponseBody struct {
	// The response code. A value of Success indicates a successful call. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of data catalogs bound to the agent, including entries and pagination information.
	Data *ListLumaCatalogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message. A value of Operation success is returned if the call succeeds. A specific error description is returned if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. Use this ID for troubleshooting and when submitting a ticket.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLumaCatalogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLumaCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLumaCatalogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLumaCatalogsResponseBody) GetData() *ListLumaCatalogsResponseBodyData {
	return s.Data
}

func (s *ListLumaCatalogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLumaCatalogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLumaCatalogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLumaCatalogsResponseBody) SetCode(v string) *ListLumaCatalogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLumaCatalogsResponseBody) SetData(v *ListLumaCatalogsResponseBodyData) *ListLumaCatalogsResponseBody {
	s.Data = v
	return s
}

func (s *ListLumaCatalogsResponseBody) SetMessage(v string) *ListLumaCatalogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLumaCatalogsResponseBody) SetRequestId(v string) *ListLumaCatalogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLumaCatalogsResponseBody) SetSuccess(v bool) *ListLumaCatalogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListLumaCatalogsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLumaCatalogsResponseBodyData struct {
	// The list of data catalogs bound to the agent.
	//
	// example:
	//
	// [{"Name":"my_catalog"}]
	Catalogs []*Catalog `json:"Catalogs,omitempty" xml:"Catalogs,omitempty" type:"Repeated"`
	// The effective page size for this request. If the Limit parameter is not specified, the server default value is used. If the specified value exceeds the upper limit, the value is adjusted to the maximum allowed value.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token for the next page. Pass this value as the NextToken in the next request to retrieve the next page. An empty value indicates that no more data is available.
	//
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of data catalogs bound to the agent, regardless of the number of entries returned on the current page.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLumaCatalogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLumaCatalogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLumaCatalogsResponseBodyData) GetCatalogs() []*Catalog {
	return s.Catalogs
}

func (s *ListLumaCatalogsResponseBodyData) GetLimit() *int32 {
	return s.Limit
}

func (s *ListLumaCatalogsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaCatalogsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLumaCatalogsResponseBodyData) SetCatalogs(v []*Catalog) *ListLumaCatalogsResponseBodyData {
	s.Catalogs = v
	return s
}

func (s *ListLumaCatalogsResponseBodyData) SetLimit(v int32) *ListLumaCatalogsResponseBodyData {
	s.Limit = &v
	return s
}

func (s *ListLumaCatalogsResponseBodyData) SetNextToken(v string) *ListLumaCatalogsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListLumaCatalogsResponseBodyData) SetTotalCount(v int32) *ListLumaCatalogsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLumaCatalogsResponseBodyData) Validate() error {
	if s.Catalogs != nil {
		for _, item := range s.Catalogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
