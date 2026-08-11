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
	// The template name.
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
	// The cycle recording duration. Unit: seconds. If this parameter is not specified, the default value is 6 hours.
	//
	// > - If a live stream is interrupted during a recording cycle but resumes within 3 minutes, the recording continues in the same recording file.
	//
	// - A live stream must be interrupted for more than 3 minutes before the last recording file is generated. If you need to modify the default 3-minute interruption time, submit a ticket.
	//
	// example:
	//
	// 3600
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The format.
	//
	// >If you select the m3u8 format, you must also set the request parameters SliceOssObjectPrefix and SliceDuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The name of the recording file stored in OSS.
	//
	// - The file name must be less than 256 bytes and supports variable matching, including {JobId}, {Sequence}, {StartTime}, {EndTime}, {EscapedStartTime}, and {EscapedEndTime}.
	//
	// - The parameter value must contain {StartTime} or {EscapedStartTime} and {EndTime} or {EscapedEndTime}.
	//
	// example:
	//
	// record/{JobId}/{Sequence}_{EscapedStartTime}_{EscapedEndTime}
	OssObjectPrefix *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	// The duration of a single slice. Unit: seconds.
	//
	// >This parameter takes effect only when Format is set to m3u8.
	//
	// If this parameter is not specified, the default value is 30 seconds. Valid values: 5 to 30.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// The name of the TS slice.
	//
	// >This parameter is required only when Format is set to m3u8.
	//
	// - The default slice duration is 30 seconds. The name must be less than 256 bytes and supports variable matching, including {JobId}, {UnixTimestamp}, and {Sequence}.
	//
	// - The parameter value must contain the {UnixTimestamp} and {Sequence} variables.
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
