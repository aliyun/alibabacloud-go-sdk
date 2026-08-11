// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveSnapshotTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOverwriteFormat(v string) *CreateLiveSnapshotTemplateRequest
	GetOverwriteFormat() *string
	SetSequenceFormat(v string) *CreateLiveSnapshotTemplateRequest
	GetSequenceFormat() *string
	SetTemplateName(v string) *CreateLiveSnapshotTemplateRequest
	GetTemplateName() *string
	SetTimeInterval(v int32) *CreateLiveSnapshotTemplateRequest
	GetTimeInterval() *int32
}

type CreateLiveSnapshotTemplateRequest struct {
	// The overwrite snapshot file format.
	//
	// - The value cannot start with "/". Only the .jpg suffix is supported.
	//
	// - Maximum length: 255.
	//
	// - Supported placeholder: {JobId}: snapshot task ID.
	//
	// - The placeholders {UnixTimestamp}, {Sequence}, and {Date} are not allowed.
	//
	// - At least one of the overwrite snapshot format or sequence snapshot format must be specified.
	//
	// example:
	//
	// snapshot/{JobId}.jpg
	OverwriteFormat *string `json:"OverwriteFormat,omitempty" xml:"OverwriteFormat,omitempty"`
	// The sequence snapshot file format.
	//
	// - The value cannot start with "/". Only the .jpg suffix is supported.
	//
	// - Maximum length: 255.
	//
	// - Supported placeholders: {JobId}: snapshot task ID, {Date}: snapshot date, {UnixTimestamp}: timestamp, {Sequence}: serial number. At least one of {UnixTimestamp} or {Sequence} must be specified.
	//
	// - At least one of the overwrite snapshot format or sequence snapshot format must be specified.
	//
	// example:
	//
	// snapshot/{JobId}/{UnixTimestamp}.jpg
	SequenceFormat *string `json:"SequenceFormat,omitempty" xml:"SequenceFormat,omitempty"`
	// The template name.
	//
	// - Maximum length: 128.
	//
	// This parameter is required.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The snapshot time interval. Unit: seconds.
	//
	// - Valid values: 5 to 3600.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
}

func (s CreateLiveSnapshotTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveSnapshotTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveSnapshotTemplateRequest) GetOverwriteFormat() *string {
	return s.OverwriteFormat
}

func (s *CreateLiveSnapshotTemplateRequest) GetSequenceFormat() *string {
	return s.SequenceFormat
}

func (s *CreateLiveSnapshotTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateLiveSnapshotTemplateRequest) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *CreateLiveSnapshotTemplateRequest) SetOverwriteFormat(v string) *CreateLiveSnapshotTemplateRequest {
	s.OverwriteFormat = &v
	return s
}

func (s *CreateLiveSnapshotTemplateRequest) SetSequenceFormat(v string) *CreateLiveSnapshotTemplateRequest {
	s.SequenceFormat = &v
	return s
}

func (s *CreateLiveSnapshotTemplateRequest) SetTemplateName(v string) *CreateLiveSnapshotTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateLiveSnapshotTemplateRequest) SetTimeInterval(v int32) *CreateLiveSnapshotTemplateRequest {
	s.TimeInterval = &v
	return s
}

func (s *CreateLiveSnapshotTemplateRequest) Validate() error {
	return dara.Validate(s)
}
