// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGroupDirectoriesResponseBody
	GetCode() *string
	SetDirectories(v []*ListGroupDirectoriesResponseBodyDirectories) *ListGroupDirectoriesResponseBody
	GetDirectories() []*ListGroupDirectoriesResponseBodyDirectories
	SetMessage(v string) *ListGroupDirectoriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGroupDirectoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListGroupDirectoriesResponseBody
	GetTotalCount() *int64
}

type ListGroupDirectoriesResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The query root itself and all its descendant directories, including visible referenced directories in the space. The results are not paginated.
	//
	// example:
	//
	// []
	Directories []*ListGroupDirectoriesResponseBodyDirectories `json:"directories,omitempty" xml:"directories,omitempty" type:"Repeated"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The number of returned directories, which equals the length of the directories array.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListGroupDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupDirectoriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGroupDirectoriesResponseBody) GetDirectories() []*ListGroupDirectoriesResponseBodyDirectories {
	return s.Directories
}

func (s *ListGroupDirectoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGroupDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupDirectoriesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGroupDirectoriesResponseBody) SetCode(v string) *ListGroupDirectoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListGroupDirectoriesResponseBody) SetDirectories(v []*ListGroupDirectoriesResponseBodyDirectories) *ListGroupDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *ListGroupDirectoriesResponseBody) SetMessage(v string) *ListGroupDirectoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListGroupDirectoriesResponseBody) SetRequestId(v string) *ListGroupDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupDirectoriesResponseBody) SetTotalCount(v int64) *ListGroupDirectoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupDirectoriesResponseBody) Validate() error {
	if s.Directories != nil {
		for _, item := range s.Directories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupDirectoriesResponseBodyDirectories struct {
	// The directory description.
	//
	// example:
	//
	// Project description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID, including the query root itself and its descendants.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The original directory type. The value is GROUP for physical directories in the space. Referenced directories retain their original type.
	//
	// example:
	//
	// GROUP
	DirectoryType *string `json:"directoryType,omitempty" xml:"directoryType,omitempty"`
	// The directory name.
	//
	// example:
	//
	// Project resources
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parent directory ID. This value is empty for the internal root of the space.
	//
	// example:
	//
	// dir_parent
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// Indicates whether the directory is a read-only referenced directory. A value of false still requires creator or administrator permissions to modify the directory. The internal root is always unmodifiable.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s ListGroupDirectoriesResponseBodyDirectories) String() string {
	return dara.Prettify(s)
}

func (s ListGroupDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *ListGroupDirectoriesResponseBodyDirectories) GetDescription() *string {
	return s.Description
}

func (s *ListGroupDirectoriesResponseBodyDirectories) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListGroupDirectoriesResponseBodyDirectories) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *ListGroupDirectoriesResponseBodyDirectories) GetName() *string {
	return s.Name
}

func (s *ListGroupDirectoriesResponseBodyDirectories) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *ListGroupDirectoriesResponseBodyDirectories) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ListGroupDirectoriesResponseBodyDirectories) SetDescription(v string) *ListGroupDirectoriesResponseBodyDirectories {
	s.Description = &v
	return s
}

func (s *ListGroupDirectoriesResponseBodyDirectories) SetDirectoryId(v string) *ListGroupDirectoriesResponseBodyDirectories {
	s.DirectoryId = &v
	return s
}

func (s *ListGroupDirectoriesResponseBodyDirectories) SetDirectoryType(v string) *ListGroupDirectoriesResponseBodyDirectories {
	s.DirectoryType = &v
	return s
}

func (s *ListGroupDirectoriesResponseBodyDirectories) SetName(v string) *ListGroupDirectoriesResponseBodyDirectories {
	s.Name = &v
	return s
}

func (s *ListGroupDirectoriesResponseBodyDirectories) SetParentDirectoryId(v string) *ListGroupDirectoriesResponseBodyDirectories {
	s.ParentDirectoryId = &v
	return s
}

func (s *ListGroupDirectoriesResponseBodyDirectories) SetReadOnly(v bool) *ListGroupDirectoriesResponseBodyDirectories {
	s.ReadOnly = &v
	return s
}

func (s *ListGroupDirectoriesResponseBodyDirectories) Validate() error {
	return dara.Validate(s)
}
