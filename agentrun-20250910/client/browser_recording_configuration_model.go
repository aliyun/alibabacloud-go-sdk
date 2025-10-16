// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowserRecordingConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *BrowserRecordingConfiguration
	GetEnabled() *bool
	SetOssLocation(v *BrowserOssLocation) *BrowserRecordingConfiguration
	GetOssLocation() *BrowserOssLocation
}

type BrowserRecordingConfiguration struct {
	Enabled     *bool               `json:"enabled,omitempty" xml:"enabled,omitempty"`
	OssLocation *BrowserOssLocation `json:"ossLocation,omitempty" xml:"ossLocation,omitempty"`
}

func (s BrowserRecordingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s BrowserRecordingConfiguration) GoString() string {
	return s.String()
}

func (s *BrowserRecordingConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *BrowserRecordingConfiguration) GetOssLocation() *BrowserOssLocation {
	return s.OssLocation
}

func (s *BrowserRecordingConfiguration) SetEnabled(v bool) *BrowserRecordingConfiguration {
	s.Enabled = &v
	return s
}

func (s *BrowserRecordingConfiguration) SetOssLocation(v *BrowserOssLocation) *BrowserRecordingConfiguration {
	s.OssLocation = v
	return s
}

func (s *BrowserRecordingConfiguration) Validate() error {
	if s.OssLocation != nil {
		if err := s.OssLocation.Validate(); err != nil {
			return err
		}
	}
	return nil
}
