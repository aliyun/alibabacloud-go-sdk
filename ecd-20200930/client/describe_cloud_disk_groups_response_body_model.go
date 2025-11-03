// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudDriveGroups(v []*DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) *DescribeCloudDiskGroupsResponseBody
	GetCloudDriveGroups() []*DescribeCloudDiskGroupsResponseBodyCloudDriveGroups
	SetCount(v int64) *DescribeCloudDiskGroupsResponseBody
	GetCount() *int64
	SetNextToken(v string) *DescribeCloudDiskGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeCloudDiskGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCloudDiskGroupsResponseBody
	GetSuccess() *bool
}

type DescribeCloudDiskGroupsResponseBody struct {
	CloudDriveGroups []*DescribeCloudDiskGroupsResponseBodyCloudDriveGroups `json:"CloudDriveGroups,omitempty" xml:"CloudDriveGroups,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// D648DBF7-9476-53D6-98AB-674836021DFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudDiskGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskGroupsResponseBody) GetCloudDriveGroups() []*DescribeCloudDiskGroupsResponseBodyCloudDriveGroups {
	return s.CloudDriveGroups
}

func (s *DescribeCloudDiskGroupsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCloudDiskGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudDiskGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudDiskGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCloudDiskGroupsResponseBody) SetCloudDriveGroups(v []*DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) *DescribeCloudDiskGroupsResponseBody {
	s.CloudDriveGroups = v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBody) SetCount(v int64) *DescribeCloudDiskGroupsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBody) SetNextToken(v string) *DescribeCloudDiskGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBody) SetRequestId(v string) *DescribeCloudDiskGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBody) SetSuccess(v bool) *DescribeCloudDiskGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBody) Validate() error {
	if s.CloudDriveGroups != nil {
		for _, item := range s.CloudDriveGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudDiskGroupsResponseBodyCloudDriveGroups struct {
	// example:
	//
	// 2022-04-11T07:44:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-959593****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// 1234
	DriveId *string `json:"DriveId,omitempty" xml:"DriveId,omitempty"`
	// example:
	//
	// cg-e70ga4ixp30ur****
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// org-aliyun-wy-org-id
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 5368709120
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// example:
	//
	// 1024000000
	UsedSize *string `json:"UsedSize,omitempty" xml:"UsedSize,omitempty"`
}

func (s DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) GetDriveId() *string {
	return s.DriveId
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) GetOrgId() *string {
	return s.OrgId
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) GetUsedSize() *string {
	return s.UsedSize
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) SetCreateTime(v string) *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) SetDirectoryId(v string) *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups {
	s.DirectoryId = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) SetDriveId(v string) *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups {
	s.DriveId = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) SetGroupId(v string) *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) SetGroupName(v string) *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) SetOrgId(v string) *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups {
	s.OrgId = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) SetStatus(v string) *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups {
	s.Status = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) SetTotalSize(v int64) *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups {
	s.TotalSize = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) SetUsedSize(v string) *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups {
	s.UsedSize = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponseBodyCloudDriveGroups) Validate() error {
	return dara.Validate(s)
}
