// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVirusScanAdditionalListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalLists(v []*AddVirusScanAdditionalListsRequestAdditionalLists) *AddVirusScanAdditionalListsRequest
	GetAdditionalLists() []*AddVirusScanAdditionalListsRequestAdditionalLists
	SetDevType(v string) *AddVirusScanAdditionalListsRequest
	GetDevType() *string
}

type AddVirusScanAdditionalListsRequest struct {
	// The list of entries to append. At least one entry is required.
	AdditionalLists []*AddVirusScanAdditionalListsRequestAdditionalLists `json:"AdditionalLists,omitempty" xml:"AdditionalLists,omitempty" type:"Repeated"`
	// The operating system type for which the list takes effect. Valid values:
	//
	// - **windows**: Windows.
	//
	// - **macOS**: macOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// windows
	DevType *string `json:"DevType,omitempty" xml:"DevType,omitempty"`
}

func (s AddVirusScanAdditionalListsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddVirusScanAdditionalListsRequest) GoString() string {
	return s.String()
}

func (s *AddVirusScanAdditionalListsRequest) GetAdditionalLists() []*AddVirusScanAdditionalListsRequestAdditionalLists {
	return s.AdditionalLists
}

func (s *AddVirusScanAdditionalListsRequest) GetDevType() *string {
	return s.DevType
}

func (s *AddVirusScanAdditionalListsRequest) SetAdditionalLists(v []*AddVirusScanAdditionalListsRequestAdditionalLists) *AddVirusScanAdditionalListsRequest {
	s.AdditionalLists = v
	return s
}

func (s *AddVirusScanAdditionalListsRequest) SetDevType(v string) *AddVirusScanAdditionalListsRequest {
	s.DevType = &v
	return s
}

func (s *AddVirusScanAdditionalListsRequest) Validate() error {
	if s.AdditionalLists != nil {
		for _, item := range s.AdditionalLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddVirusScanAdditionalListsRequestAdditionalLists struct {
	// The matching dimension of the list entry. Valid values:
	//
	// - **FileSuffix**: matches by file name extension.
	//
	// - **FileName**: matches by file name.
	//
	// - **FolderName**: matches by folder name.
	//
	// - **FilePath**: matches by file path.
	//
	// - **FileMd5**: matches by file MD5 value.
	//
	// example:
	//
	// FileSuffix
	AdditionalType *string `json:"AdditionalType,omitempty" xml:"AdditionalType,omitempty"`
	// The content of the list entry. The value cannot exceed 255 characters. The meaning is determined by AdditionalType: when AdditionalType is set to FileSuffix, specify a file name extension. When set to FileName, specify a file name. When set to FolderName, specify a folder name. When set to FilePath, specify a file path. When set to FileMd5, specify the MD5 value of a file.
	//
	// example:
	//
	// .tmp
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The list type. Valid values:
	//
	// - **Blacklist**: blacklist. Files that match are directly identified as virus files.
	//
	// - **Whitelist**: whitelist. Files that match are excluded from virus detection.
	//
	// example:
	//
	// Whitelist
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
}

func (s AddVirusScanAdditionalListsRequestAdditionalLists) String() string {
	return dara.Prettify(s)
}

func (s AddVirusScanAdditionalListsRequestAdditionalLists) GoString() string {
	return s.String()
}

func (s *AddVirusScanAdditionalListsRequestAdditionalLists) GetAdditionalType() *string {
	return s.AdditionalType
}

func (s *AddVirusScanAdditionalListsRequestAdditionalLists) GetDetail() *string {
	return s.Detail
}

func (s *AddVirusScanAdditionalListsRequestAdditionalLists) GetListType() *string {
	return s.ListType
}

func (s *AddVirusScanAdditionalListsRequestAdditionalLists) SetAdditionalType(v string) *AddVirusScanAdditionalListsRequestAdditionalLists {
	s.AdditionalType = &v
	return s
}

func (s *AddVirusScanAdditionalListsRequestAdditionalLists) SetDetail(v string) *AddVirusScanAdditionalListsRequestAdditionalLists {
	s.Detail = &v
	return s
}

func (s *AddVirusScanAdditionalListsRequestAdditionalLists) SetListType(v string) *AddVirusScanAdditionalListsRequestAdditionalLists {
	s.ListType = &v
	return s
}

func (s *AddVirusScanAdditionalListsRequestAdditionalLists) Validate() error {
	return dara.Validate(s)
}
