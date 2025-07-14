// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllSwimmingLaneGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllSwimmingLaneGroupsResponseBody
	GetCode() *string
	SetData(v []*ListAllSwimmingLaneGroupsResponseBodyData) *ListAllSwimmingLaneGroupsResponseBody
	GetData() []*ListAllSwimmingLaneGroupsResponseBodyData
	SetErrorCode(v string) *ListAllSwimmingLaneGroupsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListAllSwimmingLaneGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllSwimmingLaneGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllSwimmingLaneGroupsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListAllSwimmingLaneGroupsResponseBody
	GetTraceId() *string
}

type ListAllSwimmingLaneGroupsResponseBody struct {
	// example:
	//
	// 200
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListAllSwimmingLaneGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 30375C38-F4ED-4135-A0AE-5C75DC7F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// ac1a0b2215622246421415014e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListAllSwimmingLaneGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLaneGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLaneGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllSwimmingLaneGroupsResponseBody) GetData() []*ListAllSwimmingLaneGroupsResponseBodyData {
	return s.Data
}

func (s *ListAllSwimmingLaneGroupsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAllSwimmingLaneGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllSwimmingLaneGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllSwimmingLaneGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllSwimmingLaneGroupsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListAllSwimmingLaneGroupsResponseBody) SetCode(v string) *ListAllSwimmingLaneGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBody) SetData(v []*ListAllSwimmingLaneGroupsResponseBodyData) *ListAllSwimmingLaneGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBody) SetErrorCode(v string) *ListAllSwimmingLaneGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBody) SetMessage(v string) *ListAllSwimmingLaneGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBody) SetRequestId(v string) *ListAllSwimmingLaneGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBody) SetSuccess(v bool) *ListAllSwimmingLaneGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBody) SetTraceId(v string) *ListAllSwimmingLaneGroupsResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAllSwimmingLaneGroupsResponseBodyData struct {
	AppIds []*string                                        `json:"AppIds,omitempty" xml:"AppIds,omitempty" type:"Repeated"`
	Apps   []*ListAllSwimmingLaneGroupsResponseBodyDataApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	CanaryModel *int32                                             `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	EntryApp    *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp `json:"EntryApp,omitempty" xml:"EntryApp,omitempty" type:"Struct"`
	// example:
	//
	// mse_ingresspost-cn-axc49******
	EntryAppId *string `json:"EntryAppId,omitempty" xml:"EntryAppId,omitempty"`
	// example:
	//
	// mse-gw
	EntryAppType *string `json:"EntryAppType,omitempty" xml:"EntryAppType,omitempty"`
	// example:
	//
	// 2074
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// mse-test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// sae-test
	MseNamespaceId *string `json:"MseNamespaceId,omitempty" xml:"MseNamespaceId,omitempty"`
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// 2
	SwimVersion *string `json:"SwimVersion,omitempty" xml:"SwimVersion,omitempty"`
}

func (s ListAllSwimmingLaneGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLaneGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) GetAppIds() []*string {
	return s.AppIds
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) GetApps() []*ListAllSwimmingLaneGroupsResponseBodyDataApps {
	return s.Apps
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) GetCanaryModel() *int32 {
	return s.CanaryModel
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) GetEntryApp() *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp {
	return s.EntryApp
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) GetEntryAppId() *string {
	return s.EntryAppId
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) GetEntryAppType() *string {
	return s.EntryAppType
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) GetMseNamespaceId() *string {
	return s.MseNamespaceId
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) GetSwimVersion() *string {
	return s.SwimVersion
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) SetAppIds(v []*string) *ListAllSwimmingLaneGroupsResponseBodyData {
	s.AppIds = v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) SetApps(v []*ListAllSwimmingLaneGroupsResponseBodyDataApps) *ListAllSwimmingLaneGroupsResponseBodyData {
	s.Apps = v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) SetCanaryModel(v int32) *ListAllSwimmingLaneGroupsResponseBodyData {
	s.CanaryModel = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) SetEntryApp(v *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) *ListAllSwimmingLaneGroupsResponseBodyData {
	s.EntryApp = v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) SetEntryAppId(v string) *ListAllSwimmingLaneGroupsResponseBodyData {
	s.EntryAppId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) SetEntryAppType(v string) *ListAllSwimmingLaneGroupsResponseBodyData {
	s.EntryAppType = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) SetGroupId(v int64) *ListAllSwimmingLaneGroupsResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) SetGroupName(v string) *ListAllSwimmingLaneGroupsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) SetMseNamespaceId(v string) *ListAllSwimmingLaneGroupsResponseBodyData {
	s.MseNamespaceId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) SetNamespaceId(v string) *ListAllSwimmingLaneGroupsResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) SetSwimVersion(v string) *ListAllSwimmingLaneGroupsResponseBodyData {
	s.SwimVersion = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAllSwimmingLaneGroupsResponseBodyDataApps struct {
	// example:
	//
	// f5aad0d0-3e56-44cd-8199-9887a0******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// mse-cn-53y49******
	MseAppId *string `json:"MseAppId,omitempty" xml:"MseAppId,omitempty"`
	// example:
	//
	// demo
	MseAppName *string `json:"MseAppName,omitempty" xml:"MseAppName,omitempty"`
	// example:
	//
	// 6733e538-d52f-48e6-91a4-192f91******
	MseNamespaceId *string `json:"MseNamespaceId,omitempty" xml:"MseNamespaceId,omitempty"`
}

func (s ListAllSwimmingLaneGroupsResponseBodyDataApps) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLaneGroupsResponseBodyDataApps) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataApps) GetAppId() *string {
	return s.AppId
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataApps) GetAppName() *string {
	return s.AppName
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataApps) GetMseAppId() *string {
	return s.MseAppId
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataApps) GetMseAppName() *string {
	return s.MseAppName
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataApps) GetMseNamespaceId() *string {
	return s.MseNamespaceId
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataApps) SetAppId(v string) *ListAllSwimmingLaneGroupsResponseBodyDataApps {
	s.AppId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataApps) SetAppName(v string) *ListAllSwimmingLaneGroupsResponseBodyDataApps {
	s.AppName = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataApps) SetMseAppId(v string) *ListAllSwimmingLaneGroupsResponseBodyDataApps {
	s.MseAppId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataApps) SetMseAppName(v string) *ListAllSwimmingLaneGroupsResponseBodyDataApps {
	s.MseAppName = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataApps) SetMseNamespaceId(v string) *ListAllSwimmingLaneGroupsResponseBodyDataApps {
	s.MseNamespaceId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataApps) Validate() error {
	return dara.Validate(s)
}

type ListAllSwimmingLaneGroupsResponseBodyDataEntryApp struct {
	// example:
	//
	// 09805e5d-9b8d-42cd-9442-03c498******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// mse
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// mse-cn-53y49******
	MseAppId *string `json:"MseAppId,omitempty" xml:"MseAppId,omitempty"`
	// example:
	//
	// test
	MseAppName *string `json:"MseAppName,omitempty" xml:"MseAppName,omitempty"`
	// example:
	//
	// demo
	MseNamespaceId *string `json:"MseNamespaceId,omitempty" xml:"MseNamespaceId,omitempty"`
}

func (s ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) GetAppId() *string {
	return s.AppId
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) GetAppName() *string {
	return s.AppName
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) GetAppType() *string {
	return s.AppType
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) GetMseAppId() *string {
	return s.MseAppId
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) GetMseAppName() *string {
	return s.MseAppName
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) GetMseNamespaceId() *string {
	return s.MseNamespaceId
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) SetAppId(v string) *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp {
	s.AppId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) SetAppName(v string) *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp {
	s.AppName = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) SetAppType(v string) *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp {
	s.AppType = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) SetMseAppId(v string) *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp {
	s.MseAppId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) SetMseAppName(v string) *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp {
	s.MseAppName = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) SetMseNamespaceId(v string) *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp {
	s.MseNamespaceId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponseBodyDataEntryApp) Validate() error {
	return dara.Validate(s)
}
