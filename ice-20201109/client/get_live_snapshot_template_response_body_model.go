// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveSnapshotTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetLiveSnapshotTemplateResponseBody
	GetCreateTime() *string
	SetLastModified(v string) *GetLiveSnapshotTemplateResponseBody
	GetLastModified() *string
	SetOverwriteFormat(v string) *GetLiveSnapshotTemplateResponseBody
	GetOverwriteFormat() *string
	SetRequestId(v string) *GetLiveSnapshotTemplateResponseBody
	GetRequestId() *string
	SetSequenceFormat(v string) *GetLiveSnapshotTemplateResponseBody
	GetSequenceFormat() *string
	SetTemplateId(v string) *GetLiveSnapshotTemplateResponseBody
	GetTemplateId() *string
	SetTemplateName(v string) *GetLiveSnapshotTemplateResponseBody
	GetTemplateName() *string
	SetTimeInterval(v int32) *GetLiveSnapshotTemplateResponseBody
	GetTimeInterval() *int32
	SetType(v string) *GetLiveSnapshotTemplateResponseBody
	GetType() *string
}

type GetLiveSnapshotTemplateResponseBody struct {
	// The time when the configuration was modified.
	//
	// example:
	//
	// 2022-02-02T22:22:22Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2022-02-02T22:22:22Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The naming format of the snapshot captured in overwrite mode.
	//
	// example:
	//
	// snapshot/{JobId}.jpg
	OverwriteFormat *string `json:"OverwriteFormat,omitempty" xml:"OverwriteFormat,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The naming format of the snapshot captured in time series mode.
	//
	// example:
	//
	// snapshot/{JobId}/{UnixTimestamp}.jpg
	SequenceFormat *string `json:"SequenceFormat,omitempty" xml:"SequenceFormat,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The interval between two adjacent snapshots.
	//
	// example:
	//
	// 5
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// The type of the template.
	//
	// Valid values:
	//
	// 	- system
	//
	// 	- custom
	//
	// example:
	//
	// custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLiveSnapshotTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveSnapshotTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveSnapshotTemplateResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLiveSnapshotTemplateResponseBody) GetLastModified() *string {
	return s.LastModified
}

func (s *GetLiveSnapshotTemplateResponseBody) GetOverwriteFormat() *string {
	return s.OverwriteFormat
}

func (s *GetLiveSnapshotTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLiveSnapshotTemplateResponseBody) GetSequenceFormat() *string {
	return s.SequenceFormat
}

func (s *GetLiveSnapshotTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLiveSnapshotTemplateResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetLiveSnapshotTemplateResponseBody) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *GetLiveSnapshotTemplateResponseBody) GetType() *string {
	return s.Type
}

func (s *GetLiveSnapshotTemplateResponseBody) SetCreateTime(v string) *GetLiveSnapshotTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLiveSnapshotTemplateResponseBody) SetLastModified(v string) *GetLiveSnapshotTemplateResponseBody {
	s.LastModified = &v
	return s
}

func (s *GetLiveSnapshotTemplateResponseBody) SetOverwriteFormat(v string) *GetLiveSnapshotTemplateResponseBody {
	s.OverwriteFormat = &v
	return s
}

func (s *GetLiveSnapshotTemplateResponseBody) SetRequestId(v string) *GetLiveSnapshotTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveSnapshotTemplateResponseBody) SetSequenceFormat(v string) *GetLiveSnapshotTemplateResponseBody {
	s.SequenceFormat = &v
	return s
}

func (s *GetLiveSnapshotTemplateResponseBody) SetTemplateId(v string) *GetLiveSnapshotTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetLiveSnapshotTemplateResponseBody) SetTemplateName(v string) *GetLiveSnapshotTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetLiveSnapshotTemplateResponseBody) SetTimeInterval(v int32) *GetLiveSnapshotTemplateResponseBody {
	s.TimeInterval = &v
	return s
}

func (s *GetLiveSnapshotTemplateResponseBody) SetType(v string) *GetLiveSnapshotTemplateResponseBody {
	s.Type = &v
	return s
}

func (s *GetLiveSnapshotTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
