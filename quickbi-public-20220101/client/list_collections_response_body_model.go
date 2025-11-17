// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListCollectionsResponseBody
	GetRequestId() *string
	SetResult(v []*ListCollectionsResponseBodyResult) *ListCollectionsResponseBody
	GetResult() []*ListCollectionsResponseBodyResult
	SetSuccess(v bool) *ListCollectionsResponseBody
	GetSuccess() *bool
}

type ListCollectionsResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 162A632E-0A88-51CF-98F8-94FDEE82DB7D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the list of reports favored by the user.
	Result []*ListCollectionsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The primary key ID of the favorite record.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCollectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCollectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCollectionsResponseBody) GetResult() []*ListCollectionsResponseBodyResult {
	return s.Result
}

func (s *ListCollectionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCollectionsResponseBody) SetRequestId(v string) *ListCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCollectionsResponseBody) SetResult(v []*ListCollectionsResponseBodyResult) *ListCollectionsResponseBody {
	s.Result = v
	return s
}

func (s *ListCollectionsResponseBody) SetSuccess(v bool) *ListCollectionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCollectionsResponseBody) Validate() error {
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

type ListCollectionsResponseBodyResult struct {
	// The primary key ID of the favorite record.
	//
	// example:
	//
	// 12373
	FavoriteId *int32 `json:"FavoriteId,omitempty" xml:"FavoriteId,omitempty"`
	// The user ID of the work owner. This refers to the UserID in Quick BI, not the Alibaba Cloud UID.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the work.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// The name of the work.
	//
	// example:
	//
	// test
	WorksName *string `json:"WorksName,omitempty" xml:"WorksName,omitempty"`
	// The type of the work. Possible values:
	//
	// - DATAPRODUCT: Data Portal
	//
	// - PAGE: Dashboard
	//
	// - REPORT: Spreadsheet
	//
	// - dataForm: Data Entry Form
	//
	// - dashboardOfflineQuery: Self-service Data Extraction
	//
	// example:
	//
	// dashboardOfflineQuery
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
	// Workspace ID.
	//
	// example:
	//
	// 9337d121-a78f-4c1b-a8bc-f81de117****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// Workspace Name.
	//
	// example:
	//
	// test
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListCollectionsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCollectionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCollectionsResponseBodyResult) GetFavoriteId() *int32 {
	return s.FavoriteId
}

func (s *ListCollectionsResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListCollectionsResponseBodyResult) GetWorksId() *string {
	return s.WorksId
}

func (s *ListCollectionsResponseBodyResult) GetWorksName() *string {
	return s.WorksName
}

func (s *ListCollectionsResponseBodyResult) GetWorksType() *string {
	return s.WorksType
}

func (s *ListCollectionsResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListCollectionsResponseBodyResult) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListCollectionsResponseBodyResult) SetFavoriteId(v int32) *ListCollectionsResponseBodyResult {
	s.FavoriteId = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetOwnerId(v string) *ListCollectionsResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetWorksId(v string) *ListCollectionsResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetWorksName(v string) *ListCollectionsResponseBodyResult {
	s.WorksName = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetWorksType(v string) *ListCollectionsResponseBodyResult {
	s.WorksType = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetWorkspaceId(v string) *ListCollectionsResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetWorkspaceName(v string) *ListCollectionsResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
