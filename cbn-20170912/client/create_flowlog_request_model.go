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
	// Generate a value for this parameter from your client. Make sure that the value is unique for each request. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the flow log.
	//
	// The description can be empty or 1 to 256 characters in length. It cannot start with \\`http\\://\\` or \\`https\\://\\`.
	//
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the flow log.
	//
	// The name can be empty or 1 to 128 characters in length. It cannot start with \\`http\\://\\` or \\`https\\://\\`.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The aggregation interval for the flow log. Unit: seconds. Valid values: **60*	- and **600**. Default value: **600**.
	//
	// example:
	//
	// 600
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// A custom string of log fields for the flow log.
	//
	// The format is defined as:
	//
	// `${field 1}${field 2}${field 3}...${field n}`
	//
	// - If you do not specify this parameter, all default fields are logged.
	//
	// - If you specify this parameter, you must start the string with `${srcaddr}${dstaddr}${bytes}` because these are required parameters. For more information about all supported log fields, see [Configure a flow log](https://help.aliyun.com/document_detail/339822.html).
	//
	// example:
	//
	// ${srcaddr}${dstaddr}${bytes}
	LogFormatString *string `json:"LogFormatString,omitempty" xml:"LogFormatString,omitempty"`
	// The Logstore that is used to store the captured traffic.
	//
	// - If you have already created a Logstore in the current region, enter the name of the Logstore.
	//
	// - If you have not created a Logstore in the current region, you can specify a custom name for the Logstore. The system automatically creates the Logstore.
	//
	//   The name of the Logstore must meet the following requirements:
	//
	//   - The name must be unique within the same project.
	//
	//   - It can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	//   - It must start and end with a lowercase letter or a digit.
	//
	//   - It must be 3 to 63 characters in length.
	//
	// example:
	//
	// flowlog-logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The project that is used to store the captured traffic.
	//
	// - If you have already created a project in the current region, enter the name of the project.
	//
	// - If you have not created a project in the current region, you can specify a custom name for the project. The system automatically creates the project.
	//
	//   A project name must be globally unique within an Alibaba Cloud region and cannot be modified after the project is created. The name must meet the following requirements:
	//
	//   - The name must be globally unique.
	//
	//   - It can contain only lowercase letters, digits, and hyphens (-).
	//
	//   - It must start and end with a lowercase letter or a digit.
	//
	//   - It must be 3 to 63 characters in length.
	//
	// example:
	//
	// flowlog-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID of the flow log.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag.
	//
	// You can add up to 20 tags in each call.
	Tag []*CreateFlowlogRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC connection, VPN connection, VBR connection, ECR connection, or inter-region connection.
	//
	// If you want to configure a flow log for a transit router instance, do not specify this parameter.
	//
	// example:
	//
	// tr-attach-r6g0m3epjehw57****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the transit router instance.
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

type CreateFlowlogRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https:// `.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be an empty string or a string of up to 128 characters. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https:// `.
	//
	// Each tag key must have a unique tag value. You can specify up to 20 tag values.
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
