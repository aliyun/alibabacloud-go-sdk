// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v *ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody
	GetApplications() *ListApplicationsResponseBodyApplications
	SetPageNumber(v int32) *ListApplicationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListApplicationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApplicationsResponseBody
	GetTotalCount() *int32
}

type ListApplicationsResponseBody struct {
	// Details about applications.
	Applications *ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 50373E71-7710-4620-8AAB-133CCE49451C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 49
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) GetApplications() *ListApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationsResponseBody) SetApplications(v *ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBody) SetPageNumber(v int32) *ListApplicationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsResponseBody) SetPageSize(v int32) *ListApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) SetTotalCount(v int32) *ListApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsResponseBodyApplications struct {
	Application []*ListApplicationsResponseBodyApplicationsApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) GetApplication() []*ListApplicationsResponseBodyApplicationsApplication {
	return s.Application
}

func (s *ListApplicationsResponseBodyApplications) SetApplication(v []*ListApplicationsResponseBodyApplicationsApplication) *ListApplicationsResponseBodyApplications {
	s.Application = v
	return s
}

func (s *ListApplicationsResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsResponseBodyApplicationsApplication struct {
	// Details about the application.
	AppList *ListApplicationsResponseBodyApplicationsApplicationAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Struct"`
	// The name of the cluster.
	//
	// example:
	//
	// ay-ads-hz-h
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s ListApplicationsResponseBodyApplicationsApplication) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplication) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetAppList() *ListApplicationsResponseBodyApplicationsApplicationAppList {
	return s.AppList
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAppList(v *ListApplicationsResponseBodyApplicationsApplicationAppList) *ListApplicationsResponseBodyApplicationsApplication {
	s.AppList = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetClusterName(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.ClusterName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsResponseBodyApplicationsApplicationAppList struct {
	App []*ListApplicationsResponseBodyApplicationsApplicationAppListApp `json:"App,omitempty" xml:"App,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppList) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppList) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppList) GetApp() []*ListApplicationsResponseBodyApplicationsApplicationAppListApp {
	return s.App
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppList) SetApp(v []*ListApplicationsResponseBodyApplicationsApplicationAppListApp) *ListApplicationsResponseBodyApplicationsApplicationAppList {
	s.App = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppList) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsResponseBodyApplicationsApplicationAppListApp struct {
	// The ID of the application.
	//
	// example:
	//
	// e76f8985-7965-41fc-925b-9648bb6bf650
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The information about the application, such as the resource specification, parameter configuration, and resources.
	//
	// example:
	//
	// {
	//
	//     "AppStatus":{
	//
	//         "Phase":"RUNNING",
	//
	//         "StatusDescrip":"{\\"status\\":\\"UPDATE_SUCCESS\\",\\"step\\":\\"DONE\\",\\"descrip\\":\\"update to version:1022 success\\",\\"start_time\\":\\"2022-03-01 16:18:22\\"}",
	//
	//         "UpdateTime":"2022-03-01 16:18:22",
	//
	//         "OrderStatus":null
	//
	//     },
	//
	//     "ResourceAttribute":{
	//
	//         "NetSecurityInfo":null,
	//
	//         "InitConfig":null,
	//
	//         "InventoryType":"Ens",
	//
	//         "InstanceSpec":"ens.gi6s-c12g1.large",
	//
	//         "SystemDiskSize":100,
	//
	//         "DataDiskSize":0,
	//
	//         "BandwithOut":5000,
	//
	//         "SchedulingStrategy":"Disperse",
	//
	//         "ImageId":"m-5or73kzkuxytv7hh6wxr6yc5q",
	//
	//         "ResourceType":"Linux",
	//
	//         "AreaLevel":"National",
	//
	//         "IpType":"PublicIP"
	//
	//     },
	//
	//     "WorkloadAttribute":[
	//
	//         {
	//
	//             "Name":"andorid",
	//
	//             "Count":15,
	//
	//             "ServiceConfig":{
	//
	//                 "PortsBindConfig":{
	//
	//                     "NetServiceContainer":"uravi-service",
	//
	//                     "Ports":[
	//
	//                         {
	//
	//                             "Protocol":"TCP",
	//
	//                             "BindType":"Mapping",
	//
	//                             "StartNodePorts":"31000-31009",
	//
	//                             "ContainerPorts":"4440-4449"
	//
	//                         },
	//
	//                         {
	//
	//                             "Protocol":"TCP",
	//
	//                             "BindType":"PassThrough",
	//
	//                             "StartNodePorts":"59000-59000",
	//
	//                             "ContainerPorts":"59000-59000"
	//
	//                         },
	//
	//                         {
	//
	//                             "Protocol":"UDP",
	//
	//                             "BindType":"PassThrough",
	//
	//                             "StartNodePorts":"40001-40010",
	//
	//                             "ContainerPorts":"40001-40010"
	//
	//                         }
	//
	//                     ]
	//
	//                 },
	//
	//                 "ServiceContainerName":"android"
	//
	//             }
	//
	//         },
	//
	//         {
	//
	//             "Name":"coturn",
	//
	//             "Count":1,
	//
	//             "ServiceConfig":{
	//
	//                 "PortsBindConfig":{
	//
	//                     "NetServiceContainer":"coturn",
	//
	//                     "Ports":[
	//
	//                         {
	//
	//                             "Protocol":"TCP",
	//
	//                             "BindType":"PassThrough",
	//
	//                             "StartNodePorts":"3478-3478",
	//
	//                             "ContainerPorts":"3478-3478"
	//
	//                         },
	//
	//                         {
	//
	//                             "Protocol":"UDP",
	//
	//                             "BindType":"PassThrough",
	//
	//                             "StartNodePorts":"3478-3478",
	//
	//                             "ContainerPorts":"3478-3478"
	//
	//                         }
	//
	//                     ]
	//
	//                 },
	//
	//                 "ServiceContainerName":"coturn"
	//
	//             }
	//
	//         },
	//
	//         {
	//
	//             "Name":"aic-manager",
	//
	//             "Count":1,
	//
	//             "ServiceConfig":null
	//
	//         }
	//
	//     ]
	//
	// }
	AppInfo *string `json:"AppInfo,omitempty" xml:"AppInfo,omitempty"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppListApp) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationAppListApp) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) GetAppInfo() *string {
	return s.AppInfo
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) SetAppId(v string) *ListApplicationsResponseBodyApplicationsApplicationAppListApp {
	s.AppId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) SetAppInfo(v string) *ListApplicationsResponseBodyApplicationsApplicationAppListApp {
	s.AppInfo = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationAppListApp) Validate() error {
	return dara.Validate(s)
}
