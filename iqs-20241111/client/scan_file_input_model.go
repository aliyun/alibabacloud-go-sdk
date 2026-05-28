// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanFileInput interface {
	dara.Model
	String() string
	GoString() string
	SetImageBase64(v string) *ScanFileInput
	GetImageBase64() *string
	SetImageUrl(v string) *ScanFileInput
	GetImageUrl() *string
	SetScanFileInputConfig(v *ScanFileInputConfig) *ScanFileInput
	GetScanFileInputConfig() *ScanFileInputConfig
}

type ScanFileInput struct {
	// example:
	//
	// wrwqr
	ImageBase64 *string `json:"imageBase64,omitempty" xml:"imageBase64,omitempty"`
	// example:
	//
	// https://www.1241.png
	ImageUrl            *string              `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	ScanFileInputConfig *ScanFileInputConfig `json:"scanFileInputConfig,omitempty" xml:"scanFileInputConfig,omitempty"`
}

func (s ScanFileInput) String() string {
	return dara.Prettify(s)
}

func (s ScanFileInput) GoString() string {
	return s.String()
}

func (s *ScanFileInput) GetImageBase64() *string {
	return s.ImageBase64
}

func (s *ScanFileInput) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ScanFileInput) GetScanFileInputConfig() *ScanFileInputConfig {
	return s.ScanFileInputConfig
}

func (s *ScanFileInput) SetImageBase64(v string) *ScanFileInput {
	s.ImageBase64 = &v
	return s
}

func (s *ScanFileInput) SetImageUrl(v string) *ScanFileInput {
	s.ImageUrl = &v
	return s
}

func (s *ScanFileInput) SetScanFileInputConfig(v *ScanFileInputConfig) *ScanFileInput {
	s.ScanFileInputConfig = v
	return s
}

func (s *ScanFileInput) Validate() error {
	if s.ScanFileInputConfig != nil {
		if err := s.ScanFileInputConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
