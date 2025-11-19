// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordingConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *RecordingConfiguration
	GetEnabled() *bool
	SetOssLocation(v *OssConfiguration) *RecordingConfiguration
	GetOssLocation() *OssConfiguration
}

type RecordingConfiguration struct {
	// 是否启用录制
	//
	// This parameter is required.
	Enabled     *bool             `json:"enabled,omitempty" xml:"enabled,omitempty"`
	OssLocation *OssConfiguration `json:"ossLocation,omitempty" xml:"ossLocation,omitempty"`
}

func (s RecordingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s RecordingConfiguration) GoString() string {
	return s.String()
}

func (s *RecordingConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *RecordingConfiguration) GetOssLocation() *OssConfiguration {
	return s.OssLocation
}

func (s *RecordingConfiguration) SetEnabled(v bool) *RecordingConfiguration {
	s.Enabled = &v
	return s
}

func (s *RecordingConfiguration) SetOssLocation(v *OssConfiguration) *RecordingConfiguration {
	s.OssLocation = v
	return s
}

func (s *RecordingConfiguration) Validate() error {
	if s.OssLocation != nil {
		if err := s.OssLocation.Validate(); err != nil {
			return err
		}
	}
	return nil
}
