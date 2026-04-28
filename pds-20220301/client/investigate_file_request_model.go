// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvestigateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveFileIds(v []*InvestigateFileRequestDriveFileIds) *InvestigateFileRequest
	GetDriveFileIds() []*InvestigateFileRequestDriveFileIds
}

type InvestigateFileRequest struct {
	// This parameter is required.
	DriveFileIds []*InvestigateFileRequestDriveFileIds `json:"drive_file_ids,omitempty" xml:"drive_file_ids,omitempty" type:"Repeated"`
}

func (s InvestigateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s InvestigateFileRequest) GoString() string {
	return s.String()
}

func (s *InvestigateFileRequest) GetDriveFileIds() []*InvestigateFileRequestDriveFileIds {
	return s.DriveFileIds
}

func (s *InvestigateFileRequest) SetDriveFileIds(v []*InvestigateFileRequestDriveFileIds) *InvestigateFileRequest {
	s.DriveFileIds = v
	return s
}

func (s *InvestigateFileRequest) Validate() error {
	if s.DriveFileIds != nil {
		for _, item := range s.DriveFileIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InvestigateFileRequestDriveFileIds struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s InvestigateFileRequestDriveFileIds) String() string {
	return dara.Prettify(s)
}

func (s InvestigateFileRequestDriveFileIds) GoString() string {
	return s.String()
}

func (s *InvestigateFileRequestDriveFileIds) GetDriveId() *string {
	return s.DriveId
}

func (s *InvestigateFileRequestDriveFileIds) GetFileId() *string {
	return s.FileId
}

func (s *InvestigateFileRequestDriveFileIds) SetDriveId(v string) *InvestigateFileRequestDriveFileIds {
	s.DriveId = &v
	return s
}

func (s *InvestigateFileRequestDriveFileIds) SetFileId(v string) *InvestigateFileRequestDriveFileIds {
	s.FileId = &v
	return s
}

func (s *InvestigateFileRequestDriveFileIds) Validate() error {
	return dara.Validate(s)
}
