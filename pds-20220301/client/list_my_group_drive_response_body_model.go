// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyGroupDriveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Drive) *ListMyGroupDriveResponseBody
	GetItems() []*Drive
	SetNextMarker(v string) *ListMyGroupDriveResponseBody
	GetNextMarker() *string
	SetRootGroupDrive(v *Drive) *ListMyGroupDriveResponseBody
	GetRootGroupDrive() *Drive
}

type ListMyGroupDriveResponseBody struct {
	// The information about the team drives.
	Items []*Drive `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker     *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	RootGroupDrive *Drive  `json:"root_group_drive,omitempty" xml:"root_group_drive,omitempty"`
}

func (s ListMyGroupDriveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMyGroupDriveResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyGroupDriveResponseBody) GetItems() []*Drive {
	return s.Items
}

func (s *ListMyGroupDriveResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListMyGroupDriveResponseBody) GetRootGroupDrive() *Drive {
	return s.RootGroupDrive
}

func (s *ListMyGroupDriveResponseBody) SetItems(v []*Drive) *ListMyGroupDriveResponseBody {
	s.Items = v
	return s
}

func (s *ListMyGroupDriveResponseBody) SetNextMarker(v string) *ListMyGroupDriveResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListMyGroupDriveResponseBody) SetRootGroupDrive(v *Drive) *ListMyGroupDriveResponseBody {
	s.RootGroupDrive = v
	return s
}

func (s *ListMyGroupDriveResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RootGroupDrive != nil {
		if err := s.RootGroupDrive.Validate(); err != nil {
			return err
		}
	}
	return nil
}
