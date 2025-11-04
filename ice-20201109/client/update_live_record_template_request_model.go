// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveRecordTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateLiveRecordTemplateRequest
	GetName() *string
	SetRecordFormat(v []*UpdateLiveRecordTemplateRequestRecordFormat) *UpdateLiveRecordTemplateRequest
	GetRecordFormat() []*UpdateLiveRecordTemplateRequestRecordFormat
	SetTemplateId(v string) *UpdateLiveRecordTemplateRequest
	GetTemplateId() *string
}

type UpdateLiveRecordTemplateRequest struct {
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of recording formats.
	//
	// This parameter is required.
	RecordFormat []*UpdateLiveRecordTemplateRequestRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Repeated"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 69e1f9fe-1e97-11ed-ba64-0c42a1b73d66
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateLiveRecordTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLiveRecordTemplateRequest) GetRecordFormat() []*UpdateLiveRecordTemplateRequestRecordFormat {
	return s.RecordFormat
}

func (s *UpdateLiveRecordTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateLiveRecordTemplateRequest) SetName(v string) *UpdateLiveRecordTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateLiveRecordTemplateRequest) SetRecordFormat(v []*UpdateLiveRecordTemplateRequestRecordFormat) *UpdateLiveRecordTemplateRequest {
	s.RecordFormat = v
	return s
}

func (s *UpdateLiveRecordTemplateRequest) SetTemplateId(v string) *UpdateLiveRecordTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateLiveRecordTemplateRequest) Validate() error {
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

type UpdateLiveRecordTemplateRequestRecordFormat struct {
	// The duration of the recording cycle. Unit: seconds If you do not specify this parameter, the default value 6 hours is used.
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
	// The format of recording files.
	//
	// >  If you set this parameter to m3u8, you must also specify the SliceOssObjectPrefix and SliceDuration parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The name of the recording that is stored in Object Storage Service (OSS).
	//
	// 	- The name must be less than 256 bytes in length and can contain the {JobId}, {Sequence}, {StartTime}, {EndTime}, {EscapedStartTime}, and {EscapedEndTime} variables.
	//
	// 	- The name must contain the {StartTime} and {EndTime} variables or the {EscapedStartTime} and {EscapedEndTime} variables.
	//
	// example:
	//
	// record/{JobId}/{Sequence}_{EscapedStartTime}_{EscapedEndTime}
	OssObjectPrefix *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	// The duration of a single segment. Unit: seconds
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
	// >  This parameter is required only if you set Format to m3u8. By default, the duration of a segment is 30 seconds. The segment name must be less than 256 bytes in length and can contain the {JobId}, {UnixTimestamp}, and {Sequence} variables.
	//
	// The segment name must contain the {UnixTimestamp} and {Sequence} variables.
	//
	// example:
	//
	// record/{JobId}/{UnixTimestamp}_{Sequence}
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty"`
}

func (s UpdateLiveRecordTemplateRequestRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordTemplateRequestRecordFormat) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordTemplateRequestRecordFormat) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *UpdateLiveRecordTemplateRequestRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *UpdateLiveRecordTemplateRequestRecordFormat) GetOssObjectPrefix() *string {
	return s.OssObjectPrefix
}

func (s *UpdateLiveRecordTemplateRequestRecordFormat) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *UpdateLiveRecordTemplateRequestRecordFormat) GetSliceOssObjectPrefix() *string {
	return s.SliceOssObjectPrefix
}

func (s *UpdateLiveRecordTemplateRequestRecordFormat) SetCycleDuration(v int32) *UpdateLiveRecordTemplateRequestRecordFormat {
	s.CycleDuration = &v
	return s
}

func (s *UpdateLiveRecordTemplateRequestRecordFormat) SetFormat(v string) *UpdateLiveRecordTemplateRequestRecordFormat {
	s.Format = &v
	return s
}

func (s *UpdateLiveRecordTemplateRequestRecordFormat) SetOssObjectPrefix(v string) *UpdateLiveRecordTemplateRequestRecordFormat {
	s.OssObjectPrefix = &v
	return s
}

func (s *UpdateLiveRecordTemplateRequestRecordFormat) SetSliceDuration(v int32) *UpdateLiveRecordTemplateRequestRecordFormat {
	s.SliceDuration = &v
	return s
}

func (s *UpdateLiveRecordTemplateRequestRecordFormat) SetSliceOssObjectPrefix(v string) *UpdateLiveRecordTemplateRequestRecordFormat {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *UpdateLiveRecordTemplateRequestRecordFormat) Validate() error {
	return dara.Validate(s)
}
