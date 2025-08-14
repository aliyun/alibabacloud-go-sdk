// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveSnapshotTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOverwriteFormat(v string) *UpdateLiveSnapshotTemplateRequest
	GetOverwriteFormat() *string
	SetSequenceFormat(v string) *UpdateLiveSnapshotTemplateRequest
	GetSequenceFormat() *string
	SetTemplateId(v string) *UpdateLiveSnapshotTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *UpdateLiveSnapshotTemplateRequest
	GetTemplateName() *string
	SetTimeInterval(v int32) *UpdateLiveSnapshotTemplateRequest
	GetTimeInterval() *int32
}

type UpdateLiveSnapshotTemplateRequest struct {
	// The naming format of the snapshot captured in overwrite mode.
	//
	// 	- The value cannot start with a forward slash (/). Only the suffix .jpg is supported.
	//
	// 	- The value cannot exceed 255 characters in length.
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
	// 	- The value cannot exceed 255 characters in length.
	//
	// 	- The {JobId}, {Date}, {UnixTimestamp}, and {Sequence} placeholders are supported. {JobId} specifies the ID of the snapshot job. {Date} specifies the date on which the snapshot is captured. {UnixTimestamp} specifies the timestamp of the snapshot. {Sequence} specifies the sequence number of the snapshot. You must specify at least one of the {UnixTimestamp} and {Sequence} placeholders.
	//
	// 	- You must specify at least one of the OverwriteFormat and SequenceFormat parameters.
	//
	// example:
	//
	// snapshot/{JobId}/{UnixTimestamp}.jpg
	SequenceFormat *string `json:"SequenceFormat,omitempty" xml:"SequenceFormat,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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

func (s UpdateLiveSnapshotTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveSnapshotTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveSnapshotTemplateRequest) GetOverwriteFormat() *string {
	return s.OverwriteFormat
}

func (s *UpdateLiveSnapshotTemplateRequest) GetSequenceFormat() *string {
	return s.SequenceFormat
}

func (s *UpdateLiveSnapshotTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateLiveSnapshotTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateLiveSnapshotTemplateRequest) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *UpdateLiveSnapshotTemplateRequest) SetOverwriteFormat(v string) *UpdateLiveSnapshotTemplateRequest {
	s.OverwriteFormat = &v
	return s
}

func (s *UpdateLiveSnapshotTemplateRequest) SetSequenceFormat(v string) *UpdateLiveSnapshotTemplateRequest {
	s.SequenceFormat = &v
	return s
}

func (s *UpdateLiveSnapshotTemplateRequest) SetTemplateId(v string) *UpdateLiveSnapshotTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateLiveSnapshotTemplateRequest) SetTemplateName(v string) *UpdateLiveSnapshotTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateLiveSnapshotTemplateRequest) SetTimeInterval(v int32) *UpdateLiveSnapshotTemplateRequest {
	s.TimeInterval = &v
	return s
}

func (s *UpdateLiveSnapshotTemplateRequest) Validate() error {
	return dara.Validate(s)
}
