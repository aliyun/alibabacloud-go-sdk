// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDataSourceResponseBody
	GetRequestId() *string
	SetResult(v []*ListDataSourceResponseBodyResult) *ListDataSourceResponseBody
	GetResult() []*ListDataSourceResponseBodyResult
	SetSuccess(v bool) *ListDataSourceResponseBody
	GetSuccess() *bool
}

type ListDataSourceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 7FC9A6A6-****-5CED-B*****E891E4075
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Array of data source information.
	Result []*ListDataSourceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourceResponseBody) GetResult() []*ListDataSourceResponseBodyResult {
	return s.Result
}

func (s *ListDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataSourceResponseBody) SetRequestId(v string) *ListDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceResponseBody) SetResult(v []*ListDataSourceResponseBodyResult) *ListDataSourceResponseBody {
	s.Result = v
	return s
}

func (s *ListDataSourceResponseBody) SetSuccess(v bool) *ListDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataSourceResponseBody) Validate() error {
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

type ListDataSourceResponseBodyResult struct {
	// Quick BI user ID of the creator.
	//
	// example:
	//
	// 281*****-485******-8
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// Owner\\"s nickname.
	//
	// example:
	//
	// system
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// Data source ID.
	//
	// example:
	//
	// 7FC9A6A6-****-5CED-B*****E891E4075
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// Data source type.
	//
	// example:
	//
	// odps
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// Creation time of the data source, in yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2024-04-16 13:17:39
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 2024-08-15 10:06:31
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Primary data source type for multi-engine data sources.
	//
	// example:
	//
	// dataphin
	ParentDsType *string `json:"ParentDsType,omitempty" xml:"ParentDsType,omitempty"`
	// Display name of the data source.
	//
	// example:
	//
	// 0327
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s ListDataSourceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSourceResponseBodyResult) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDataSourceResponseBodyResult) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListDataSourceResponseBodyResult) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *ListDataSourceResponseBodyResult) GetDsType() *string {
	return s.DsType
}

func (s *ListDataSourceResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDataSourceResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDataSourceResponseBodyResult) GetParentDsType() *string {
	return s.ParentDsType
}

func (s *ListDataSourceResponseBodyResult) GetShowName() *string {
	return s.ShowName
}

func (s *ListDataSourceResponseBodyResult) SetCreatorId(v string) *ListDataSourceResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *ListDataSourceResponseBodyResult) SetCreatorName(v string) *ListDataSourceResponseBodyResult {
	s.CreatorName = &v
	return s
}

func (s *ListDataSourceResponseBodyResult) SetDatasourceId(v string) *ListDataSourceResponseBodyResult {
	s.DatasourceId = &v
	return s
}

func (s *ListDataSourceResponseBodyResult) SetDsType(v string) *ListDataSourceResponseBodyResult {
	s.DsType = &v
	return s
}

func (s *ListDataSourceResponseBodyResult) SetGmtCreate(v string) *ListDataSourceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListDataSourceResponseBodyResult) SetGmtModified(v string) *ListDataSourceResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListDataSourceResponseBodyResult) SetParentDsType(v string) *ListDataSourceResponseBodyResult {
	s.ParentDsType = &v
	return s
}

func (s *ListDataSourceResponseBodyResult) SetShowName(v string) *ListDataSourceResponseBodyResult {
	s.ShowName = &v
	return s
}

func (s *ListDataSourceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
