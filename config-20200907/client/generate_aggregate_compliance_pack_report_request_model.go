// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAggregateCompliancePackReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GenerateAggregateCompliancePackReportRequest
	GetAggregatorId() *string
	SetClientToken(v string) *GenerateAggregateCompliancePackReportRequest
	GetClientToken() *string
	SetCompliancePackId(v string) *GenerateAggregateCompliancePackReportRequest
	GetCompliancePackId() *string
	SetMultiFiles(v bool) *GenerateAggregateCompliancePackReportRequest
	GetMultiFiles() *bool
}

type GenerateAggregateCompliancePackReportRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	MultiFiles       *bool   `json:"MultiFiles,omitempty" xml:"MultiFiles,omitempty"`
}

func (s GenerateAggregateCompliancePackReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAggregateCompliancePackReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateAggregateCompliancePackReportRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GenerateAggregateCompliancePackReportRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GenerateAggregateCompliancePackReportRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GenerateAggregateCompliancePackReportRequest) GetMultiFiles() *bool {
	return s.MultiFiles
}

func (s *GenerateAggregateCompliancePackReportRequest) SetAggregatorId(v string) *GenerateAggregateCompliancePackReportRequest {
	s.AggregatorId = &v
	return s
}

func (s *GenerateAggregateCompliancePackReportRequest) SetClientToken(v string) *GenerateAggregateCompliancePackReportRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateAggregateCompliancePackReportRequest) SetCompliancePackId(v string) *GenerateAggregateCompliancePackReportRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GenerateAggregateCompliancePackReportRequest) SetMultiFiles(v bool) *GenerateAggregateCompliancePackReportRequest {
	s.MultiFiles = &v
	return s
}

func (s *GenerateAggregateCompliancePackReportRequest) Validate() error {
	return dara.Validate(s)
}
