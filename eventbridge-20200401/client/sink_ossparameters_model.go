// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSinkOSSParameters interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *SinkOSSParameters
	GetBucketName() *string
	SetCompressionType(v string) *SinkOSSParameters
	GetCompressionType() *string
	SetContentTransform(v *SinkOSSParametersContentTransform) *SinkOSSParameters
	GetContentTransform() *SinkOSSParametersContentTransform
	SetEndpoint(v string) *SinkOSSParameters
	GetEndpoint() *string
	SetPathFormat(v string) *SinkOSSParameters
	GetPathFormat() *string
	SetRegionId(v string) *SinkOSSParameters
	GetRegionId() *string
	SetRoleArn(v string) *SinkOSSParameters
	GetRoleArn() *string
	SetRotateIntervalMs(v string) *SinkOSSParameters
	GetRotateIntervalMs() *string
	SetRotateSizeBytes(v string) *SinkOSSParameters
	GetRotateSizeBytes() *string
	SetSSLEnabled(v bool) *SinkOSSParameters
	GetSSLEnabled() *bool
	SetTaskConcurrency(v string) *SinkOSSParameters
	GetTaskConcurrency() *string
	SetTimeZone(v string) *SinkOSSParameters
	GetTimeZone() *string
}

type SinkOSSParameters struct {
	BucketName       *string                            `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	CompressionType  *string                            `json:"CompressionType,omitempty" xml:"CompressionType,omitempty"`
	ContentTransform *SinkOSSParametersContentTransform `json:"ContentTransform,omitempty" xml:"ContentTransform,omitempty" type:"Struct"`
	Endpoint         *string                            `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	PathFormat       *string                            `json:"PathFormat,omitempty" xml:"PathFormat,omitempty"`
	RegionId         *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleArn          *string                            `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	RotateIntervalMs *string                            `json:"RotateIntervalMs,omitempty" xml:"RotateIntervalMs,omitempty"`
	RotateSizeBytes  *string                            `json:"RotateSizeBytes,omitempty" xml:"RotateSizeBytes,omitempty"`
	SSLEnabled       *bool                              `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	TaskConcurrency  *string                            `json:"TaskConcurrency,omitempty" xml:"TaskConcurrency,omitempty"`
	TimeZone         *string                            `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s SinkOSSParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkOSSParameters) GoString() string {
	return s.String()
}

func (s *SinkOSSParameters) GetBucketName() *string {
	return s.BucketName
}

func (s *SinkOSSParameters) GetCompressionType() *string {
	return s.CompressionType
}

func (s *SinkOSSParameters) GetContentTransform() *SinkOSSParametersContentTransform {
	return s.ContentTransform
}

func (s *SinkOSSParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *SinkOSSParameters) GetPathFormat() *string {
	return s.PathFormat
}

func (s *SinkOSSParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *SinkOSSParameters) GetRoleArn() *string {
	return s.RoleArn
}

func (s *SinkOSSParameters) GetRotateIntervalMs() *string {
	return s.RotateIntervalMs
}

func (s *SinkOSSParameters) GetRotateSizeBytes() *string {
	return s.RotateSizeBytes
}

func (s *SinkOSSParameters) GetSSLEnabled() *bool {
	return s.SSLEnabled
}

func (s *SinkOSSParameters) GetTaskConcurrency() *string {
	return s.TaskConcurrency
}

func (s *SinkOSSParameters) GetTimeZone() *string {
	return s.TimeZone
}

func (s *SinkOSSParameters) SetBucketName(v string) *SinkOSSParameters {
	s.BucketName = &v
	return s
}

func (s *SinkOSSParameters) SetCompressionType(v string) *SinkOSSParameters {
	s.CompressionType = &v
	return s
}

func (s *SinkOSSParameters) SetContentTransform(v *SinkOSSParametersContentTransform) *SinkOSSParameters {
	s.ContentTransform = v
	return s
}

func (s *SinkOSSParameters) SetEndpoint(v string) *SinkOSSParameters {
	s.Endpoint = &v
	return s
}

func (s *SinkOSSParameters) SetPathFormat(v string) *SinkOSSParameters {
	s.PathFormat = &v
	return s
}

func (s *SinkOSSParameters) SetRegionId(v string) *SinkOSSParameters {
	s.RegionId = &v
	return s
}

func (s *SinkOSSParameters) SetRoleArn(v string) *SinkOSSParameters {
	s.RoleArn = &v
	return s
}

func (s *SinkOSSParameters) SetRotateIntervalMs(v string) *SinkOSSParameters {
	s.RotateIntervalMs = &v
	return s
}

func (s *SinkOSSParameters) SetRotateSizeBytes(v string) *SinkOSSParameters {
	s.RotateSizeBytes = &v
	return s
}

func (s *SinkOSSParameters) SetSSLEnabled(v bool) *SinkOSSParameters {
	s.SSLEnabled = &v
	return s
}

func (s *SinkOSSParameters) SetTaskConcurrency(v string) *SinkOSSParameters {
	s.TaskConcurrency = &v
	return s
}

func (s *SinkOSSParameters) SetTimeZone(v string) *SinkOSSParameters {
	s.TimeZone = &v
	return s
}

func (s *SinkOSSParameters) Validate() error {
	if s.ContentTransform != nil {
		if err := s.ContentTransform.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SinkOSSParametersContentTransform struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkOSSParametersContentTransform) String() string {
	return dara.Prettify(s)
}

func (s SinkOSSParametersContentTransform) GoString() string {
	return s.String()
}

func (s *SinkOSSParametersContentTransform) GetForm() *string {
	return s.Form
}

func (s *SinkOSSParametersContentTransform) GetTemplate() *string {
	return s.Template
}

func (s *SinkOSSParametersContentTransform) GetValue() *string {
	return s.Value
}

func (s *SinkOSSParametersContentTransform) SetForm(v string) *SinkOSSParametersContentTransform {
	s.Form = &v
	return s
}

func (s *SinkOSSParametersContentTransform) SetTemplate(v string) *SinkOSSParametersContentTransform {
	s.Template = &v
	return s
}

func (s *SinkOSSParametersContentTransform) SetValue(v string) *SinkOSSParametersContentTransform {
	s.Value = &v
	return s
}

func (s *SinkOSSParametersContentTransform) Validate() error {
	return dara.Validate(s)
}
