// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDevicesFromCSVRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *AddDevicesFromCSVRequest
	GetFileName() *string
	SetFileType(v int32) *AddDevicesFromCSVRequest
	GetFileType() *int32
	SetSeatCol(v int32) *AddDevicesFromCSVRequest
	GetSeatCol() *int32
	SetSiteId(v string) *AddDevicesFromCSVRequest
	GetSiteId() *string
	SetSiteName(v string) *AddDevicesFromCSVRequest
	GetSiteName() *string
}

type AddDevicesFromCSVRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileType *int32  `json:"FileType,omitempty" xml:"FileType,omitempty"`
	SeatCol  *int32  `json:"SeatCol,omitempty" xml:"SeatCol,omitempty"`
	SiteId   *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s AddDevicesFromCSVRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDevicesFromCSVRequest) GoString() string {
	return s.String()
}

func (s *AddDevicesFromCSVRequest) GetFileName() *string {
	return s.FileName
}

func (s *AddDevicesFromCSVRequest) GetFileType() *int32 {
	return s.FileType
}

func (s *AddDevicesFromCSVRequest) GetSeatCol() *int32 {
	return s.SeatCol
}

func (s *AddDevicesFromCSVRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *AddDevicesFromCSVRequest) GetSiteName() *string {
	return s.SiteName
}

func (s *AddDevicesFromCSVRequest) SetFileName(v string) *AddDevicesFromCSVRequest {
	s.FileName = &v
	return s
}

func (s *AddDevicesFromCSVRequest) SetFileType(v int32) *AddDevicesFromCSVRequest {
	s.FileType = &v
	return s
}

func (s *AddDevicesFromCSVRequest) SetSeatCol(v int32) *AddDevicesFromCSVRequest {
	s.SeatCol = &v
	return s
}

func (s *AddDevicesFromCSVRequest) SetSiteId(v string) *AddDevicesFromCSVRequest {
	s.SiteId = &v
	return s
}

func (s *AddDevicesFromCSVRequest) SetSiteName(v string) *AddDevicesFromCSVRequest {
	s.SiteName = &v
	return s
}

func (s *AddDevicesFromCSVRequest) Validate() error {
	return dara.Validate(s)
}
