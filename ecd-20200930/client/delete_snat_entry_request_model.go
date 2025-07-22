// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteSnatEntryRequest
	GetRegionId() *string
	SetSnatEntryId(v string) *DeleteSnatEntryRequest
	GetSnatEntryId() *string
	SetSnatTableId(v string) *DeleteSnatEntryRequest
	GetSnatTableId() *string
}

type DeleteSnatEntryRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// This parameter is required.
	SnatTableId *string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty"`
}

func (s DeleteSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnatEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSnatEntryRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DeleteSnatEntryRequest) GetSnatTableId() *string {
	return s.SnatTableId
}

func (s *DeleteSnatEntryRequest) SetRegionId(v string) *DeleteSnatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetSnatEntryId(v string) *DeleteSnatEntryRequest {
	s.SnatEntryId = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetSnatTableId(v string) *DeleteSnatEntryRequest {
	s.SnatTableId = &v
	return s
}

func (s *DeleteSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
