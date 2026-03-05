// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleResultsValue interface {
	dara.Model
	String() string
	GoString() string
	SetPassed(v bool) *ModuleResultsValue
	GetPassed() *bool
	SetResourceCode(v string) *ModuleResultsValue
	GetResourceCode() *string
	SetErrorCode(v string) *ModuleResultsValue
	GetErrorCode() *string
	SetErrorMessage(v string) *ModuleResultsValue
	GetErrorMessage() *string
}

type ModuleResultsValue struct {
	// example:
	//
	// true
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// InspirationTokens
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// example:
	//
	// Resource.Control.No.Usage
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s ModuleResultsValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleResultsValue) GoString() string {
	return s.String()
}

func (s *ModuleResultsValue) GetPassed() *bool {
	return s.Passed
}

func (s *ModuleResultsValue) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *ModuleResultsValue) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModuleResultsValue) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModuleResultsValue) SetPassed(v bool) *ModuleResultsValue {
	s.Passed = &v
	return s
}

func (s *ModuleResultsValue) SetResourceCode(v string) *ModuleResultsValue {
	s.ResourceCode = &v
	return s
}

func (s *ModuleResultsValue) SetErrorCode(v string) *ModuleResultsValue {
	s.ErrorCode = &v
	return s
}

func (s *ModuleResultsValue) SetErrorMessage(v string) *ModuleResultsValue {
	s.ErrorMessage = &v
	return s
}

func (s *ModuleResultsValue) Validate() error {
	return dara.Validate(s)
}
