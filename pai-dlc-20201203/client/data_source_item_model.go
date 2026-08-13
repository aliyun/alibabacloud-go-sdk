// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSourceItem interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *DataSourceItem
	GetDataSourceId() *string
	SetDataSourceType(v string) *DataSourceItem
	GetDataSourceType() *string
	SetDescription(v string) *DataSourceItem
	GetDescription() *string
	SetDisplayName(v string) *DataSourceItem
	GetDisplayName() *string
	SetEndpoint(v string) *DataSourceItem
	GetEndpoint() *string
	SetFileSystemId(v string) *DataSourceItem
	GetFileSystemId() *string
	SetGmtCreateTime(v string) *DataSourceItem
	GetGmtCreateTime() *string
	SetGmtModifyTime(v string) *DataSourceItem
	GetGmtModifyTime() *string
	SetMountPath(v string) *DataSourceItem
	GetMountPath() *string
	SetOptions(v string) *DataSourceItem
	GetOptions() *string
	SetPath(v string) *DataSourceItem
	GetPath() *string
	SetUserId(v string) *DataSourceItem
	GetUserId() *string
}

type DataSourceItem struct {
	// The data source ID.
	//
	// example:
	//
	// data-20210114104214-xxxxxxx
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The data source type. Only File Storage NAS (NAS) file systems are supported.
	//
	// example:
	//
	// nas
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The data source description.
	//
	// example:
	//
	// data source of dlc examples
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The data source name.
	//
	// example:
	//
	// nas-data
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The endpoint of the Alibaba Cloud OSS file system.
	//
	// example:
	//
	// oss-cn-beijing-internal.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The ID of the NAS file system. For more information, see [Alibaba Cloud documentation](https://help.aliyun.com/document_detail/185212.html).
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The time when the data source was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-12T14:35:00Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the data source was last modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-12T14:36:00Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// The local path of the container that is attached to the runtime of the task.
	//
	// example:
	//
	// /root/data/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The configuration options of the Alibaba Cloud OSS file system.
	//
	// example:
	//
	// {"key": "value"}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The path of the Alibaba Cloud Object Storage Service (OSS) file system.
	//
	// example:
	//
	// oss://mybucket/path/to/dir
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The ID of the Alibaba Cloud account that creates the data source.
	//
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DataSourceItem) String() string {
	return dara.Prettify(s)
}

func (s DataSourceItem) GoString() string {
	return s.String()
}

func (s *DataSourceItem) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DataSourceItem) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *DataSourceItem) GetDescription() *string {
	return s.Description
}

func (s *DataSourceItem) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DataSourceItem) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DataSourceItem) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DataSourceItem) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *DataSourceItem) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *DataSourceItem) GetMountPath() *string {
	return s.MountPath
}

func (s *DataSourceItem) GetOptions() *string {
	return s.Options
}

func (s *DataSourceItem) GetPath() *string {
	return s.Path
}

func (s *DataSourceItem) GetUserId() *string {
	return s.UserId
}

func (s *DataSourceItem) SetDataSourceId(v string) *DataSourceItem {
	s.DataSourceId = &v
	return s
}

func (s *DataSourceItem) SetDataSourceType(v string) *DataSourceItem {
	s.DataSourceType = &v
	return s
}

func (s *DataSourceItem) SetDescription(v string) *DataSourceItem {
	s.Description = &v
	return s
}

func (s *DataSourceItem) SetDisplayName(v string) *DataSourceItem {
	s.DisplayName = &v
	return s
}

func (s *DataSourceItem) SetEndpoint(v string) *DataSourceItem {
	s.Endpoint = &v
	return s
}

func (s *DataSourceItem) SetFileSystemId(v string) *DataSourceItem {
	s.FileSystemId = &v
	return s
}

func (s *DataSourceItem) SetGmtCreateTime(v string) *DataSourceItem {
	s.GmtCreateTime = &v
	return s
}

func (s *DataSourceItem) SetGmtModifyTime(v string) *DataSourceItem {
	s.GmtModifyTime = &v
	return s
}

func (s *DataSourceItem) SetMountPath(v string) *DataSourceItem {
	s.MountPath = &v
	return s
}

func (s *DataSourceItem) SetOptions(v string) *DataSourceItem {
	s.Options = &v
	return s
}

func (s *DataSourceItem) SetPath(v string) *DataSourceItem {
	s.Path = &v
	return s
}

func (s *DataSourceItem) SetUserId(v string) *DataSourceItem {
	s.UserId = &v
	return s
}

func (s *DataSourceItem) Validate() error {
	return dara.Validate(s)
}
