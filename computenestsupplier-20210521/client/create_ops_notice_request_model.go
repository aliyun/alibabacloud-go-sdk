// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpsNoticeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]interface{}) *CreateOpsNoticeRequest
	GetAttributes() map[string]interface{}
	SetCategory(v string) *CreateOpsNoticeRequest
	GetCategory() *string
	SetClientToken(v string) *CreateOpsNoticeRequest
	GetClientToken() *string
	SetContent(v string) *CreateOpsNoticeRequest
	GetContent() *string
	SetRegionId(v string) *CreateOpsNoticeRequest
	GetRegionId() *string
	SetServiceId(v string) *CreateOpsNoticeRequest
	GetServiceId() *string
	SetServiceVersion(v []*string) *CreateOpsNoticeRequest
	GetServiceVersion() []*string
	SetSeverity(v string) *CreateOpsNoticeRequest
	GetSeverity() *string
	SetSolutions(v string) *CreateOpsNoticeRequest
	GetSolutions() *string
	SetType(v string) *CreateOpsNoticeRequest
	GetType() *string
}

type CreateOpsNoticeRequest struct {
	// example:
	//
	// {"cveId":"CVE-2021-4034"}
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Availability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId      *string   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceVersion []*string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// You need to upgrade the service instance
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ServiceInstanceUpgrade
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateOpsNoticeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpsNoticeRequest) GoString() string {
	return s.String()
}

func (s *CreateOpsNoticeRequest) GetAttributes() map[string]interface{} {
	return s.Attributes
}

func (s *CreateOpsNoticeRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateOpsNoticeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOpsNoticeRequest) GetContent() *string {
	return s.Content
}

func (s *CreateOpsNoticeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOpsNoticeRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateOpsNoticeRequest) GetServiceVersion() []*string {
	return s.ServiceVersion
}

func (s *CreateOpsNoticeRequest) GetSeverity() *string {
	return s.Severity
}

func (s *CreateOpsNoticeRequest) GetSolutions() *string {
	return s.Solutions
}

func (s *CreateOpsNoticeRequest) GetType() *string {
	return s.Type
}

func (s *CreateOpsNoticeRequest) SetAttributes(v map[string]interface{}) *CreateOpsNoticeRequest {
	s.Attributes = v
	return s
}

func (s *CreateOpsNoticeRequest) SetCategory(v string) *CreateOpsNoticeRequest {
	s.Category = &v
	return s
}

func (s *CreateOpsNoticeRequest) SetClientToken(v string) *CreateOpsNoticeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOpsNoticeRequest) SetContent(v string) *CreateOpsNoticeRequest {
	s.Content = &v
	return s
}

func (s *CreateOpsNoticeRequest) SetRegionId(v string) *CreateOpsNoticeRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOpsNoticeRequest) SetServiceId(v string) *CreateOpsNoticeRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateOpsNoticeRequest) SetServiceVersion(v []*string) *CreateOpsNoticeRequest {
	s.ServiceVersion = v
	return s
}

func (s *CreateOpsNoticeRequest) SetSeverity(v string) *CreateOpsNoticeRequest {
	s.Severity = &v
	return s
}

func (s *CreateOpsNoticeRequest) SetSolutions(v string) *CreateOpsNoticeRequest {
	s.Solutions = &v
	return s
}

func (s *CreateOpsNoticeRequest) SetType(v string) *CreateOpsNoticeRequest {
	s.Type = &v
	return s
}

func (s *CreateOpsNoticeRequest) Validate() error {
	return dara.Validate(s)
}
