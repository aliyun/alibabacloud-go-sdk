// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowlogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateFlowlogRequest
	GetCenId() *string
	SetClientToken(v string) *CreateFlowlogRequest
	GetClientToken() *string
	SetDescription(v string) *CreateFlowlogRequest
	GetDescription() *string
	SetFlowLogName(v string) *CreateFlowlogRequest
	GetFlowLogName() *string
	SetInterval(v int64) *CreateFlowlogRequest
	GetInterval() *int64
	SetLogFormatString(v string) *CreateFlowlogRequest
	GetLogFormatString() *string
	SetLogStoreName(v string) *CreateFlowlogRequest
	GetLogStoreName() *string
	SetOwnerAccount(v string) *CreateFlowlogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateFlowlogRequest
	GetOwnerId() *int64
	SetProjectName(v string) *CreateFlowlogRequest
	GetProjectName() *string
	SetRegionId(v string) *CreateFlowlogRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateFlowlogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateFlowlogRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateFlowlogRequestTag) *CreateFlowlogRequest
	GetTag() []*CreateFlowlogRequestTag
	SetTransitRouterAttachmentId(v string) *CreateFlowlogRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterId(v string) *CreateFlowlogRequest
	GetTransitRouterId() *string
}

type CreateFlowlogRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, ClientToken is set to the value of RequestId. The value of RequestId for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the flow log.
	//
	// The description is optional. If you enter a description, it must be 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The flow log name.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The time window for collecting log data. Unit: seconds. Valid values: **60*	- and **600**. Default value: **600**.
	//
	// example:
	//
	// 600
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The strings that define the fields in the flow log.
	//
	// Format: `${Field 1}${Field 2}${Field 3}...{Field n}`
	//
	// 	- If you do not configure this parameter, all fields are included in the flow log.
	//
	// 	- If you configure this parameter, start the string with `${srcaddr}${dstaddr}${bytes}` because `${srcaddr}${dstaddr}${bytes}` are required variables. For more information about the fields supported by flow logs, see [Configure a flow log](https://help.aliyun.com/document_detail/339822.html).
	//
	// example:
	//
	// ${srcaddr}${dstaddr}${bytes}
	LogFormatString *string `json:"LogFormatString,omitempty" xml:"LogFormatString,omitempty"`
	// The Logstore that stores the captured traffic data.
	//
	// 	- If a Logstore is already created in the selected region, enter the name of the Logstore.
	//
	// 	- If no Logstores are created in the selected region, enter a name and the system automatically creates a Logstore. The name of the Logstore. The name must meet the following requirements:
	//
	//     	- The name must be unique in a project.
	//
	//     	- The name can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	//     	- The name must start and end with a lowercase letter or a digit.
	//
	//     	- The name must be 3 to 63 characters in length,
	//
	// example:
	//
	// FlowLogStore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The project that stores the captured traffic data.
	//
	// 	- If a project is already created in the selected region, enter the name of the project.
	//
	// 	- If no projects are created in the selected region, enter a name and the system automatically creates a project.
	//
	//     The project name must be unique in a region. You cannot change the name after the project is created. The name must meet the following requirements:
	//
	//     	- The name must be globally unique.
	//
	//     	- The name can contain only lowercase letters, digits, and hyphens (-).
	//
	//     	- The name must start and end with a lowercase letter or a digit.
	//
	//     	- The name must be 3 to 63 characters in length,
	//
	// example:
	//
	// FlowLogProject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the region where the flow log is deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	//
	// You can specify at most 20 tags.
	Tag []*CreateFlowlogRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC connection, VPN connection, VBR connection, ECR connection, or inter-region connection.
	//
	// If you create the flow log for a transfer router, skip this parameter.
	//
	// example:
	//
	// tr-attach-r6g0m3epjehw57****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-bp1rmwxnk221e3fas****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateFlowlogRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowlogRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowlogRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateFlowlogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFlowlogRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowlogRequest) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *CreateFlowlogRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *CreateFlowlogRequest) GetLogFormatString() *string {
	return s.LogFormatString
}

func (s *CreateFlowlogRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *CreateFlowlogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateFlowlogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFlowlogRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFlowlogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFlowlogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFlowlogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateFlowlogRequest) GetTag() []*CreateFlowlogRequestTag {
	return s.Tag
}

func (s *CreateFlowlogRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *CreateFlowlogRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateFlowlogRequest) SetCenId(v string) *CreateFlowlogRequest {
	s.CenId = &v
	return s
}

func (s *CreateFlowlogRequest) SetClientToken(v string) *CreateFlowlogRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowlogRequest) SetDescription(v string) *CreateFlowlogRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowlogRequest) SetFlowLogName(v string) *CreateFlowlogRequest {
	s.FlowLogName = &v
	return s
}

func (s *CreateFlowlogRequest) SetInterval(v int64) *CreateFlowlogRequest {
	s.Interval = &v
	return s
}

func (s *CreateFlowlogRequest) SetLogFormatString(v string) *CreateFlowlogRequest {
	s.LogFormatString = &v
	return s
}

func (s *CreateFlowlogRequest) SetLogStoreName(v string) *CreateFlowlogRequest {
	s.LogStoreName = &v
	return s
}

func (s *CreateFlowlogRequest) SetOwnerAccount(v string) *CreateFlowlogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateFlowlogRequest) SetOwnerId(v int64) *CreateFlowlogRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFlowlogRequest) SetProjectName(v string) *CreateFlowlogRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFlowlogRequest) SetRegionId(v string) *CreateFlowlogRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowlogRequest) SetResourceOwnerAccount(v string) *CreateFlowlogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFlowlogRequest) SetResourceOwnerId(v int64) *CreateFlowlogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFlowlogRequest) SetTag(v []*CreateFlowlogRequestTag) *CreateFlowlogRequest {
	s.Tag = v
	return s
}

func (s *CreateFlowlogRequest) SetTransitRouterAttachmentId(v string) *CreateFlowlogRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *CreateFlowlogRequest) SetTransitRouterId(v string) *CreateFlowlogRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateFlowlogRequest) Validate() error {
	return dara.Validate(s)
}

type CreateFlowlogRequestTag struct {
	// The tag keys.
	//
	// The tag keys cannot be an empty string. The tag keys can be up to 64 characters in length. The tag keys cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	//
	// The tag values can be an empty string or up to 128 characters in length. The tag values cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each key-value must be unique. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFlowlogRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowlogRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFlowlogRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateFlowlogRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateFlowlogRequestTag) SetKey(v string) *CreateFlowlogRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFlowlogRequestTag) SetValue(v string) *CreateFlowlogRequestTag {
	s.Value = &v
	return s
}

func (s *CreateFlowlogRequestTag) Validate() error {
	return dara.Validate(s)
}
