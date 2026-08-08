// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCrossAccountsResponseBodyData) *ListCrossAccountsResponseBody
	GetData() *ListCrossAccountsResponseBodyData
	SetRequestId(v string) *ListCrossAccountsResponseBody
	GetRequestId() *string
}

type ListCrossAccountsResponseBody struct {
	Data *ListCrossAccountsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 5B2F09BF-CEBD-5A7E-AC01-E7F86169A5E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCrossAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCrossAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrossAccountsResponseBody) GetData() *ListCrossAccountsResponseBodyData {
	return s.Data
}

func (s *ListCrossAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrossAccountsResponseBody) SetData(v *ListCrossAccountsResponseBodyData) *ListCrossAccountsResponseBody {
	s.Data = v
	return s
}

func (s *ListCrossAccountsResponseBody) SetRequestId(v string) *ListCrossAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrossAccountsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCrossAccountsResponseBodyData struct {
	Content []*ListCrossAccountsResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eKDyCM0zFQ5op7jVMWmNNA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCrossAccountsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCrossAccountsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCrossAccountsResponseBodyData) GetContent() []*ListCrossAccountsResponseBodyDataContent {
	return s.Content
}

func (s *ListCrossAccountsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCrossAccountsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCrossAccountsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCrossAccountsResponseBodyData) SetContent(v []*ListCrossAccountsResponseBodyDataContent) *ListCrossAccountsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListCrossAccountsResponseBodyData) SetMaxResults(v int32) *ListCrossAccountsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListCrossAccountsResponseBodyData) SetNextToken(v string) *ListCrossAccountsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListCrossAccountsResponseBodyData) SetTotalCount(v int64) *ListCrossAccountsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCrossAccountsResponseBodyData) Validate() error {
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

type ListCrossAccountsResponseBodyDataContent struct {
	// example:
	//
	// 1773738311
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// MANUAL
	ManagementMode *string `json:"ManagementMode,omitempty" xml:"ManagementMode,omitempty"`
	// example:
	//
	// r-***
	ParentTargetId *string `json:"ParentTargetId,omitempty" xml:"ParentTargetId,omitempty"`
	// example:
	//
	// 123***7890
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// example:
	//
	// ACCOUNT
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListCrossAccountsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListCrossAccountsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListCrossAccountsResponseBodyDataContent) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListCrossAccountsResponseBodyDataContent) GetManagementMode() *string {
	return s.ManagementMode
}

func (s *ListCrossAccountsResponseBodyDataContent) GetParentTargetId() *string {
	return s.ParentTargetId
}

func (s *ListCrossAccountsResponseBodyDataContent) GetTargetId() *string {
	return s.TargetId
}

func (s *ListCrossAccountsResponseBodyDataContent) GetTargetType() *string {
	return s.TargetType
}

func (s *ListCrossAccountsResponseBodyDataContent) SetCreatedTime(v int64) *ListCrossAccountsResponseBodyDataContent {
	s.CreatedTime = &v
	return s
}

func (s *ListCrossAccountsResponseBodyDataContent) SetManagementMode(v string) *ListCrossAccountsResponseBodyDataContent {
	s.ManagementMode = &v
	return s
}

func (s *ListCrossAccountsResponseBodyDataContent) SetParentTargetId(v string) *ListCrossAccountsResponseBodyDataContent {
	s.ParentTargetId = &v
	return s
}

func (s *ListCrossAccountsResponseBodyDataContent) SetTargetId(v string) *ListCrossAccountsResponseBodyDataContent {
	s.TargetId = &v
	return s
}

func (s *ListCrossAccountsResponseBodyDataContent) SetTargetType(v string) *ListCrossAccountsResponseBodyDataContent {
	s.TargetType = &v
	return s
}

func (s *ListCrossAccountsResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
