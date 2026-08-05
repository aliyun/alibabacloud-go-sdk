// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperienceDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *CreateExperienceDataRequest
	GetContentType() *string
	SetDataSize(v int64) *CreateExperienceDataRequest
	GetDataSize() *int64
	SetDataType(v string) *CreateExperienceDataRequest
	GetDataType() *string
	SetDataValue(v string) *CreateExperienceDataRequest
	GetDataValue() *string
	SetName(v string) *CreateExperienceDataRequest
	GetName() *string
	SetServiceType(v string) *CreateExperienceDataRequest
	GetServiceType() *string
	SetDryRun(v bool) *CreateExperienceDataRequest
	GetDryRun() *bool
}

type CreateExperienceDataRequest struct {
	// The data content type. Valid values:
	//
	// - pdf
	//
	// - text
	//
	// - html
	//
	// - doc.
	//
	// example:
	//
	// text
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The data size.
	//
	// example:
	//
	// 100
	DataSize *int64 `json:"dataSize,omitempty" xml:"dataSize,omitempty"`
	// The data type. Valid values:
	//
	// - file: file
	//
	// - url: URL.
	//
	// example:
	//
	// file
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// The data content.
	//
	// - If dataType is set to file, this field specifies the OSS address of the file.
	//
	// - If dataType is set to url, this field specifies the HTTP URL of the data.
	//
	// example:
	//
	// https://xxx
	DataValue *string `json:"dataValue,omitempty" xml:"dataValue,omitempty"`
	// The data name. This parameter is required when dataType is set to file.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service type.
	//
	// example:
	//
	// document-analyze
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// Specifies whether to perform a dry run request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateExperienceDataRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExperienceDataRequest) GoString() string {
	return s.String()
}

func (s *CreateExperienceDataRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreateExperienceDataRequest) GetDataSize() *int64 {
	return s.DataSize
}

func (s *CreateExperienceDataRequest) GetDataType() *string {
	return s.DataType
}

func (s *CreateExperienceDataRequest) GetDataValue() *string {
	return s.DataValue
}

func (s *CreateExperienceDataRequest) GetName() *string {
	return s.Name
}

func (s *CreateExperienceDataRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *CreateExperienceDataRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateExperienceDataRequest) SetContentType(v string) *CreateExperienceDataRequest {
	s.ContentType = &v
	return s
}

func (s *CreateExperienceDataRequest) SetDataSize(v int64) *CreateExperienceDataRequest {
	s.DataSize = &v
	return s
}

func (s *CreateExperienceDataRequest) SetDataType(v string) *CreateExperienceDataRequest {
	s.DataType = &v
	return s
}

func (s *CreateExperienceDataRequest) SetDataValue(v string) *CreateExperienceDataRequest {
	s.DataValue = &v
	return s
}

func (s *CreateExperienceDataRequest) SetName(v string) *CreateExperienceDataRequest {
	s.Name = &v
	return s
}

func (s *CreateExperienceDataRequest) SetServiceType(v string) *CreateExperienceDataRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateExperienceDataRequest) SetDryRun(v bool) *CreateExperienceDataRequest {
	s.DryRun = &v
	return s
}

func (s *CreateExperienceDataRequest) Validate() error {
	return dara.Validate(s)
}
