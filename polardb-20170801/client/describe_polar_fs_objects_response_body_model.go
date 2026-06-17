// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribePolarFsObjectsResponseBodyItems) *DescribePolarFsObjectsResponseBody
	GetItems() []*DescribePolarFsObjectsResponseBodyItems
	SetPageRecordCount(v string) *DescribePolarFsObjectsResponseBody
	GetPageRecordCount() *string
	SetPageSize(v string) *DescribePolarFsObjectsResponseBody
	GetPageSize() *string
	SetPath(v string) *DescribePolarFsObjectsResponseBody
	GetPath() *string
	SetPolarFsInstanceId(v string) *DescribePolarFsObjectsResponseBody
	GetPolarFsInstanceId() *string
	SetRequestId(v string) *DescribePolarFsObjectsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribePolarFsObjectsResponseBody
	GetTotalRecordCount() *string
}

type DescribePolarFsObjectsResponseBody struct {
	// The files and subdirectories in the specified path.
	Items []*DescribePolarFsObjectsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of records returned on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The number of records to return per page. Valid values: 30, 50, and 100. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The path that was queried.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The PolarFs instance ID.
	//
	// example:
	//
	// pfs-2ze0i7*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that match the query.
	//
	// example:
	//
	// 50
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribePolarFsObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarFsObjectsResponseBody) GetItems() []*DescribePolarFsObjectsResponseBodyItems {
	return s.Items
}

func (s *DescribePolarFsObjectsResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribePolarFsObjectsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribePolarFsObjectsResponseBody) GetPath() *string {
	return s.Path
}

func (s *DescribePolarFsObjectsResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarFsObjectsResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribePolarFsObjectsResponseBody) SetItems(v []*DescribePolarFsObjectsResponseBodyItems) *DescribePolarFsObjectsResponseBody {
	s.Items = v
	return s
}

func (s *DescribePolarFsObjectsResponseBody) SetPageRecordCount(v string) *DescribePolarFsObjectsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBody) SetPageSize(v string) *DescribePolarFsObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBody) SetPath(v string) *DescribePolarFsObjectsResponseBody {
	s.Path = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBody) SetPolarFsInstanceId(v string) *DescribePolarFsObjectsResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBody) SetRequestId(v string) *DescribePolarFsObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBody) SetTotalRecordCount(v string) *DescribePolarFsObjectsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarFsObjectsResponseBodyItems struct {
	// The size of the item in bytes.
	//
	// example:
	//
	// 100
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The time when the item was created, as a UNIX timestamp in seconds.
	//
	// example:
	//
	// 2025-03-25T09:37:10Z
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The time the item was last modified, as a UNIX timestamp in seconds.
	//
	// example:
	//
	// 2025-03-25T09:37:10Z
	LastModified *int64 `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The target path of the symbolic link.
	//
	// example:
	//
	// test
	LinkTarget *string `json:"LinkTarget,omitempty" xml:"LinkTarget,omitempty"`
	// The file system permissions in octal format.
	//
	// example:
	//
	// 755
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The number of hard links to the item.
	//
	// example:
	//
	// test
	Nlink *int32 `json:"Nlink,omitempty" xml:"Nlink,omitempty"`
	// The owner of the file or directory.
	//
	// example:
	//
	// test
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The absolute path of the file or subdirectory.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The file type.
	//
	// example:
	//
	// directory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePolarFsObjectsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsObjectsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribePolarFsObjectsResponseBodyItems) GetCapacity() *string {
	return s.Capacity
}

func (s *DescribePolarFsObjectsResponseBodyItems) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribePolarFsObjectsResponseBodyItems) GetLastModified() *int64 {
	return s.LastModified
}

func (s *DescribePolarFsObjectsResponseBodyItems) GetLinkTarget() *string {
	return s.LinkTarget
}

func (s *DescribePolarFsObjectsResponseBodyItems) GetMode() *string {
	return s.Mode
}

func (s *DescribePolarFsObjectsResponseBodyItems) GetNlink() *int32 {
	return s.Nlink
}

func (s *DescribePolarFsObjectsResponseBodyItems) GetOwner() *string {
	return s.Owner
}

func (s *DescribePolarFsObjectsResponseBodyItems) GetPath() *string {
	return s.Path
}

func (s *DescribePolarFsObjectsResponseBodyItems) GetType() *string {
	return s.Type
}

func (s *DescribePolarFsObjectsResponseBodyItems) SetCapacity(v string) *DescribePolarFsObjectsResponseBodyItems {
	s.Capacity = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBodyItems) SetCreationTime(v int64) *DescribePolarFsObjectsResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBodyItems) SetLastModified(v int64) *DescribePolarFsObjectsResponseBodyItems {
	s.LastModified = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBodyItems) SetLinkTarget(v string) *DescribePolarFsObjectsResponseBodyItems {
	s.LinkTarget = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBodyItems) SetMode(v string) *DescribePolarFsObjectsResponseBodyItems {
	s.Mode = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBodyItems) SetNlink(v int32) *DescribePolarFsObjectsResponseBodyItems {
	s.Nlink = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBodyItems) SetOwner(v string) *DescribePolarFsObjectsResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBodyItems) SetPath(v string) *DescribePolarFsObjectsResponseBodyItems {
	s.Path = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBodyItems) SetType(v string) *DescribePolarFsObjectsResponseBodyItems {
	s.Type = &v
	return s
}

func (s *DescribePolarFsObjectsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
