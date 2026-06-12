// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpsNoticeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributesShrink(v string) *CreateOpsNoticeShrinkRequest
	GetAttributesShrink() *string
	SetCategory(v string) *CreateOpsNoticeShrinkRequest
	GetCategory() *string
	SetClientToken(v string) *CreateOpsNoticeShrinkRequest
	GetClientToken() *string
	SetContent(v string) *CreateOpsNoticeShrinkRequest
	GetContent() *string
	SetRegionId(v string) *CreateOpsNoticeShrinkRequest
	GetRegionId() *string
	SetServiceId(v string) *CreateOpsNoticeShrinkRequest
	GetServiceId() *string
	SetServiceVersion(v []*string) *CreateOpsNoticeShrinkRequest
	GetServiceVersion() []*string
	SetSeverity(v string) *CreateOpsNoticeShrinkRequest
	GetSeverity() *string
	SetSolutions(v string) *CreateOpsNoticeShrinkRequest
	GetSolutions() *string
	SetType(v string) *CreateOpsNoticeShrinkRequest
	GetType() *string
}

type CreateOpsNoticeShrinkRequest struct {
	// The properties of the O\\&M item.
	//
	// example:
	//
	// {"cveId":"CVE-2021-4034"}
	AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The category of the notice. Valid values:
	//
	// - Availability
	//
	// - Cost
	//
	// - Performance
	//
	// - Recovery
	//
	// - Security
	//
	// This parameter is required.
	//
	// example:
	//
	// Availability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// A client token to ensure that the request is idempotent. Generate a unique token for each request. The token can contain only ASCII characters. Note: If you do not set this parameter, the system uses the RequestId as the ClientToken. The RequestId may be different for each request.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The content of the notice.
	//
	// This parameter is required.
	//
	// example:
	//
	// content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service versions to which the notice applies.
	ServiceVersion []*string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty" type:"Repeated"`
	// The severity level of the notice. Valid values:
	//
	// - Critical
	//
	// - High
	//
	// - Medium
	//
	// - Low
	//
	// This parameter is required.
	//
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solution.
	//
	// example:
	//
	// You need to upgrade the service instance
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// The type of the notice. Valid values:
	//
	// - ServiceInstanceUpgrade: Upgrade
	//
	// - VulnerabilityFix: Vulnerability
	//
	// This parameter is required.
	//
	// example:
	//
	// ServiceInstanceUpgrade
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateOpsNoticeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpsNoticeShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOpsNoticeShrinkRequest) GetAttributesShrink() *string {
	return s.AttributesShrink
}

func (s *CreateOpsNoticeShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateOpsNoticeShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOpsNoticeShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *CreateOpsNoticeShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOpsNoticeShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateOpsNoticeShrinkRequest) GetServiceVersion() []*string {
	return s.ServiceVersion
}

func (s *CreateOpsNoticeShrinkRequest) GetSeverity() *string {
	return s.Severity
}

func (s *CreateOpsNoticeShrinkRequest) GetSolutions() *string {
	return s.Solutions
}

func (s *CreateOpsNoticeShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateOpsNoticeShrinkRequest) SetAttributesShrink(v string) *CreateOpsNoticeShrinkRequest {
	s.AttributesShrink = &v
	return s
}

func (s *CreateOpsNoticeShrinkRequest) SetCategory(v string) *CreateOpsNoticeShrinkRequest {
	s.Category = &v
	return s
}

func (s *CreateOpsNoticeShrinkRequest) SetClientToken(v string) *CreateOpsNoticeShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOpsNoticeShrinkRequest) SetContent(v string) *CreateOpsNoticeShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateOpsNoticeShrinkRequest) SetRegionId(v string) *CreateOpsNoticeShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOpsNoticeShrinkRequest) SetServiceId(v string) *CreateOpsNoticeShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateOpsNoticeShrinkRequest) SetServiceVersion(v []*string) *CreateOpsNoticeShrinkRequest {
	s.ServiceVersion = v
	return s
}

func (s *CreateOpsNoticeShrinkRequest) SetSeverity(v string) *CreateOpsNoticeShrinkRequest {
	s.Severity = &v
	return s
}

func (s *CreateOpsNoticeShrinkRequest) SetSolutions(v string) *CreateOpsNoticeShrinkRequest {
	s.Solutions = &v
	return s
}

func (s *CreateOpsNoticeShrinkRequest) SetType(v string) *CreateOpsNoticeShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateOpsNoticeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
