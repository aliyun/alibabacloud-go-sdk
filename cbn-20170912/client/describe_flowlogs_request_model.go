// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowlogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeFlowlogsRequest
	GetCenId() *string
	SetClientToken(v string) *DescribeFlowlogsRequest
	GetClientToken() *string
	SetDescription(v string) *DescribeFlowlogsRequest
	GetDescription() *string
	SetFlowLogId(v string) *DescribeFlowlogsRequest
	GetFlowLogId() *string
	SetFlowLogName(v string) *DescribeFlowlogsRequest
	GetFlowLogName() *string
	SetFlowLogVersion(v string) *DescribeFlowlogsRequest
	GetFlowLogVersion() *string
	SetInterval(v int32) *DescribeFlowlogsRequest
	GetInterval() *int32
	SetLogStoreName(v string) *DescribeFlowlogsRequest
	GetLogStoreName() *string
	SetOwnerAccount(v string) *DescribeFlowlogsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeFlowlogsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeFlowlogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFlowlogsRequest
	GetPageSize() *int32
	SetProjectName(v string) *DescribeFlowlogsRequest
	GetProjectName() *string
	SetRegionId(v string) *DescribeFlowlogsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeFlowlogsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeFlowlogsRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeFlowlogsRequest
	GetStatus() *string
	SetTag(v []*DescribeFlowlogsRequestTag) *DescribeFlowlogsRequest
	GetTag() []*DescribeFlowlogsRequestTag
	SetTransitRouterAttachmentId(v string) *DescribeFlowlogsRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterId(v string) *DescribeFlowlogsRequest
	GetTransitRouterId() *string
}

type DescribeFlowlogsRequest struct {
	// The Cloud Enterprise Network (CEN) instance ID.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId value as the ClientToken value. The RequestId value may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the flow log.
	//
	// The description can be empty or 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The flow log ID.
	//
	// example:
	//
	// flowlog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The name of the flow log.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The version of the flow log.
	//
	// When a flow log is created, the latest version supported by the system is automatically used. The current version is **3**.
	//
	// example:
	//
	// 3
	FlowLogVersion *string `json:"FlowLogVersion,omitempty" xml:"FlowLogVersion,omitempty"`
	// The capture window duration of the flow log. Unit: seconds. Valid values: **60*	- or **600**. Default value: **600**.
	//
	// example:
	//
	// 600
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Logstore that stores the captured traffic.
	//
	// The Logstore name must be 3 to 63 characters in length, and must start and end with a lowercase letter or digit. It can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// FlowLogStore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paging queries. Minimum value: **1**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the project that stores the captured traffic.
	//
	// The project name must be 3 to 63 characters in length, and must start and end with a lowercase letter or digit. It can contain only lowercase letters, digits, and hyphens (-).
	//
	// example:
	//
	// FlowLogProject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID of the flow log.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the flow log. Valid values:
	//
	// - **Active**: activated.
	//
	// - **Inactive**: not activated.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags at a time.
	Tag []*DescribeFlowlogsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The network instance connection ID.
	//
	// example:
	//
	// tr-attach-qieks13jnt1cchy***
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The transit router instance ID.
	//
	// example:
	//
	// tr-uf654ttymmljlvh2x****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DescribeFlowlogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowlogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeFlowlogsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeFlowlogsRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeFlowlogsRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DescribeFlowlogsRequest) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *DescribeFlowlogsRequest) GetFlowLogVersion() *string {
	return s.FlowLogVersion
}

func (s *DescribeFlowlogsRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeFlowlogsRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeFlowlogsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeFlowlogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeFlowlogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFlowlogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFlowlogsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeFlowlogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowlogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeFlowlogsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeFlowlogsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeFlowlogsRequest) GetTag() []*DescribeFlowlogsRequestTag {
	return s.Tag
}

func (s *DescribeFlowlogsRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DescribeFlowlogsRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeFlowlogsRequest) SetCenId(v string) *DescribeFlowlogsRequest {
	s.CenId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetClientToken(v string) *DescribeFlowlogsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetDescription(v string) *DescribeFlowlogsRequest {
	s.Description = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetFlowLogId(v string) *DescribeFlowlogsRequest {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetFlowLogName(v string) *DescribeFlowlogsRequest {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetFlowLogVersion(v string) *DescribeFlowlogsRequest {
	s.FlowLogVersion = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetInterval(v int32) *DescribeFlowlogsRequest {
	s.Interval = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetLogStoreName(v string) *DescribeFlowlogsRequest {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetOwnerAccount(v string) *DescribeFlowlogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetOwnerId(v int64) *DescribeFlowlogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetPageNumber(v int32) *DescribeFlowlogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetPageSize(v int32) *DescribeFlowlogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetProjectName(v string) *DescribeFlowlogsRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetRegionId(v string) *DescribeFlowlogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetResourceOwnerAccount(v string) *DescribeFlowlogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetResourceOwnerId(v int64) *DescribeFlowlogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetStatus(v string) *DescribeFlowlogsRequest {
	s.Status = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetTag(v []*DescribeFlowlogsRequestTag) *DescribeFlowlogsRequest {
	s.Tag = v
	return s
}

func (s *DescribeFlowlogsRequest) SetTransitRouterAttachmentId(v string) *DescribeFlowlogsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DescribeFlowlogsRequest) SetTransitRouterId(v string) *DescribeFlowlogsRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeFlowlogsRequest) Validate() error {
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

type DescribeFlowlogsRequestTag struct {
	// The tag key of the resource.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// The tag value can be an empty string or up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFlowlogsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowlogsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeFlowlogsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeFlowlogsRequestTag) SetKey(v string) *DescribeFlowlogsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeFlowlogsRequestTag) SetValue(v string) *DescribeFlowlogsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeFlowlogsRequestTag) Validate() error {
	return dara.Validate(s)
}
