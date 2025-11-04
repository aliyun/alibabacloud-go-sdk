// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveRecordTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordTemplate(v *GetLiveRecordTemplateResponseBodyRecordTemplate) *GetLiveRecordTemplateResponseBody
	GetRecordTemplate() *GetLiveRecordTemplateResponseBodyRecordTemplate
	SetRequestId(v string) *GetLiveRecordTemplateResponseBody
	GetRequestId() *string
}

type GetLiveRecordTemplateResponseBody struct {
	// The recording template.
	RecordTemplate *GetLiveRecordTemplateResponseBodyRecordTemplate `json:"RecordTemplate,omitempty" xml:"RecordTemplate,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C892855F-95DF-50D6-A28C-279ABDB76810
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLiveRecordTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveRecordTemplateResponseBody) GetRecordTemplate() *GetLiveRecordTemplateResponseBodyRecordTemplate {
	return s.RecordTemplate
}

func (s *GetLiveRecordTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLiveRecordTemplateResponseBody) SetRecordTemplate(v *GetLiveRecordTemplateResponseBodyRecordTemplate) *GetLiveRecordTemplateResponseBody {
	s.RecordTemplate = v
	return s
}

func (s *GetLiveRecordTemplateResponseBody) SetRequestId(v string) *GetLiveRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveRecordTemplateResponseBody) Validate() error {
	if s.RecordTemplate != nil {
		if err := s.RecordTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLiveRecordTemplateResponseBodyRecordTemplate struct {
	// The time when the job was created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-07-20T03:26:36Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the template was last modified.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-07-20T03:26:36Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The template name.
	//
	// example:
	//
	// test template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of recording formats.
	RecordFormatList []*GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList `json:"RecordFormatList,omitempty" xml:"RecordFormatList,omitempty" type:"Repeated"`
	// The template ID.
	//
	// example:
	//
	// 69e1f9fe-1e97-11ed-ba64-0c42a1b73d66
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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

func (s GetLiveRecordTemplateResponseBodyRecordTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetLiveRecordTemplateResponseBodyRecordTemplate) GoString() string {
	return s.String()
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) GetLastModified() *string {
	return s.LastModified
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) GetName() *string {
	return s.Name
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) GetRecordFormatList() []*GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList {
	return s.RecordFormatList
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) GetType() *string {
	return s.Type
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) SetCreateTime(v string) *GetLiveRecordTemplateResponseBodyRecordTemplate {
	s.CreateTime = &v
	return s
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) SetLastModified(v string) *GetLiveRecordTemplateResponseBodyRecordTemplate {
	s.LastModified = &v
	return s
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) SetName(v string) *GetLiveRecordTemplateResponseBodyRecordTemplate {
	s.Name = &v
	return s
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) SetRecordFormatList(v []*GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) *GetLiveRecordTemplateResponseBodyRecordTemplate {
	s.RecordFormatList = v
	return s
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) SetTemplateId(v string) *GetLiveRecordTemplateResponseBodyRecordTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) SetType(v string) *GetLiveRecordTemplateResponseBodyRecordTemplate {
	s.Type = &v
	return s
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplate) Validate() error {
	if s.RecordFormatList != nil {
		for _, item := range s.RecordFormatList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList struct {
	// The duration of the recording cycle. Unit: seconds. If you do not specify this parameter, the default value 6 hours is used.
	//
	// example:
	//
	// 7200
	CycleDuration *int32 `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	// The output file format.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The name of the recording file that is stored in Object Storage Service (OSS).
	//
	// example:
	//
	// record/{JobId}/{Sequence}{EscapedStartTime}{EscapedEndTime}
	OssObjectPrefix *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	// The duration of a single segment. Unit: seconds.
	//
	// example:
	//
	// 30
	SliceDuration *int32 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// The name of the TS segment.
	//
	// example:
	//
	// record/{JobId}/{UnixTimestamp}_{Sequence}
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty"`
}

func (s GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) String() string {
	return dara.Prettify(s)
}

func (s GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) GoString() string {
	return s.String()
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) GetFormat() *string {
	return s.Format
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) GetOssObjectPrefix() *string {
	return s.OssObjectPrefix
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) GetSliceOssObjectPrefix() *string {
	return s.SliceOssObjectPrefix
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) SetCycleDuration(v int32) *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList {
	s.CycleDuration = &v
	return s
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) SetFormat(v string) *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList {
	s.Format = &v
	return s
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) SetOssObjectPrefix(v string) *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList {
	s.OssObjectPrefix = &v
	return s
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) SetSliceDuration(v int32) *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList {
	s.SliceDuration = &v
	return s
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) SetSliceOssObjectPrefix(v string) *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *GetLiveRecordTemplateResponseBodyRecordTemplateRecordFormatList) Validate() error {
	return dara.Validate(s)
}
