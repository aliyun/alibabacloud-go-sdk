// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBMiniEngineVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBMiniEngineVersionsResponseBody
	GetDBInstanceId() *string
	SetMaxRecordsPerPage(v int32) *DescribeDBMiniEngineVersionsResponseBody
	GetMaxRecordsPerPage() *int32
	SetMinorVersionItems(v []*DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) *DescribeDBMiniEngineVersionsResponseBody
	GetMinorVersionItems() []*DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems
	SetPageNumbers(v int32) *DescribeDBMiniEngineVersionsResponseBody
	GetPageNumbers() *int32
	SetRequestId(v string) *DescribeDBMiniEngineVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDBMiniEngineVersionsResponseBody
	GetTotalCount() *int32
}

type DescribeDBMiniEngineVersionsResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	MaxRecordsPerPage *int32 `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	// The list of minor engine versions.
	MinorVersionItems []*DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems `json:"MinorVersionItems,omitempty" xml:"MinorVersionItems,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumbers *int32 `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EFB6083A-7699-489B-8278-C0CB4793A96E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBMiniEngineVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBMiniEngineVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBMiniEngineVersionsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBMiniEngineVersionsResponseBody) GetMaxRecordsPerPage() *int32 {
	return s.MaxRecordsPerPage
}

func (s *DescribeDBMiniEngineVersionsResponseBody) GetMinorVersionItems() []*DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	return s.MinorVersionItems
}

func (s *DescribeDBMiniEngineVersionsResponseBody) GetPageNumbers() *int32 {
	return s.PageNumbers
}

func (s *DescribeDBMiniEngineVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBMiniEngineVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDBMiniEngineVersionsResponseBody) SetDBInstanceId(v string) *DescribeDBMiniEngineVersionsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBody) SetMaxRecordsPerPage(v int32) *DescribeDBMiniEngineVersionsResponseBody {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBody) SetMinorVersionItems(v []*DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) *DescribeDBMiniEngineVersionsResponseBody {
	s.MinorVersionItems = v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBody) SetPageNumbers(v int32) *DescribeDBMiniEngineVersionsResponseBody {
	s.PageNumbers = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBody) SetRequestId(v string) *DescribeDBMiniEngineVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBody) SetTotalCount(v int32) *DescribeDBMiniEngineVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBody) Validate() error {
	if s.MinorVersionItems != nil {
		for _, item := range s.MinorVersionItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems struct {
	// The community minor version that corresponds to the minor engine version.
	//
	// example:
	//
	// 5.7.38
	CommunityMinorVersion *string `json:"CommunityMinorVersion,omitempty" xml:"CommunityMinorVersion,omitempty"`
	// The database engine that corresponds to the minor version.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version that corresponds to the minor version.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The expiration time of the minor engine version.
	//
	// example:
	//
	// 20231213
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The expiration status of the minor engine version. Valid values:
	//
	// - **vaild**: Milvus version is valid.
	//
	// - **expired**: Milvus version has expired.
	//
	// > If the offline status is Offline, Milvus version has been taken offline and the expiration status is ignored. If the offline status is Online and the expiration status is expired, Milvus version has exceeded its lifecycle. If the offline status is Online and the expiration status is vaild, Milvus version is still within its lifecycle.
	//
	// example:
	//
	// vaild
	ExpireStatus *string `json:"ExpireStatus,omitempty" xml:"ExpireStatus,omitempty"`
	// An internal parameter. You can ignore this parameter.
	//
	// example:
	//
	// True
	IsHotfixVersion *bool `json:"IsHotfixVersion,omitempty" xml:"IsHotfixVersion,omitempty"`
	// The version number of the minor engine version.
	//
	// example:
	//
	// rds_20220731
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The instance edition that corresponds to the minor version. Valid values:
	//
	// 	- **Basic**: Basic Edition.
	//
	// 	- **HighAvailability**: high-availability series.
	//
	// 	- **Finance**: RDS Enterprise Edition.
	//
	// example:
	//
	// HighAvailability
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The URL of the release notes for the minor version.
	//
	// example:
	//
	// https://example.com
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The release type. Valid values:
	//
	// 	- **LTS**: Long-term support version.
	//
	// 	- **BETA**: Preview version.
	//
	// example:
	//
	// BETA
	ReleaseType *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	// The offline status of the minor engine version. Valid values:
	//
	// - **Offline**: Milvus version has been taken offline.
	//
	// - **Online**: Milvus version is online.
	//
	// > If the offline status is Offline, Milvus version has been taken offline and the expiration status is ignored. If the offline status is Online and the expiration status is expired, Milvus version has exceeded its lifecycle. If the offline status is Online and the expiration status is vaild, Milvus version is still within its lifecycle.
	//
	// example:
	//
	// Online
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The tag that corresponds to the minor engine version. Valid values:
	//
	// - **pgsql_docker_image**: general instance tag.
	//
	// - **pgsql_babelfish_image**: Babelfish instance tag.
	//
	// > This value is returned only for **PostgreSQL**.
	//
	// example:
	//
	// pgsql_babelfish_image
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GoString() string {
	return s.String()
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetCommunityMinorVersion() *string {
	return s.CommunityMinorVersion
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetExpireStatus() *string {
	return s.ExpireStatus
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetIsHotfixVersion() *bool {
	return s.IsHotfixVersion
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) GetTag() *string {
	return s.Tag
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetCommunityMinorVersion(v string) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.CommunityMinorVersion = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetEngine(v string) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.Engine = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetEngineVersion(v string) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetExpireDate(v string) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetExpireStatus(v string) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.ExpireStatus = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetIsHotfixVersion(v bool) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.IsHotfixVersion = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetMinorVersion(v string) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetNodeType(v string) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.NodeType = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetReleaseNote(v string) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetReleaseType(v string) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.ReleaseType = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetStatusDesc(v string) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.StatusDesc = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) SetTag(v string) *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems {
	s.Tag = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems) Validate() error {
	return dara.Validate(s)
}
