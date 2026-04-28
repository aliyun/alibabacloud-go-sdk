// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressGroup interface {
	dara.Model
	String() string
	GoString() string
	SetAddressDetail(v *Address) *AddressGroup
	GetAddressDetail() *Address
	SetCount(v int64) *AddressGroup
	GetCount() *int64
	SetCoverFileId(v string) *AddressGroup
	GetCoverFileId() *string
	SetCoverUrl(v string) *AddressGroup
	GetCoverUrl() *string
	SetLocation(v string) *AddressGroup
	GetLocation() *string
	SetName(v string) *AddressGroup
	GetName() *string
}

type AddressGroup struct {
	// The information about the site.
	AddressDetail *Address `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	// The number of files in the site group.
	//
	// example:
	//
	// 10
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The ID of the cover image of the site.
	CoverFileId *string `json:"cover_file_id,omitempty" xml:"cover_file_id,omitempty"`
	// The URL of the cover image of the site.
	CoverUrl *string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// The latitude and longitude of the site.
	//
	// example:
	//
	// 30.12231,120.1212
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// The name of the site.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s AddressGroup) String() string {
	return dara.Prettify(s)
}

func (s AddressGroup) GoString() string {
	return s.String()
}

func (s *AddressGroup) GetAddressDetail() *Address {
	return s.AddressDetail
}

func (s *AddressGroup) GetCount() *int64 {
	return s.Count
}

func (s *AddressGroup) GetCoverFileId() *string {
	return s.CoverFileId
}

func (s *AddressGroup) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *AddressGroup) GetLocation() *string {
	return s.Location
}

func (s *AddressGroup) GetName() *string {
	return s.Name
}

func (s *AddressGroup) SetAddressDetail(v *Address) *AddressGroup {
	s.AddressDetail = v
	return s
}

func (s *AddressGroup) SetCount(v int64) *AddressGroup {
	s.Count = &v
	return s
}

func (s *AddressGroup) SetCoverFileId(v string) *AddressGroup {
	s.CoverFileId = &v
	return s
}

func (s *AddressGroup) SetCoverUrl(v string) *AddressGroup {
	s.CoverUrl = &v
	return s
}

func (s *AddressGroup) SetLocation(v string) *AddressGroup {
	s.Location = &v
	return s
}

func (s *AddressGroup) SetName(v string) *AddressGroup {
	s.Name = &v
	return s
}

func (s *AddressGroup) Validate() error {
	if s.AddressDetail != nil {
		if err := s.AddressDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
