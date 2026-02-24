// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemediationTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListRemediationTemplatesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRemediationTemplatesResponseBody
	GetPageSize() *int64
	SetRemediationTemplates(v []*ListRemediationTemplatesResponseBodyRemediationTemplates) *ListRemediationTemplatesResponseBody
	GetRemediationTemplates() []*ListRemediationTemplatesResponseBodyRemediationTemplates
	SetRequestId(v string) *ListRemediationTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRemediationTemplatesResponseBody
	GetTotalCount() *string
}

type ListRemediationTemplatesResponseBody struct {
	PageNumber           *int64                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int64                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RemediationTemplates []*ListRemediationTemplatesResponseBodyRemediationTemplates `json:"RemediationTemplates,omitempty" xml:"RemediationTemplates,omitempty" type:"Repeated"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *string                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRemediationTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRemediationTemplatesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRemediationTemplatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRemediationTemplatesResponseBody) GetRemediationTemplates() []*ListRemediationTemplatesResponseBodyRemediationTemplates {
	return s.RemediationTemplates
}

func (s *ListRemediationTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRemediationTemplatesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRemediationTemplatesResponseBody) SetPageNumber(v int64) *ListRemediationTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRemediationTemplatesResponseBody) SetPageSize(v int64) *ListRemediationTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRemediationTemplatesResponseBody) SetRemediationTemplates(v []*ListRemediationTemplatesResponseBodyRemediationTemplates) *ListRemediationTemplatesResponseBody {
	s.RemediationTemplates = v
	return s
}

func (s *ListRemediationTemplatesResponseBody) SetRequestId(v string) *ListRemediationTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRemediationTemplatesResponseBody) SetTotalCount(v string) *ListRemediationTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRemediationTemplatesResponseBody) Validate() error {
	if s.RemediationTemplates != nil {
		for _, item := range s.RemediationTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRemediationTemplatesResponseBodyRemediationTemplates struct {
	RemediationType     *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	TemplateDefinition  *string `json:"TemplateDefinition,omitempty" xml:"TemplateDefinition,omitempty"`
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	TemplateIdentifier  *string `json:"TemplateIdentifier,omitempty" xml:"TemplateIdentifier,omitempty"`
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListRemediationTemplatesResponseBodyRemediationTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationTemplatesResponseBodyRemediationTemplates) GoString() string {
	return s.String()
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) GetRemediationType() *string {
	return s.RemediationType
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) GetTemplateDefinition() *string {
	return s.TemplateDefinition
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) GetTemplateDescription() *string {
	return s.TemplateDescription
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) GetTemplateIdentifier() *string {
	return s.TemplateIdentifier
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetRemediationType(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.RemediationType = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetTemplateDefinition(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.TemplateDefinition = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetTemplateDescription(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.TemplateDescription = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetTemplateIdentifier(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.TemplateIdentifier = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetTemplateName(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) Validate() error {
	return dara.Validate(s)
}
