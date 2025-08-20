// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *GetTemplateSummaryResponseBody
	GetDescription() *string
	SetMetadata(v map[string]interface{}) *GetTemplateSummaryResponseBody
	GetMetadata() map[string]interface{}
	SetParameters(v []map[string]interface{}) *GetTemplateSummaryResponseBody
	GetParameters() []map[string]interface{}
	SetRequestId(v string) *GetTemplateSummaryResponseBody
	GetRequestId() *string
	SetResourceIdentifierSummaries(v []*GetTemplateSummaryResponseBodyResourceIdentifierSummaries) *GetTemplateSummaryResponseBody
	GetResourceIdentifierSummaries() []*GetTemplateSummaryResponseBodyResourceIdentifierSummaries
	SetResourceTypes(v []*string) *GetTemplateSummaryResponseBody
	GetResourceTypes() []*string
	SetVersion(v string) *GetTemplateSummaryResponseBody
	GetVersion() *string
}

type GetTemplateSummaryResponseBody struct {
	// The description of the stack template.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The metadata that is defined in the template.
	//
	// example:
	//
	// {"key": "value"}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The declarations of the parameters in the template.
	Parameters []map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource identifier summaries.\\
	//
	// A summary describes the resource that you want to import and the properties that are used to identify the resource during the import. For example, VpcId is an identifier property of ALIYUN::ECS::VPC.
	ResourceIdentifierSummaries []*GetTemplateSummaryResponseBodyResourceIdentifierSummaries `json:"ResourceIdentifierSummaries,omitempty" xml:"ResourceIdentifierSummaries,omitempty" type:"Repeated"`
	// All resource types that are used in the template.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The version of the template.
	//
	// example:
	//
	// 2015-09-01
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetTemplateSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetTemplateSummaryResponseBody) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *GetTemplateSummaryResponseBody) GetParameters() []map[string]interface{} {
	return s.Parameters
}

func (s *GetTemplateSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateSummaryResponseBody) GetResourceIdentifierSummaries() []*GetTemplateSummaryResponseBodyResourceIdentifierSummaries {
	return s.ResourceIdentifierSummaries
}

func (s *GetTemplateSummaryResponseBody) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *GetTemplateSummaryResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetTemplateSummaryResponseBody) SetDescription(v string) *GetTemplateSummaryResponseBody {
	s.Description = &v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetMetadata(v map[string]interface{}) *GetTemplateSummaryResponseBody {
	s.Metadata = v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetParameters(v []map[string]interface{}) *GetTemplateSummaryResponseBody {
	s.Parameters = v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetRequestId(v string) *GetTemplateSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetResourceIdentifierSummaries(v []*GetTemplateSummaryResponseBodyResourceIdentifierSummaries) *GetTemplateSummaryResponseBody {
	s.ResourceIdentifierSummaries = v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetResourceTypes(v []*string) *GetTemplateSummaryResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetVersion(v string) *GetTemplateSummaryResponseBody {
	s.Version = &v
	return s
}

func (s *GetTemplateSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTemplateSummaryResponseBodyResourceIdentifierSummaries struct {
	// The logical IDs of all resources of the type that is specified by ResouceType in the template.
	LogicalResourceIds []*string `json:"LogicalResourceIds,omitempty" xml:"LogicalResourceIds,omitempty" type:"Repeated"`
	// The resource properties. You can use a resource property to identify the resource that you want to manage. For example, VpcId is an identifier property of ALIYUN::ECS::VPC.
	ResourceIdentifiers []*string `json:"ResourceIdentifiers,omitempty" xml:"ResourceIdentifiers,omitempty" type:"Repeated"`
	// The resource type.
	//
	// > The resource import feature is supported for the resource type.
	//
	// example:
	//
	// ALIYUN::ECS::VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTemplateSummaryResponseBodyResourceIdentifierSummaries) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateSummaryResponseBodyResourceIdentifierSummaries) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) GetLogicalResourceIds() []*string {
	return s.LogicalResourceIds
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) GetResourceIdentifiers() []*string {
	return s.ResourceIdentifiers
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) SetLogicalResourceIds(v []*string) *GetTemplateSummaryResponseBodyResourceIdentifierSummaries {
	s.LogicalResourceIds = v
	return s
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) SetResourceIdentifiers(v []*string) *GetTemplateSummaryResponseBodyResourceIdentifierSummaries {
	s.ResourceIdentifiers = v
	return s
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) SetResourceType(v string) *GetTemplateSummaryResponseBodyResourceIdentifierSummaries {
	s.ResourceType = &v
	return s
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) Validate() error {
	return dara.Validate(s)
}
