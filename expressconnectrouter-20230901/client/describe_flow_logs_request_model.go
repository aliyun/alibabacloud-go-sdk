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
}

type DescribeFlowLogsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-a5xqrgbeidz1w8****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the flow log.
	//
	// example:
	//
	// flowlog-jqnr0veifo5d****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The flow log name. The name must be 0 to 128 characters in length.
	//
	// example:
	//
	// same-flowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The ID of the VBR associated with the ECR.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Logstore that stores the captured traffic data.
	//
	// 	- If a Logstore is already created in the selected region, enter the name of the Logstore.
	//
	// 	- If no Logstores are created in the selected region, enter a name and the system automatically creates a Logstore. The name of the Logstore. The name must meet the following requirements:
	//
	// 	- The name must be unique in a project.
	//
	// 	- It can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must start and end with a lowercase letter or a digit.
	//
	// 	- The name must be 3 to 63 characters in length.
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
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// - You do not need to specify this parameter for the first request.
	//
	// - You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// gAAAAABkWw*******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The project that stores the captured traffic data.
	//
	// 	- If a project is already created in the selected region, enter the name of the project.
	//
	// 	- If no projects are created in the selected region, enter a name and the system automatically creates a project.
	//
	// The project name must be unique in a region. You cannot change the name after the project is created. The name must meet the following requirements:
	//
	// 	- The name must be globally unique.
	//
	// 	- The name can contain only lowercase letters,
	//
	// 	- digits, and hyphens (-).
	//
	// 	- The name must start and end with a lowercase letter or a digit.
	//
	// 	- The name must be 3 to 63 characters in length.
	//
	// example:
	//
	// myFlowlog
	ProjectName     *string                       `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ResourceGroupId *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeFlowLogsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *DescribeFlowLogsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowLogsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
