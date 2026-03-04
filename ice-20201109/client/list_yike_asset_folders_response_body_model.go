// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYikeAssetFoldersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolderList(v []*ListYikeAssetFoldersResponseBodyFolderList) *ListYikeAssetFoldersResponseBody
	GetFolderList() []*ListYikeAssetFoldersResponseBodyFolderList
	SetMaxResults(v int32) *ListYikeAssetFoldersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListYikeAssetFoldersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListYikeAssetFoldersResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListYikeAssetFoldersResponseBody
	GetTotal() *int32
}

type ListYikeAssetFoldersResponseBody struct {
	FolderList []*ListYikeAssetFoldersResponseBodyFolderList `json:"FolderList,omitempty" xml:"FolderList,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 78
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListYikeAssetFoldersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListYikeAssetFoldersResponseBody) GoString() string {
	return s.String()
}

func (s *ListYikeAssetFoldersResponseBody) GetFolderList() []*ListYikeAssetFoldersResponseBodyFolderList {
	return s.FolderList
}

func (s *ListYikeAssetFoldersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListYikeAssetFoldersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListYikeAssetFoldersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListYikeAssetFoldersResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListYikeAssetFoldersResponseBody) SetFolderList(v []*ListYikeAssetFoldersResponseBodyFolderList) *ListYikeAssetFoldersResponseBody {
	s.FolderList = v
	return s
}

func (s *ListYikeAssetFoldersResponseBody) SetMaxResults(v int32) *ListYikeAssetFoldersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListYikeAssetFoldersResponseBody) SetNextToken(v string) *ListYikeAssetFoldersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListYikeAssetFoldersResponseBody) SetRequestId(v string) *ListYikeAssetFoldersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListYikeAssetFoldersResponseBody) SetTotal(v int32) *ListYikeAssetFoldersResponseBody {
	s.Total = &v
	return s
}

func (s *ListYikeAssetFoldersResponseBody) Validate() error {
	if s.FolderList != nil {
		for _, item := range s.FolderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListYikeAssetFoldersResponseBodyFolderList struct {
	// example:
	//
	// 2026-01-22T02:07:06Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// fd-EXRyxc5SHY
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// example:
	//
	// default
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// example:
	//
	// 1
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// pd_183320223010****
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// example:
	//
	// llm-odl2p61i4vfbph4g
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListYikeAssetFoldersResponseBodyFolderList) String() string {
	return dara.Prettify(s)
}

func (s ListYikeAssetFoldersResponseBodyFolderList) GoString() string {
	return s.String()
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) GetFolderId() *string {
	return s.FolderId
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) GetFolderName() *string {
	return s.FolderName
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) GetProductionId() *string {
	return s.ProductionId
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) SetCreateTime(v string) *ListYikeAssetFoldersResponseBodyFolderList {
	s.CreateTime = &v
	return s
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) SetFolderId(v string) *ListYikeAssetFoldersResponseBodyFolderList {
	s.FolderId = &v
	return s
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) SetFolderName(v string) *ListYikeAssetFoldersResponseBodyFolderList {
	s.FolderName = &v
	return s
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) SetIsDefault(v bool) *ListYikeAssetFoldersResponseBodyFolderList {
	s.IsDefault = &v
	return s
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) SetProductionId(v string) *ListYikeAssetFoldersResponseBodyFolderList {
	s.ProductionId = &v
	return s
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) SetWorkspaceId(v string) *ListYikeAssetFoldersResponseBodyFolderList {
	s.WorkspaceId = &v
	return s
}

func (s *ListYikeAssetFoldersResponseBodyFolderList) Validate() error {
	return dara.Validate(s)
}
