// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveRecordTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateLiveRecordTemplateRequest
	GetName() *string
	SetRecordFormat(v []*CreateLiveRecordTemplateRequestRecordFormat) *CreateLiveRecordTemplateRequest
	GetRecordFormat() []*CreateLiveRecordTemplateRequestRecordFormat
}

type CreateLiveRecordTemplateRequest struct {
	// The name of the Live Record Template.
	//
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of recording formats.
	//
	// This parameter is required.
	RecordFormat []*CreateLiveRecordTemplateRequestRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Repeated"`
}

func (s CreateLiveRecordTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRecordTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateLiveRecordTemplateRequest) GetRecordFormat() []*CreateLiveRecordTemplateRequestRecordFormat {
	return s.RecordFormat
}

func (s *CreateLiveRecordTemplateRequest) SetName(v string) *CreateLiveRecordTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateLiveRecordTemplateRequest) SetRecordFormat(v []*CreateLiveRecordTemplateRequestRecordFormat) *CreateLiveRecordTemplateRequest {
	s.RecordFormat = v
	return s
}

func (s *CreateLiveRecordTemplateRequest) Validate() error {
	if s.RecordFormat != nil {
		for _, item := range s.RecordFormat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLiveRecordTemplateRequestRecordFormat struct {
	// The duration of a recording cycle in seconds. If you omit this parameter, it defaults to 6 hours.
	//
	// > - If a stream interruption during a recording cycle lasts less than 3 minutes, the recording continues in the same Recording File.
	//
	// - A Recording File is finalized only after a stream interruption lasts for more than 3 minutes. To change this default 3-minute threshold, submit a ticket.
	//
	// example:
	//
	// 3600
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The recording format.
	//
	// > If you set this parameter to `m3u8`, you must also specify the `SliceOssObjectPrefix` and `SliceDuration` parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The name of the Recording File stored in Object Storage Service (OSS).
	//
	// - The file name must be less than 256 bytes and supports the following variables: {JobId}, {Sequence}, {StartTime}, {EndTime}, {EscapedStartTime}, and {EscapedEndTime}.
	//
	// - The value must include either the {StartTime} or {EscapedStartTime} variable and either the {EndTime} or {EscapedEndTime} variable.
	//
	// example:
	//
	// record/{JobId}/{Sequence}_{EscapedStartTime}_{EscapedEndTime}
	OssObjectPrefix *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	// The duration of each slice in seconds.
	//
	// > This parameter is valid only when `Format` is set to `m3u8`.
	//
	// The default value is 30. The value must be an integer from 5 to 30.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// The name of the TS slice.
	//
	// > This parameter is required only when `Format` is set to `m3u8`.
	//
	// - The file name must be less than 256 bytes and supports the following variables: {JobId}, {UnixTimestamp}, and {Sequence}.
	//
	// - The value must include the {UnixTimestamp} and {Sequence} variables.
	//
	// example:
	//
	// record/{JobId}/{UnixTimestamp}_{Sequence}
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty"`
}

func (s CreateLiveRecordTemplateRequestRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveRecordTemplateRequestRecordFormat) GoString() string {
	return s.String()
}

func (s *CreateLiveRecordTemplateRequestRecordFormat) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *CreateLiveRecordTemplateRequestRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *CreateLiveRecordTemplateRequestRecordFormat) GetOssObjectPrefix() *string {
	return s.OssObjectPrefix
}

func (s *CreateLiveRecordTemplateRequestRecordFormat) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *CreateLiveRecordTemplateRequestRecordFormat) GetSliceOssObjectPrefix() *string {
	return s.SliceOssObjectPrefix
}

func (s *CreateLiveRecordTemplateRequestRecordFormat) SetCycleDuration(v int32) *CreateLiveRecordTemplateRequestRecordFormat {
	s.CycleDuration = &v
	return s
}

func (s *CreateLiveRecordTemplateRequestRecordFormat) SetFormat(v string) *CreateLiveRecordTemplateRequestRecordFormat {
	s.Format = &v
	return s
}

func (s *CreateLiveRecordTemplateRequestRecordFormat) SetOssObjectPrefix(v string) *CreateLiveRecordTemplateRequestRecordFormat {
	s.OssObjectPrefix = &v
	return s
}

func (s *CreateLiveRecordTemplateRequestRecordFormat) SetSliceDuration(v int32) *CreateLiveRecordTemplateRequestRecordFormat {
	s.SliceDuration = &v
	return s
}

func (s *CreateLiveRecordTemplateRequestRecordFormat) SetSliceOssObjectPrefix(v string) *CreateLiveRecordTemplateRequestRecordFormat {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *CreateLiveRecordTemplateRequestRecordFormat) Validate() error {
	return dara.Validate(s)
}
