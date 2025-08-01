// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewaySlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGatewaySlbResponseBody
	GetCode() *int32
	SetData(v []*ListGatewaySlbResponseBodyData) *ListGatewaySlbResponseBody
	GetData() []*ListGatewaySlbResponseBodyData
	SetHttpStatusCode(v int32) *ListGatewaySlbResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGatewaySlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewaySlbResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGatewaySlbResponseBody
	GetSuccess() *bool
}

type ListGatewaySlbResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data entries returned.
	Data []*ListGatewaySlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EAB345F4-3AC3-560C-B653-65717703****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewaySlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaySlbResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewaySlbResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewaySlbResponseBody) GetData() []*ListGatewaySlbResponseBodyData {
	return s.Data
}

func (s *ListGatewaySlbResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewaySlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewaySlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewaySlbResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewaySlbResponseBody) SetCode(v int32) *ListGatewaySlbResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewaySlbResponseBody) SetData(v []*ListGatewaySlbResponseBodyData) *ListGatewaySlbResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewaySlbResponseBody) SetHttpStatusCode(v int32) *ListGatewaySlbResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewaySlbResponseBody) SetMessage(v string) *ListGatewaySlbResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewaySlbResponseBody) SetRequestId(v string) *ListGatewaySlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewaySlbResponseBody) SetSuccess(v bool) *ListGatewaySlbResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewaySlbResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGatewaySlbResponseBodyData struct {
	// Indicates whether the edit operation is supported.
	//
	// example:
	//
	// false
	EditEnable *bool `json:"EditEnable,omitempty" xml:"EditEnable,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 1
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The mode of the SLB instance.
	//
	// example:
	//
	// UserSelect
	GatewaySlbMode *string `json:"GatewaySlbMode,omitempty" xml:"GatewaySlbMode,omitempty"`
	// The association status.
	//
	// example:
	//
	// Ready
	GatewaySlbStatus *string `json:"GatewaySlbStatus,omitempty" xml:"GatewaySlbStatus,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-14 14:39:16
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The port number of the HTTP virtual service group.
	//
	// example:
	//
	// 80
	HttpPort *int32 `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	// The port number of the HTTPS virtual service group.
	//
	// example:
	//
	// 443
	HttpsPort *int32 `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	// The ID of the HTTPS virtual service group.
	//
	// example:
	//
	// 353
	HttpsVServerGroupId *string `json:"HttpsVServerGroupId,omitempty" xml:"HttpsVServerGroupId,omitempty"`
	// The ID.
	//
	// example:
	//
	// ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The service weight.
	//
	// example:
	//
	// 80
	ServiceWeight *int32 `json:"ServiceWeight,omitempty" xml:"ServiceWeight,omitempty"`
	// The ID of the SLB instance.
	//
	// example:
	//
	// lb-bp1kmnli3hdpreptw2ah3
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The IP address of the SLB instance.
	//
	// example:
	//
	// 121.199.XX.XX
	SlbIp *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	// The port number of the SLB instance.
	//
	// example:
	//
	// 80,443
	SlbPort *string `json:"SlbPort,omitempty" xml:"SlbPort,omitempty"`
	SlbType *string `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
	// The description of the status.
	//
	// example:
	//
	// Associating
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The type.
	//
	// example:
	//
	// PUB_NET
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the HTTP virtual service group.
	//
	// example:
	//
	// 353
	VServerGroupId *string                                       `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	VServiceList   []*ListGatewaySlbResponseBodyDataVServiceList `json:"VServiceList,omitempty" xml:"VServiceList,omitempty" type:"Repeated"`
	VsMetaInfo     *string                                       `json:"VsMetaInfo,omitempty" xml:"VsMetaInfo,omitempty"`
}

func (s ListGatewaySlbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaySlbResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewaySlbResponseBodyData) GetEditEnable() *bool {
	return s.EditEnable
}

func (s *ListGatewaySlbResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewaySlbResponseBodyData) GetGatewaySlbMode() *string {
	return s.GatewaySlbMode
}

func (s *ListGatewaySlbResponseBodyData) GetGatewaySlbStatus() *string {
	return s.GatewaySlbStatus
}

func (s *ListGatewaySlbResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListGatewaySlbResponseBodyData) GetHttpPort() *int32 {
	return s.HttpPort
}

func (s *ListGatewaySlbResponseBodyData) GetHttpsPort() *int32 {
	return s.HttpsPort
}

func (s *ListGatewaySlbResponseBodyData) GetHttpsVServerGroupId() *string {
	return s.HttpsVServerGroupId
}

func (s *ListGatewaySlbResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListGatewaySlbResponseBodyData) GetServiceWeight() *int32 {
	return s.ServiceWeight
}

func (s *ListGatewaySlbResponseBodyData) GetSlbId() *string {
	return s.SlbId
}

func (s *ListGatewaySlbResponseBodyData) GetSlbIp() *string {
	return s.SlbIp
}

func (s *ListGatewaySlbResponseBodyData) GetSlbPort() *string {
	return s.SlbPort
}

func (s *ListGatewaySlbResponseBodyData) GetSlbType() *string {
	return s.SlbType
}

func (s *ListGatewaySlbResponseBodyData) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ListGatewaySlbResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListGatewaySlbResponseBodyData) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *ListGatewaySlbResponseBodyData) GetVServiceList() []*ListGatewaySlbResponseBodyDataVServiceList {
	return s.VServiceList
}

func (s *ListGatewaySlbResponseBodyData) GetVsMetaInfo() *string {
	return s.VsMetaInfo
}

func (s *ListGatewaySlbResponseBodyData) SetEditEnable(v bool) *ListGatewaySlbResponseBodyData {
	s.EditEnable = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetGatewayId(v string) *ListGatewaySlbResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetGatewaySlbMode(v string) *ListGatewaySlbResponseBodyData {
	s.GatewaySlbMode = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetGatewaySlbStatus(v string) *ListGatewaySlbResponseBodyData {
	s.GatewaySlbStatus = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetGmtCreate(v string) *ListGatewaySlbResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetHttpPort(v int32) *ListGatewaySlbResponseBodyData {
	s.HttpPort = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetHttpsPort(v int32) *ListGatewaySlbResponseBodyData {
	s.HttpsPort = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetHttpsVServerGroupId(v string) *ListGatewaySlbResponseBodyData {
	s.HttpsVServerGroupId = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetId(v string) *ListGatewaySlbResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetServiceWeight(v int32) *ListGatewaySlbResponseBodyData {
	s.ServiceWeight = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetSlbId(v string) *ListGatewaySlbResponseBodyData {
	s.SlbId = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetSlbIp(v string) *ListGatewaySlbResponseBodyData {
	s.SlbIp = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetSlbPort(v string) *ListGatewaySlbResponseBodyData {
	s.SlbPort = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetSlbType(v string) *ListGatewaySlbResponseBodyData {
	s.SlbType = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetStatusDesc(v string) *ListGatewaySlbResponseBodyData {
	s.StatusDesc = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetType(v string) *ListGatewaySlbResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetVServerGroupId(v string) *ListGatewaySlbResponseBodyData {
	s.VServerGroupId = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetVServiceList(v []*ListGatewaySlbResponseBodyDataVServiceList) *ListGatewaySlbResponseBodyData {
	s.VServiceList = v
	return s
}

func (s *ListGatewaySlbResponseBodyData) SetVsMetaInfo(v string) *ListGatewaySlbResponseBodyData {
	s.VsMetaInfo = &v
	return s
}

func (s *ListGatewaySlbResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListGatewaySlbResponseBodyDataVServiceList struct {
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol         *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	VServerGroupId   *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	VServerGroupName *string `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s ListGatewaySlbResponseBodyDataVServiceList) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaySlbResponseBodyDataVServiceList) GoString() string {
	return s.String()
}

func (s *ListGatewaySlbResponseBodyDataVServiceList) GetPort() *string {
	return s.Port
}

func (s *ListGatewaySlbResponseBodyDataVServiceList) GetProtocol() *string {
	return s.Protocol
}

func (s *ListGatewaySlbResponseBodyDataVServiceList) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *ListGatewaySlbResponseBodyDataVServiceList) GetVServerGroupName() *string {
	return s.VServerGroupName
}

func (s *ListGatewaySlbResponseBodyDataVServiceList) SetPort(v string) *ListGatewaySlbResponseBodyDataVServiceList {
	s.Port = &v
	return s
}

func (s *ListGatewaySlbResponseBodyDataVServiceList) SetProtocol(v string) *ListGatewaySlbResponseBodyDataVServiceList {
	s.Protocol = &v
	return s
}

func (s *ListGatewaySlbResponseBodyDataVServiceList) SetVServerGroupId(v string) *ListGatewaySlbResponseBodyDataVServiceList {
	s.VServerGroupId = &v
	return s
}

func (s *ListGatewaySlbResponseBodyDataVServiceList) SetVServerGroupName(v string) *ListGatewaySlbResponseBodyDataVServiceList {
	s.VServerGroupName = &v
	return s
}

func (s *ListGatewaySlbResponseBodyDataVServiceList) Validate() error {
	return dara.Validate(s)
}
