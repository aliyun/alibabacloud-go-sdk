// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerModuleOutputInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *CreateCustomerModuleOutputInfoRequest
	GetCustomerModuleId() *int32
	SetFinalScoreFormat(v string) *CreateCustomerModuleOutputInfoRequest
	GetFinalScoreFormat() *string
	SetProcessExpression(v string) *CreateCustomerModuleOutputInfoRequest
	GetProcessExpression() *string
	SetTestFileUrl(v string) *CreateCustomerModuleOutputInfoRequest
	GetTestFileUrl() *string
}

type CreateCustomerModuleOutputInfoRequest struct {
	// Customer model ID
	//
	// example:
	//
	// 1
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
	// Number of decimal places to retain.
	//
	// example:
	//
	// 2
	FinalScoreFormat *string `json:"FinalScoreFormat,omitempty" xml:"FinalScoreFormat,omitempty"`
	// Score processing logic.
	//
	// example:
	//
	// score
	ProcessExpression *string `json:"ProcessExpression,omitempty" xml:"ProcessExpression,omitempty"`
	// Test file URL.
	//
	// example:
	//
	// customer/xxxxxxxxx/xxxxxxxx.pmml
	TestFileUrl *string `json:"TestFileUrl,omitempty" xml:"TestFileUrl,omitempty"`
}

func (s CreateCustomerModuleOutputInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerModuleOutputInfoRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerModuleOutputInfoRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *CreateCustomerModuleOutputInfoRequest) GetFinalScoreFormat() *string {
	return s.FinalScoreFormat
}

func (s *CreateCustomerModuleOutputInfoRequest) GetProcessExpression() *string {
	return s.ProcessExpression
}

func (s *CreateCustomerModuleOutputInfoRequest) GetTestFileUrl() *string {
	return s.TestFileUrl
}

func (s *CreateCustomerModuleOutputInfoRequest) SetCustomerModuleId(v int32) *CreateCustomerModuleOutputInfoRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *CreateCustomerModuleOutputInfoRequest) SetFinalScoreFormat(v string) *CreateCustomerModuleOutputInfoRequest {
	s.FinalScoreFormat = &v
	return s
}

func (s *CreateCustomerModuleOutputInfoRequest) SetProcessExpression(v string) *CreateCustomerModuleOutputInfoRequest {
	s.ProcessExpression = &v
	return s
}

func (s *CreateCustomerModuleOutputInfoRequest) SetTestFileUrl(v string) *CreateCustomerModuleOutputInfoRequest {
	s.TestFileUrl = &v
	return s
}

func (s *CreateCustomerModuleOutputInfoRequest) Validate() error {
	return dara.Validate(s)
}
