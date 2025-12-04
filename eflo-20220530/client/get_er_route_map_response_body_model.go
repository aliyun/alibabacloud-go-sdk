// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErRouteMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetErRouteMapResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetErRouteMapResponseBody
	GetCode() *int32
	SetContent(v *GetErRouteMapResponseBodyContent) *GetErRouteMapResponseBody
	GetContent() *GetErRouteMapResponseBodyContent
	SetMessage(v string) *GetErRouteMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetErRouteMapResponseBody
	GetRequestId() *string
}

type GetErRouteMapResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Content *GetErRouteMapResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErRouteMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *GetErRouteMapResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetErRouteMapResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetErRouteMapResponseBody) GetContent() *GetErRouteMapResponseBodyContent {
	return s.Content
}

func (s *GetErRouteMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetErRouteMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetErRouteMapResponseBody) SetAccessDeniedDetail(v string) *GetErRouteMapResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetErRouteMapResponseBody) SetCode(v int32) *GetErRouteMapResponseBody {
	s.Code = &v
	return s
}

func (s *GetErRouteMapResponseBody) SetContent(v *GetErRouteMapResponseBodyContent) *GetErRouteMapResponseBody {
	s.Content = v
	return s
}

func (s *GetErRouteMapResponseBody) SetMessage(v string) *GetErRouteMapResponseBody {
	s.Message = &v
	return s
}

func (s *GetErRouteMapResponseBody) SetRequestId(v string) *GetErRouteMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErRouteMapResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetErRouteMapResponseBodyContent struct {
	// Policy behavior; optional values:
	//
	// 	- **permit**: Allow
	//
	// 	- **deny**: Rejected
	//
	// example:
	//
	// permit
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// Policy description
	//
	// example:
	//
	// ssss
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Destination CIDR block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Lingjun HUB routing policy ID
	//
	// example:
	//
	// er-rmap-uwglhzom
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// Lingjun HUB routing policy Name
	//
	// example:
	//
	// er-rmap-wulanchabu
	ErRouteMapName *string `json:"ErRouteMapName,omitempty" xml:"ErRouteMapName,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1648085472000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1648085472000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Receive Instance ID
	//
	// example:
	//
	// vpd-x25vxrb2
	ReceptionInstanceId *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	// Receive Instance Name
	//
	// example:
	//
	// vpd-receprion
	ReceptionInstanceName *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	// The tenant to which the receiving instance belongs
	//
	// example:
	//
	// 1620939556166277
	ReceptionInstanceOwner *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	// The type of the received instance. Optional values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// example:
	//
	// VPD
	ReceptionInstanceType *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzlki4ehfse4y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the policy.
	//
	// A smaller sequence number indicates a lower priority. When a route is matched, a policy with a higher priority is preferentially matched.
	//
	// **Valid values: 1001 to 2000**
	//
	// example:
	//
	// 1001
	RouteMapNum *int32 `json:"RouteMapNum,omitempty" xml:"RouteMapNum,omitempty"`
	// The status of the cache reserve instance. Valid values:
	//
	// 	- **Available**
	//
	// 	- **Not Available**: Unavailable
	//
	// 	- **Executing**: Executing
	//
	// 	- **Deleting**: The node is being deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Release Instance ID
	//
	// example:
	//
	// vpd-xgkb2kl
	TransmissionInstanceId *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	// Release Instance Name
	//
	// example:
	//
	// vpd-transimit
	TransmissionInstanceName *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	// The tenant to which the published instance belongs
	//
	// example:
	//
	// 1620939556166277
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	// Publish instance type; optional values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// example:
	//
	// VPD
	TransmissionInstanceType *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
}

func (s GetErRouteMapResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetErRouteMapResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetErRouteMapResponseBodyContent) GetAction() *string {
	return s.Action
}

func (s *GetErRouteMapResponseBodyContent) GetDescription() *string {
	return s.Description
}

func (s *GetErRouteMapResponseBodyContent) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *GetErRouteMapResponseBodyContent) GetErId() *string {
	return s.ErId
}

func (s *GetErRouteMapResponseBodyContent) GetErRouteMapId() *string {
	return s.ErRouteMapId
}

func (s *GetErRouteMapResponseBodyContent) GetErRouteMapName() *string {
	return s.ErRouteMapName
}

func (s *GetErRouteMapResponseBodyContent) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetErRouteMapResponseBodyContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetErRouteMapResponseBodyContent) GetMessage() *string {
	return s.Message
}

func (s *GetErRouteMapResponseBodyContent) GetReceptionInstanceId() *string {
	return s.ReceptionInstanceId
}

func (s *GetErRouteMapResponseBodyContent) GetReceptionInstanceName() *string {
	return s.ReceptionInstanceName
}

func (s *GetErRouteMapResponseBodyContent) GetReceptionInstanceOwner() *string {
	return s.ReceptionInstanceOwner
}

func (s *GetErRouteMapResponseBodyContent) GetReceptionInstanceType() *string {
	return s.ReceptionInstanceType
}

func (s *GetErRouteMapResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetErRouteMapResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetErRouteMapResponseBodyContent) GetRouteMapNum() *int32 {
	return s.RouteMapNum
}

func (s *GetErRouteMapResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetErRouteMapResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetErRouteMapResponseBodyContent) GetTransmissionInstanceId() *string {
	return s.TransmissionInstanceId
}

func (s *GetErRouteMapResponseBodyContent) GetTransmissionInstanceName() *string {
	return s.TransmissionInstanceName
}

func (s *GetErRouteMapResponseBodyContent) GetTransmissionInstanceOwner() *string {
	return s.TransmissionInstanceOwner
}

func (s *GetErRouteMapResponseBodyContent) GetTransmissionInstanceType() *string {
	return s.TransmissionInstanceType
}

func (s *GetErRouteMapResponseBodyContent) SetAction(v string) *GetErRouteMapResponseBodyContent {
	s.Action = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetDescription(v string) *GetErRouteMapResponseBodyContent {
	s.Description = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetDestinationCidrBlock(v string) *GetErRouteMapResponseBodyContent {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetErId(v string) *GetErRouteMapResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetErRouteMapId(v string) *GetErRouteMapResponseBodyContent {
	s.ErRouteMapId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetErRouteMapName(v string) *GetErRouteMapResponseBodyContent {
	s.ErRouteMapName = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetGmtCreate(v string) *GetErRouteMapResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetGmtModified(v string) *GetErRouteMapResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetMessage(v string) *GetErRouteMapResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetReceptionInstanceId(v string) *GetErRouteMapResponseBodyContent {
	s.ReceptionInstanceId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetReceptionInstanceName(v string) *GetErRouteMapResponseBodyContent {
	s.ReceptionInstanceName = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetReceptionInstanceOwner(v string) *GetErRouteMapResponseBodyContent {
	s.ReceptionInstanceOwner = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetReceptionInstanceType(v string) *GetErRouteMapResponseBodyContent {
	s.ReceptionInstanceType = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetRegionId(v string) *GetErRouteMapResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetResourceGroupId(v string) *GetErRouteMapResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetRouteMapNum(v int32) *GetErRouteMapResponseBodyContent {
	s.RouteMapNum = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetStatus(v string) *GetErRouteMapResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetTenantId(v string) *GetErRouteMapResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetTransmissionInstanceId(v string) *GetErRouteMapResponseBodyContent {
	s.TransmissionInstanceId = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetTransmissionInstanceName(v string) *GetErRouteMapResponseBodyContent {
	s.TransmissionInstanceName = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetTransmissionInstanceOwner(v string) *GetErRouteMapResponseBodyContent {
	s.TransmissionInstanceOwner = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) SetTransmissionInstanceType(v string) *GetErRouteMapResponseBodyContent {
	s.TransmissionInstanceType = &v
	return s
}

func (s *GetErRouteMapResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
