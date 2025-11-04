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
	// The name of the template.
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
	// The duration of the recording cycle. Unit: seconds. If you do not specify this parameter, the default value 6 hours is used.
	//
	// >
	//
	// 	- If a live stream is interrupted during a recording cycle but is resumed within 3 minutes, the stream is recorded in the same recording before and after the interruption.
	//
	// 	- If a live stream is interrupted for more than 3 minutes, a new recording is generated. To change the default stream interruption time, submit a ticket.
	//
	// example:
	//
	// 3600
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The format.
	//
	// >  If you set this parameter to m3u8, you must also specify the SliceOssObjectPrefix and SliceDuration parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The name of the recording file that is stored in Object Storage Service (OSS).
	//
	// 	- The name must be less than 256 bytes in length and can contain the {JobId}, {Sequence}, {StartTime}, {EndTime}, {EscapedStartTime}, and {EscapedEndTime} variables.
	//
	// 	- The name must contain the {StartTime} and {EndTime} variables or the {EscapedStartTime} and {EscapedEndTime} variables.
	//
	// example:
	//
	// record/{JobId}/{Sequence}_{EscapedStartTime}_{EscapedEndTime}
	OssObjectPrefix *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	// The duration of a single segment. Unit: seconds.
	//
	// >  This parameter takes effect only if you set Format to m3u8.
	//
	// If you do not specify this parameter, the default value 30 seconds is used. Valid values: 5 to 30.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// The name of the TS segment.
	//
	// >  This parameter is required only if you set Format to m3u8.
	//
	// 	- By default, the duration of a segment is 30 seconds. The segment name must be less than 256 bytes in length and can contain the {JobId}, {UnixTimestamp}, and {Sequence} variables.
	//
	// 	- The segment name must contain the {UnixTimestamp} and {Sequence} variables.
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
