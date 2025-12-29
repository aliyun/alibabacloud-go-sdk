// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportHotelConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ImportHotelConfigShrinkRequest
	GetHotelId() *string
	SetImportHotelConfigShrink(v string) *ImportHotelConfigShrinkRequest
	GetImportHotelConfigShrink() *string
}

type ImportHotelConfigShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	ImportHotelConfigShrink *string `json:"ImportHotelConfig,omitempty" xml:"ImportHotelConfig,omitempty"`
}

func (s ImportHotelConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportHotelConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ImportHotelConfigShrinkRequest) GetImportHotelConfigShrink() *string {
	return s.ImportHotelConfigShrink
}

func (s *ImportHotelConfigShrinkRequest) SetHotelId(v string) *ImportHotelConfigShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ImportHotelConfigShrinkRequest) SetImportHotelConfigShrink(v string) *ImportHotelConfigShrinkRequest {
	s.ImportHotelConfigShrink = &v
	return s
}

func (s *ImportHotelConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
