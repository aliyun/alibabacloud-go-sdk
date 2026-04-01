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
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxRecordsPerPage *int32 `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	// The details of the minor engine version.
	MinorVersionItems []*DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems `json:"MinorVersionItems,omitempty" xml:"MinorVersionItems,omitempty" type:"Repeated"`
	// The page number returned.
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
	// The total number of entries returned.
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
	// The PostgreSQL version to which the minor engine version corresponds. For more information, see [Release notes for AliPG](https://help.aliyun.com/document_detail/126002.html).
	//
	// >  This parameter is available only for instances that run **PostgreSQL**.
	//
	// example:
	//
	// 13.6
	CommunityMinorVersion *string `json:"CommunityMinorVersion,omitempty" xml:"CommunityMinorVersion,omitempty"`
	// The database engine that corresponds to the minor engine version.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version that corresponds to the minor engine version.
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
	// 	- **vaild**
	//
	// 	- **expired**
	//
	// >  If the minor engine version is in the Offline state, the minor engine version is discontinued. In this case, ignore the expiration status. If the minor engine version is in the Online state and the expiration state is expired, the minor engine version expires. If the expiration state is vaild, the minor engine version is still in its lifecycle.
	//
	// example:
	//
	// vaild
	ExpireStatus *string `json:"ExpireStatus,omitempty" xml:"ExpireStatus,omitempty"`
	// An internal parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// True
	IsHotfixVersion *bool `json:"IsHotfixVersion,omitempty" xml:"IsHotfixVersion,omitempty"`
	// The minor engine version.
	//
	// example:
	//
	// rds_20220731
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The RDS edition of the instance that runs the minor engine version. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition
	//
	// 	- **HighAvailability**: RDS High-availability Edition
	//
	// 	- **Finance**: RDS Enterprise Edition
	//
	// example:
	//
	// HighAvailability
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The URL of the release notes for the minor engine version.
	//
	// example:
	//
	// https://example.com
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The release type. Valid values:
	//
	// 	- **LTS**: a long-term version
	//
	// 	- **BETA**: a preview version
	//
	// example:
	//
	// BETA
	ReleaseType *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	// The status of the minor engine version. Valid values:
	//
	// 	- **Offline**: discontinued
	//
	// 	- **Online**: available
	//
	// >  If the minor engine version is in the Offline state, the minor engine version is discontinued. In this case, ignore the expiration status. If the minor engine version is in the Online state and the expiration state is expired, the minor engine version expires. If the expiration state is vaild, the minor engine version is still in its lifecycle.
	//
	// example:
	//
	// Online
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The tag that corresponds to the minor engine version. Valid values:
	//
	// 	- **pgsql_docker_image**: tag of common instances
	//
	// 	- **pgsql_babelfish_image**: tag of instances for which Babelfish is enabled
	//
	// >  This parameter is available only for instances that run **PostgreSQL**.
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
