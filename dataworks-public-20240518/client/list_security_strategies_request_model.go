// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityStrategiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlModule(v string) *ListSecurityStrategiesRequest
	GetControlModule() *string
	SetControlSubModule(v string) *ListSecurityStrategiesRequest
	GetControlSubModule() *string
	SetPageNumber(v int32) *ListSecurityStrategiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSecurityStrategiesRequest
	GetPageSize() *int32
}

type ListSecurityStrategiesRequest struct {
	// example:
	//
	// DataQuery
	ControlModule *string `json:"ControlModule,omitempty" xml:"ControlModule,omitempty"`
	// example:
	//
	// MyCatalog
	ControlSubModule *string `json:"ControlSubModule,omitempty" xml:"ControlSubModule,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSecurityStrategiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityStrategiesRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityStrategiesRequest) GetControlModule() *string {
	return s.ControlModule
}

func (s *ListSecurityStrategiesRequest) GetControlSubModule() *string {
	return s.ControlSubModule
}

func (s *ListSecurityStrategiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSecurityStrategiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecurityStrategiesRequest) SetControlModule(v string) *ListSecurityStrategiesRequest {
	s.ControlModule = &v
	return s
}

func (s *ListSecurityStrategiesRequest) SetControlSubModule(v string) *ListSecurityStrategiesRequest {
	s.ControlSubModule = &v
	return s
}

func (s *ListSecurityStrategiesRequest) SetPageNumber(v int32) *ListSecurityStrategiesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecurityStrategiesRequest) SetPageSize(v int32) *ListSecurityStrategiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSecurityStrategiesRequest) Validate() error {
	return dara.Validate(s)
}
