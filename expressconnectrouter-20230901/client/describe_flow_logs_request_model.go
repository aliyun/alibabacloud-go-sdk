// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeFlowLogsRequest
	GetClientToken() *string
	SetDryRun(v bool) *DescribeFlowLogsRequest
	GetDryRun() *bool
	SetEcrId(v string) *DescribeFlowLogsRequest
	GetEcrId() *string
	SetFlowLogId(v string) *DescribeFlowLogsRequest
	GetFlowLogId() *string
	SetFlowLogName(v string) *DescribeFlowLogsRequest
	GetFlowLogName() *string
	SetInstanceId(v string) *DescribeFlowLogsRequest
	GetInstanceId() *string
	SetLogStoreName(v string) *DescribeFlowLogsRequest
	GetLogStoreName() *string
	SetMaxResults(v int32) *DescribeFlowLogsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeFlowLogsRequest
	GetNextToken() *string
	SetProjectName(v string) *DescribeFlowLogsRequest
	GetProjectName() *string
	SetResourceGroupId(v string) *DescribeFlowLogsRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeFlowLogsRequestTag) *DescribeFlowLogsRequest
	GetTag() []*DescribeFlowLogsRequestTag
	SetVersion(v string) *DescribeFlowLogsRequest
	GetVersion() *string
}

type DescribeFlowLogsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a parameter value on your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only perform a dry run for this request. Valid values:
	//
	// - **true**: send a check request and do not modify service resources.
	//
	// - **false*	- (default): send a normal request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-a5xqrgbeidz1w8****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The instance ID of the flow log instance.
	//
	// example:
	//
	// ecr-flog-jqnr0veifo5d****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The name of the flow log. The name must be 0 to 128 characters in length.
	//
	// example:
	//
	// same-flowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The instance ID of the VBR instance loaded on the Express Connect Router (ECR).
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Simple Log Service Logstore used to capture traffic.
	//
	// - If you have created a Logstore in the current region, you can enter the name of the existing Logstore.
	//
	// - If you have not created a Logstore in the current region, you can specify a custom name for the Logstore, and the system automatically creates the Logstore. The naming rules are as follows:
	//
	// - The Logstore name must be unique within the same project.
	//
	// - It can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// - It must start and end with a lowercase letter or a digit.
	//
	// - It must be 3 to 63 characters in length.
	//
	// example:
	//
	// myFlowlog
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 2147483647. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If NextToken is empty, no next query is available.
	//
	// - If NextToken has a return value, the value indicates the token for the next query.
	//
	// example:
	//
	// gAAAAABkWw*******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The Simple Log Service project used to capture traffic.
	//
	// - If you have created a project in the current region, you can enter the name of the existing project.
	//
	// - If you have not created a project in the current region, you can specify a custom name for the project, and the system automatically creates the project.
	//
	//
	// The project name must be globally unique within an Alibaba Cloud region and cannot be modified after creation. The naming rules are as follows:
	//
	//
	//
	// - The project name must be globally unique.
	//
	// - It can contain only lowercase
	//
	// - letters, digits, and hyphens (-).
	//
	// - It must start and end with a lowercase letter or a digit.
	//
	// - It must be 3 to 63 characters in length.
	//
	// example:
	//
	// myFlowlog
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the resource group to which the flow log instance belongs.
	//
	// example:
	//
	// rg-aek2aq7f4va****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	//
	// You can add up to 20 tags.
	Tag     []*DescribeFlowLogsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Version *string                       `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeFlowLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeFlowLogsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeFlowLogsRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeFlowLogsRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DescribeFlowLogsRequest) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *DescribeFlowLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFlowLogsRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeFlowLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeFlowLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFlowLogsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeFlowLogsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeFlowLogsRequest) GetTag() []*DescribeFlowLogsRequestTag {
	return s.Tag
}

func (s *DescribeFlowLogsRequest) GetVersion() *string {
	return s.Version
}

func (s *DescribeFlowLogsRequest) SetClientToken(v string) *DescribeFlowLogsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetDryRun(v bool) *DescribeFlowLogsRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetEcrId(v string) *DescribeFlowLogsRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetFlowLogId(v string) *DescribeFlowLogsRequest {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetFlowLogName(v string) *DescribeFlowLogsRequest {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetInstanceId(v string) *DescribeFlowLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetLogStoreName(v string) *DescribeFlowLogsRequest {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetMaxResults(v int32) *DescribeFlowLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetNextToken(v string) *DescribeFlowLogsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetProjectName(v string) *DescribeFlowLogsRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetResourceGroupId(v string) *DescribeFlowLogsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetTag(v []*DescribeFlowLogsRequestTag) *DescribeFlowLogsRequest {
	s.Tag = v
	return s
}

func (s *DescribeFlowLogsRequest) SetVersion(v string) *DescribeFlowLogsRequest {
	s.Version = &v
	return s
}

func (s *DescribeFlowLogsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFlowLogsRequestTag struct {
	// The tag key of the resource. If you specify this parameter, it cannot be an empty string. The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20. If you specify this parameter, it can be an empty string. The tag value can be up to 128 characters in length. It cannot start with `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFlowLogsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeFlowLogsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeFlowLogsRequestTag) SetKey(v string) *DescribeFlowLogsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeFlowLogsRequestTag) SetValue(v string) *DescribeFlowLogsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeFlowLogsRequestTag) Validate() error {
	return dara.Validate(s)
}
