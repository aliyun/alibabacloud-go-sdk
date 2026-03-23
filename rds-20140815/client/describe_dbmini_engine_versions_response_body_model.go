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
	DBInstanceId      *string                                                      `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	MaxRecordsPerPage *int32                                                       `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	MinorVersionItems []*DescribeDBMiniEngineVersionsResponseBodyMinorVersionItems `json:"MinorVersionItems,omitempty" xml:"MinorVersionItems,omitempty" type:"Repeated"`
	PageNumbers       *int32                                                       `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	RequestId         *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CommunityMinorVersion *string `json:"CommunityMinorVersion,omitempty" xml:"CommunityMinorVersion,omitempty"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireDate            *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// example:
	//
	// vaild
	ExpireStatus    *string `json:"ExpireStatus,omitempty" xml:"ExpireStatus,omitempty"`
	IsHotfixVersion *bool   `json:"IsHotfixVersion,omitempty" xml:"IsHotfixVersion,omitempty"`
	MinorVersion    *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	NodeType        *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	ReleaseNote     *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseType     *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	StatusDesc      *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	Tag             *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
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
