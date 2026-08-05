// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperienceDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *ListExperienceDataRequest
	GetDataType() *string
	SetDryRun(v bool) *ListExperienceDataRequest
	GetDryRun() *bool
	SetServiceType(v string) *ListExperienceDataRequest
	GetServiceType() *string
}

type ListExperienceDataRequest struct {
	// The data type.
	//
	// example:
	//
	// file
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// - true
	//
	// - false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The service type.
	//
	// example:
	//
	// document-analyze
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListExperienceDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExperienceDataRequest) GoString() string {
	return s.String()
}

func (s *ListExperienceDataRequest) GetDataType() *string {
	return s.DataType
}

func (s *ListExperienceDataRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListExperienceDataRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListExperienceDataRequest) SetDataType(v string) *ListExperienceDataRequest {
	s.DataType = &v
	return s
}

func (s *ListExperienceDataRequest) SetDryRun(v bool) *ListExperienceDataRequest {
	s.DryRun = &v
	return s
}

func (s *ListExperienceDataRequest) SetServiceType(v string) *ListExperienceDataRequest {
	s.ServiceType = &v
	return s
}

func (s *ListExperienceDataRequest) Validate() error {
	return dara.Validate(s)
}
