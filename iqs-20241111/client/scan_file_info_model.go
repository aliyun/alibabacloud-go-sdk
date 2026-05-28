// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanFileInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAngle(v int32) *ScanFileInfo
	GetAngle() *int32
	SetHeight(v int32) *ScanFileInfo
	GetHeight() *int32
	SetImageBase64(v []byte) *ScanFileInfo
	GetImageBase64() []byte
	SetWidth(v int32) *ScanFileInfo
	GetWidth() *int32
}

type ScanFileInfo struct {
	// example:
	//
	// 100
	Angle *int32 `json:"angle,omitempty" xml:"angle,omitempty"`
	// example:
	//
	// 100
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// erwre
	ImageBase64 []byte `json:"imageBase64,omitempty" xml:"imageBase64,omitempty"`
	// example:
	//
	// 100
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s ScanFileInfo) String() string {
	return dara.Prettify(s)
}

func (s ScanFileInfo) GoString() string {
	return s.String()
}

func (s *ScanFileInfo) GetAngle() *int32 {
	return s.Angle
}

func (s *ScanFileInfo) GetHeight() *int32 {
	return s.Height
}

func (s *ScanFileInfo) GetImageBase64() []byte {
	return s.ImageBase64
}

func (s *ScanFileInfo) GetWidth() *int32 {
	return s.Width
}

func (s *ScanFileInfo) SetAngle(v int32) *ScanFileInfo {
	s.Angle = &v
	return s
}

func (s *ScanFileInfo) SetHeight(v int32) *ScanFileInfo {
	s.Height = &v
	return s
}

func (s *ScanFileInfo) SetImageBase64(v []byte) *ScanFileInfo {
	s.ImageBase64 = v
	return s
}

func (s *ScanFileInfo) SetWidth(v int32) *ScanFileInfo {
	s.Width = &v
	return s
}

func (s *ScanFileInfo) Validate() error {
	return dara.Validate(s)
}
