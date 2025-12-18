// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCompliancePackReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GenerateCompliancePackReportRequest
	GetClientToken() *string
	SetCompliancePackId(v string) *GenerateCompliancePackReportRequest
	GetCompliancePackId() *string
}

type GenerateCompliancePackReportRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AAAAAdDWBF2****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-a8a8626622af0082****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GenerateCompliancePackReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCompliancePackReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateCompliancePackReportRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GenerateCompliancePackReportRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GenerateCompliancePackReportRequest) SetClientToken(v string) *GenerateCompliancePackReportRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateCompliancePackReportRequest) SetCompliancePackId(v string) *GenerateCompliancePackReportRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GenerateCompliancePackReportRequest) Validate() error {
	return dara.Validate(s)
}
