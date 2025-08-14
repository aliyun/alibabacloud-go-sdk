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
	// The naming format of the snapshot captured in overwrite mode.
	//
	// 	- The value cannot start with a forward slash (/). Only the suffix .jpg is supported.
	//
	// 	- It cannot exceed 255 characters in length.
	//
	// 	- The {JobId} placeholder is supported. It specifies the ID of the snapshot job.
	//
	// 	- Placeholders such as {UnixTimestamp}, {Sequence}, and {Date} are not allowed.
	//
	// 	- You must specify at least one of the OverwriteFormat and SequenceFormat parameters.
	//
	// example:
	//
	// snapshot/{JobId}.jpg
	OverwriteFormat *string `json:"OverwriteFormat,omitempty" xml:"OverwriteFormat,omitempty"`
	// The naming format of the snapshot captured in time series mode.
	//
	// 	- The value cannot start with a forward slash (/). Only the suffix .jpg is supported.
	//
	// 	- It cannot exceed 255 characters in length.
	//
	// 	- The {JobId}, {Date}, {UnixTimestamp}, and {Sequence} placeholders are supported. {JobId} specifies the ID of the snapshot job. {Date} specifies the date on which the snapshot is captured. {UnixTimestamp} specifies the timestamp of the snapshot. {Sequence} specifies the sequence number of the snapshot. You must specify at least one of the {UnixTimestamp} and {Sequence} placeholders.
	//
	// 	- You must specify at least one of the OverwriteFormat and SequenceFormat parameters.
	//
	// example:
	//
	// snapshot/{JobId}/{UnixTimestamp}.jpg
	SequenceFormat *string `json:"SequenceFormat,omitempty" xml:"SequenceFormat,omitempty"`
	// The name of the template.
	//
	// 	- It cannot exceed 128 characters in length.
	//
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The interval between two adjacent snapshots. Unit: seconds.
	//
	// 	- Valid values: [5,3600].
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
