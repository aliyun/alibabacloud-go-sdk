// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUserPoolsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserPoolsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUserPoolsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUserPoolsResponseBody
	GetTotalCount() *int32
	SetUserPools(v []*ListUserPoolsResponseBodyUserPools) *ListUserPoolsResponseBody
	GetUserPools() []*ListUserPoolsResponseBodyUserPools
}

type ListUserPoolsResponseBody struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// TGlzdFVzZXJQb29sczo6MjA=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserPools  []*ListUserPoolsResponseBodyUserPools `json:"UserPools,omitempty" xml:"UserPools,omitempty" type:"Repeated"`
}

func (s ListUserPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserPoolsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserPoolsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserPoolsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserPoolsResponseBody) GetUserPools() []*ListUserPoolsResponseBodyUserPools {
	return s.UserPools
}

func (s *ListUserPoolsResponseBody) SetMaxResults(v int32) *ListUserPoolsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserPoolsResponseBody) SetNextToken(v string) *ListUserPoolsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserPoolsResponseBody) SetRequestId(v string) *ListUserPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserPoolsResponseBody) SetTotalCount(v int32) *ListUserPoolsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserPoolsResponseBody) SetUserPools(v []*ListUserPoolsResponseBodyUserPools) *ListUserPoolsResponseBody {
	s.UserPools = v
	return s
}

func (s *ListUserPoolsResponseBody) Validate() error {
	if s.UserPools != nil {
		for _, item := range s.UserPools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserPoolsResponseBodyUserPools struct {
	// example:
	//
	// 2026-05-07T06:19:17Z
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2026-05-07T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// up-xxxxxxxxxxxxxxxxxxxx
	UserPoolId *string `json:"UserPoolId,omitempty" xml:"UserPoolId,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s ListUserPoolsResponseBodyUserPools) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolsResponseBodyUserPools) GoString() string {
	return s.String()
}

func (s *ListUserPoolsResponseBodyUserPools) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserPoolsResponseBodyUserPools) GetDescription() *string {
	return s.Description
}

func (s *ListUserPoolsResponseBodyUserPools) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListUserPoolsResponseBodyUserPools) GetUserPoolId() *string {
	return s.UserPoolId
}

func (s *ListUserPoolsResponseBodyUserPools) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *ListUserPoolsResponseBodyUserPools) SetCreateTime(v string) *ListUserPoolsResponseBodyUserPools {
	s.CreateTime = &v
	return s
}

func (s *ListUserPoolsResponseBodyUserPools) SetDescription(v string) *ListUserPoolsResponseBodyUserPools {
	s.Description = &v
	return s
}

func (s *ListUserPoolsResponseBodyUserPools) SetUpdateTime(v string) *ListUserPoolsResponseBodyUserPools {
	s.UpdateTime = &v
	return s
}

func (s *ListUserPoolsResponseBodyUserPools) SetUserPoolId(v string) *ListUserPoolsResponseBodyUserPools {
	s.UserPoolId = &v
	return s
}

func (s *ListUserPoolsResponseBodyUserPools) SetUserPoolName(v string) *ListUserPoolsResponseBodyUserPools {
	s.UserPoolName = &v
	return s
}

func (s *ListUserPoolsResponseBodyUserPools) Validate() error {
	return dara.Validate(s)
}
