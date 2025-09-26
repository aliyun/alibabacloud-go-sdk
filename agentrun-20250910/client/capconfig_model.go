// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCAPConfig interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *CAPConfig
	GetFunctionName() *string
	SetName(v string) *CAPConfig
	GetName() *string
	SetTemplateId(v int32) *CAPConfig
	GetTemplateId() *int32
}

type CAPConfig struct {
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	TemplateId   *int32  `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CAPConfig) String() string {
	return dara.Prettify(s)
}

func (s CAPConfig) GoString() string {
	return s.String()
}

func (s *CAPConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *CAPConfig) GetName() *string {
	return s.Name
}

func (s *CAPConfig) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *CAPConfig) SetFunctionName(v string) *CAPConfig {
	s.FunctionName = &v
	return s
}

func (s *CAPConfig) SetName(v string) *CAPConfig {
	s.Name = &v
	return s
}

func (s *CAPConfig) SetTemplateId(v int32) *CAPConfig {
	s.TemplateId = &v
	return s
}

func (s *CAPConfig) Validate() error {
	return dara.Validate(s)
}
