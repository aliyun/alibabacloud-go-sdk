// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYikeProductionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListYikeProductionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListYikeProductionsResponseBody
	GetNextToken() *string
	SetProductionList(v []*ListYikeProductionsResponseBodyProductionList) *ListYikeProductionsResponseBody
	GetProductionList() []*ListYikeProductionsResponseBodyProductionList
	SetRequestId(v string) *ListYikeProductionsResponseBody
	GetRequestId() *string
}

type ListYikeProductionsResponseBody struct {
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// Token
	NextToken      *string                                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProductionList []*ListYikeProductionsResponseBodyProductionList `json:"ProductionList,omitempty" xml:"ProductionList,omitempty" type:"Repeated"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListYikeProductionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListYikeProductionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListYikeProductionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListYikeProductionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListYikeProductionsResponseBody) GetProductionList() []*ListYikeProductionsResponseBodyProductionList {
	return s.ProductionList
}

func (s *ListYikeProductionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListYikeProductionsResponseBody) SetMaxResults(v int32) *ListYikeProductionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListYikeProductionsResponseBody) SetNextToken(v string) *ListYikeProductionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListYikeProductionsResponseBody) SetProductionList(v []*ListYikeProductionsResponseBodyProductionList) *ListYikeProductionsResponseBody {
	s.ProductionList = v
	return s
}

func (s *ListYikeProductionsResponseBody) SetRequestId(v string) *ListYikeProductionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListYikeProductionsResponseBody) Validate() error {
	if s.ProductionList != nil {
		for _, item := range s.ProductionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListYikeProductionsResponseBodyProductionList struct {
	// example:
	//
	// Manage
	Auth *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	// example:
	//
	// https://tagvvcloud-material-center-prod.oss-cn-hangzhou.aliyuncs.com/sumvideo/utils_image/sumvideo-video-cover.png
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// example:
	//
	// 2026-01-07T02:21:36Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// aliyun_183320223010****
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// Swas_QuerySwasInstanceByRegion
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// pd_463862****
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// example:
	//
	// Harvest
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// llm-m3r546h1n9kq3mtm
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListYikeProductionsResponseBodyProductionList) String() string {
	return dara.Prettify(s)
}

func (s ListYikeProductionsResponseBodyProductionList) GoString() string {
	return s.String()
}

func (s *ListYikeProductionsResponseBodyProductionList) GetAuth() *string {
	return s.Auth
}

func (s *ListYikeProductionsResponseBodyProductionList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListYikeProductionsResponseBodyProductionList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListYikeProductionsResponseBodyProductionList) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListYikeProductionsResponseBodyProductionList) GetDescription() *string {
	return s.Description
}

func (s *ListYikeProductionsResponseBodyProductionList) GetProductionId() *string {
	return s.ProductionId
}

func (s *ListYikeProductionsResponseBodyProductionList) GetTitle() *string {
	return s.Title
}

func (s *ListYikeProductionsResponseBodyProductionList) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListYikeProductionsResponseBodyProductionList) SetAuth(v string) *ListYikeProductionsResponseBodyProductionList {
	s.Auth = &v
	return s
}

func (s *ListYikeProductionsResponseBodyProductionList) SetCoverUrl(v string) *ListYikeProductionsResponseBodyProductionList {
	s.CoverUrl = &v
	return s
}

func (s *ListYikeProductionsResponseBodyProductionList) SetCreateTime(v string) *ListYikeProductionsResponseBodyProductionList {
	s.CreateTime = &v
	return s
}

func (s *ListYikeProductionsResponseBodyProductionList) SetCreateUserName(v string) *ListYikeProductionsResponseBodyProductionList {
	s.CreateUserName = &v
	return s
}

func (s *ListYikeProductionsResponseBodyProductionList) SetDescription(v string) *ListYikeProductionsResponseBodyProductionList {
	s.Description = &v
	return s
}

func (s *ListYikeProductionsResponseBodyProductionList) SetProductionId(v string) *ListYikeProductionsResponseBodyProductionList {
	s.ProductionId = &v
	return s
}

func (s *ListYikeProductionsResponseBodyProductionList) SetTitle(v string) *ListYikeProductionsResponseBodyProductionList {
	s.Title = &v
	return s
}

func (s *ListYikeProductionsResponseBodyProductionList) SetWorkspaceId(v string) *ListYikeProductionsResponseBodyProductionList {
	s.WorkspaceId = &v
	return s
}

func (s *ListYikeProductionsResponseBodyProductionList) Validate() error {
	return dara.Validate(s)
}
