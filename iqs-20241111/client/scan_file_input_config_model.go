// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanFileInputConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCrop(v string) *ScanFileInputConfig
	GetAutoCrop() *string
	SetAutoRotate(v string) *ScanFileInputConfig
	GetAutoRotate() *string
}

type ScanFileInputConfig struct {
	AutoCrop   *string `json:"autoCrop,omitempty" xml:"autoCrop,omitempty"`
	AutoRotate *string `json:"autoRotate,omitempty" xml:"autoRotate,omitempty"`
}

func (s ScanFileInputConfig) String() string {
	return dara.Prettify(s)
}

func (s ScanFileInputConfig) GoString() string {
	return s.String()
}

func (s *ScanFileInputConfig) GetAutoCrop() *string {
	return s.AutoCrop
}

func (s *ScanFileInputConfig) GetAutoRotate() *string {
	return s.AutoRotate
}

func (s *ScanFileInputConfig) SetAutoCrop(v string) *ScanFileInputConfig {
	s.AutoCrop = &v
	return s
}

func (s *ScanFileInputConfig) SetAutoRotate(v string) *ScanFileInputConfig {
	s.AutoRotate = &v
	return s
}

func (s *ScanFileInputConfig) Validate() error {
	return dara.Validate(s)
}
