// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErRouteMapsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListErRouteMapsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListErRouteMapsResponseBody
	GetCode() *int32
	SetContent(v *ListErRouteMapsResponseBodyContent) *ListErRouteMapsResponseBody
	GetContent() *ListErRouteMapsResponseBodyContent
	SetMessage(v string) *ListErRouteMapsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListErRouteMapsResponseBody
	GetRequestId() *string
}

type ListErRouteMapsResponseBody struct {
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
	// The response parameters.
	Content *ListErRouteMapsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErRouteMapsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListErRouteMapsResponseBody) GoString() string {
	return s.String()
}

func (s *ListErRouteMapsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListErRouteMapsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListErRouteMapsResponseBody) GetContent() *ListErRouteMapsResponseBodyContent {
	return s.Content
}

func (s *ListErRouteMapsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListErRouteMapsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListErRouteMapsResponseBody) SetAccessDeniedDetail(v string) *ListErRouteMapsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListErRouteMapsResponseBody) SetCode(v int32) *ListErRouteMapsResponseBody {
	s.Code = &v
	return s
}

func (s *ListErRouteMapsResponseBody) SetContent(v *ListErRouteMapsResponseBodyContent) *ListErRouteMapsResponseBody {
	s.Content = v
	return s
}

func (s *ListErRouteMapsResponseBody) SetMessage(v string) *ListErRouteMapsResponseBody {
	s.Message = &v
	return s
}

func (s *ListErRouteMapsResponseBody) SetRequestId(v string) *ListErRouteMapsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListErRouteMapsResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListErRouteMapsResponseBodyContent struct {
	// routing policy information list
	Data []*ListErRouteMapsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListErRouteMapsResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListErRouteMapsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListErRouteMapsResponseBodyContent) GetData() []*ListErRouteMapsResponseBodyContentData {
	return s.Data
}

func (s *ListErRouteMapsResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListErRouteMapsResponseBodyContent) SetData(v []*ListErRouteMapsResponseBodyContentData) *ListErRouteMapsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListErRouteMapsResponseBodyContent) SetTotal(v int64) *ListErRouteMapsResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContent) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListErRouteMapsResponseBodyContentData struct {
	// Policy behavior; optional values:
	//
	// 	- **permit**: Allow
	//
	// 	- **deny**: Prohibited
	//
	// example:
	//
	// permit
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1601176751000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Policy description
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Destination CIDR Block
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
	// routing policy ID
	//
	// example:
	//
	// er-rmap-uwglhzom
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1601176751000
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
	// vpd-9rgxqazc
	ReceptionInstanceId *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	// Receive Instance Name
	//
	// example:
	//
	// vpd-reception
	ReceptionInstanceName *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	// The tenant to which the receiving instance belongs
	//
	// example:
	//
	// 1620939556166277
	ReceptionInstanceOwner *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	// The type of the received instance. Possible values:
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
	// rg-aek2l4sq6l7unhi
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
	// Status The status of the instance. Valid values:
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
	// vpd-8rgvqazb
	TransmissionInstanceId *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	// Release Instance Name
	//
	// example:
	//
	// vpd-transmit
	TransmissionInstanceName *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	// The tenant to which the published instance belongs
	//
	// example:
	//
	// 1620939556166277
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	// The type of the published instance. Possible values:
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

func (s ListErRouteMapsResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListErRouteMapsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListErRouteMapsResponseBodyContentData) GetAction() *string {
	return s.Action
}

func (s *ListErRouteMapsResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListErRouteMapsResponseBodyContentData) GetDescription() *string {
	return s.Description
}

func (s *ListErRouteMapsResponseBodyContentData) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListErRouteMapsResponseBodyContentData) GetErId() *string {
	return s.ErId
}

func (s *ListErRouteMapsResponseBodyContentData) GetErRouteMapId() *string {
	return s.ErRouteMapId
}

func (s *ListErRouteMapsResponseBodyContentData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListErRouteMapsResponseBodyContentData) GetMessage() *string {
	return s.Message
}

func (s *ListErRouteMapsResponseBodyContentData) GetReceptionInstanceId() *string {
	return s.ReceptionInstanceId
}

func (s *ListErRouteMapsResponseBodyContentData) GetReceptionInstanceName() *string {
	return s.ReceptionInstanceName
}

func (s *ListErRouteMapsResponseBodyContentData) GetReceptionInstanceOwner() *string {
	return s.ReceptionInstanceOwner
}

func (s *ListErRouteMapsResponseBodyContentData) GetReceptionInstanceType() *string {
	return s.ReceptionInstanceType
}

func (s *ListErRouteMapsResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListErRouteMapsResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListErRouteMapsResponseBodyContentData) GetRouteMapNum() *int32 {
	return s.RouteMapNum
}

func (s *ListErRouteMapsResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListErRouteMapsResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListErRouteMapsResponseBodyContentData) GetTransmissionInstanceId() *string {
	return s.TransmissionInstanceId
}

func (s *ListErRouteMapsResponseBodyContentData) GetTransmissionInstanceName() *string {
	return s.TransmissionInstanceName
}

func (s *ListErRouteMapsResponseBodyContentData) GetTransmissionInstanceOwner() *string {
	return s.TransmissionInstanceOwner
}

func (s *ListErRouteMapsResponseBodyContentData) GetTransmissionInstanceType() *string {
	return s.TransmissionInstanceType
}

func (s *ListErRouteMapsResponseBodyContentData) SetAction(v string) *ListErRouteMapsResponseBodyContentData {
	s.Action = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetCreateTime(v string) *ListErRouteMapsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetDescription(v string) *ListErRouteMapsResponseBodyContentData {
	s.Description = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetDestinationCidrBlock(v string) *ListErRouteMapsResponseBodyContentData {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetErId(v string) *ListErRouteMapsResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetErRouteMapId(v string) *ListErRouteMapsResponseBodyContentData {
	s.ErRouteMapId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetGmtModified(v string) *ListErRouteMapsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetMessage(v string) *ListErRouteMapsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetReceptionInstanceId(v string) *ListErRouteMapsResponseBodyContentData {
	s.ReceptionInstanceId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetReceptionInstanceName(v string) *ListErRouteMapsResponseBodyContentData {
	s.ReceptionInstanceName = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetReceptionInstanceOwner(v string) *ListErRouteMapsResponseBodyContentData {
	s.ReceptionInstanceOwner = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetReceptionInstanceType(v string) *ListErRouteMapsResponseBodyContentData {
	s.ReceptionInstanceType = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetRegionId(v string) *ListErRouteMapsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetResourceGroupId(v string) *ListErRouteMapsResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetRouteMapNum(v int32) *ListErRouteMapsResponseBodyContentData {
	s.RouteMapNum = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetStatus(v string) *ListErRouteMapsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetTenantId(v string) *ListErRouteMapsResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetTransmissionInstanceId(v string) *ListErRouteMapsResponseBodyContentData {
	s.TransmissionInstanceId = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetTransmissionInstanceName(v string) *ListErRouteMapsResponseBodyContentData {
	s.TransmissionInstanceName = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetTransmissionInstanceOwner(v string) *ListErRouteMapsResponseBodyContentData {
	s.TransmissionInstanceOwner = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) SetTransmissionInstanceType(v string) *ListErRouteMapsResponseBodyContentData {
	s.TransmissionInstanceType = &v
	return s
}

func (s *ListErRouteMapsResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}
