// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListCloudResourcesResponseBody
	GetCurrentPage() *int32
	SetData(v []*ListCloudResourcesResponseBodyData) *ListCloudResourcesResponseBody
	GetData() []*ListCloudResourcesResponseBodyData
	SetRequestId(v string) *ListCloudResourcesResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListCloudResourcesResponseBody
	GetShowSize() *int32
	SetTotal(v int64) *ListCloudResourcesResponseBody
	GetTotal() *int64
}

type ListCloudResourcesResponseBody struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data returned for the request.
	Data []*ListCloudResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 440
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCloudResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudResourcesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudResourcesResponseBody) GetData() []*ListCloudResourcesResponseBodyData {
	return s.Data
}

func (s *ListCloudResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudResourcesResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCloudResourcesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListCloudResourcesResponseBody) SetCurrentPage(v int32) *ListCloudResourcesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudResourcesResponseBody) SetData(v []*ListCloudResourcesResponseBodyData) *ListCloudResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListCloudResourcesResponseBody) SetRequestId(v string) *ListCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudResourcesResponseBody) SetShowSize(v int32) *ListCloudResourcesResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCloudResourcesResponseBody) SetTotal(v int64) *ListCloudResourcesResponseBody {
	s.Total = &v
	return s
}

func (s *ListCloudResourcesResponseBody) Validate() error {
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

type ListCloudResourcesResponseBodyData struct {
	// The end date of the certificate bound to the cloud resource. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1737795520000
	CertEndTime *string `json:"CertEndTime,omitempty" xml:"CertEndTime,omitempty"`
	// The ID of the certificate bound to the cloud resource.
	//
	// example:
	//
	// 12433121
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate bound to the cloud resource.
	//
	// example:
	//
	// shop.amsaudio.cn
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The start date of the certificate bound to the cloud resource. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1706259520000
	CertStartTime *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The AccessKey ID that is used to access cloud resources.
	//
	// >  This parameter is returned only when you deploy certificates to cloud services of third-party clouds.
	//
	// example:
	//
	// 1234
	CloudAccessId *string `json:"CloudAccessId,omitempty" xml:"CloudAccessId,omitempty"`
	// The cloud service provider.
	//
	// Valid values:
	//
	// 	- Tencent
	//
	// 	- Huawei
	//
	// 	- Aws
	//
	// 	- aliyun
	//
	// example:
	//
	// aliyun
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The cloud service.
	//
	// example:
	//
	// ALB
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The region ID of the cloud service provider to which the cloud resource belongs.
	//
	// example:
	//
	// cn-hangzhou
	CloudRegion *string `json:"CloudRegion,omitempty" xml:"CloudRegion,omitempty"`
	// Indicates whether the cloud resource is the default resource. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 0
	DefaultResource *int32 `json:"DefaultResource,omitempty" xml:"DefaultResource,omitempty"`
	// The domain name bound to the cloud resource.
	//
	// example:
	//
	// www.tkgeo.ru
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether HTTPS is enabled for the cloud resource. Valid values:
	//
	// 	- **1**: yes.
	//
	// 	- **0**: no.
	//
	// example:
	//
	// 1
	EnableHttps *int32 `json:"EnableHttps,omitempty" xml:"EnableHttps,omitempty"`
	// The time when the cloud resource was created. The time is a timestamp in seconds.
	//
	// example:
	//
	// 1673423339000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the cloud resource was last modified. The time is a timestamp in seconds.
	//
	// example:
	//
	// 1696911946000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the cloud resource.
	//
	// example:
	//
	// 2356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// nlb-rv05agjc97ovm14il5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listener ID of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// lsn-jiugof6t23et66lsnc@443
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listening port of the cloud resource.
	//
	// >  This parameter is returned only when the value of CloudProduct is SLB, NLB, ALB, or GA.
	//
	// example:
	//
	// 8047
	ListenerPort *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The region ID of the cloud resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the cloud resource.
	//
	// example:
	//
	// BUY
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether an Alibaba Cloud SSL certificate is used. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// >  This parameter is required only when you deploy certificates to services of multiple clouds.
	//
	// example:
	//
	// 1
	UseSsl *int32 `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1666884372152785
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListCloudResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCloudResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloudResourcesResponseBodyData) GetCertEndTime() *string {
	return s.CertEndTime
}

func (s *ListCloudResourcesResponseBodyData) GetCertId() *int64 {
	return s.CertId
}

func (s *ListCloudResourcesResponseBodyData) GetCertName() *string {
	return s.CertName
}

func (s *ListCloudResourcesResponseBodyData) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *ListCloudResourcesResponseBodyData) GetCloudAccessId() *string {
	return s.CloudAccessId
}

func (s *ListCloudResourcesResponseBodyData) GetCloudName() *string {
	return s.CloudName
}

func (s *ListCloudResourcesResponseBodyData) GetCloudProduct() *string {
	return s.CloudProduct
}

func (s *ListCloudResourcesResponseBodyData) GetCloudRegion() *string {
	return s.CloudRegion
}

func (s *ListCloudResourcesResponseBodyData) GetDefaultResource() *int32 {
	return s.DefaultResource
}

func (s *ListCloudResourcesResponseBodyData) GetDomain() *string {
	return s.Domain
}

func (s *ListCloudResourcesResponseBodyData) GetEnableHttps() *int32 {
	return s.EnableHttps
}

func (s *ListCloudResourcesResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListCloudResourcesResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListCloudResourcesResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListCloudResourcesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudResourcesResponseBodyData) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCloudResourcesResponseBodyData) GetListenerPort() *string {
	return s.ListenerPort
}

func (s *ListCloudResourcesResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCloudResourcesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListCloudResourcesResponseBodyData) GetUseSsl() *int32 {
	return s.UseSsl
}

func (s *ListCloudResourcesResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *ListCloudResourcesResponseBodyData) SetCertEndTime(v string) *ListCloudResourcesResponseBodyData {
	s.CertEndTime = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCertId(v int64) *ListCloudResourcesResponseBodyData {
	s.CertId = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCertName(v string) *ListCloudResourcesResponseBodyData {
	s.CertName = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCertStartTime(v string) *ListCloudResourcesResponseBodyData {
	s.CertStartTime = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCloudAccessId(v string) *ListCloudResourcesResponseBodyData {
	s.CloudAccessId = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCloudName(v string) *ListCloudResourcesResponseBodyData {
	s.CloudName = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCloudProduct(v string) *ListCloudResourcesResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetCloudRegion(v string) *ListCloudResourcesResponseBodyData {
	s.CloudRegion = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetDefaultResource(v int32) *ListCloudResourcesResponseBodyData {
	s.DefaultResource = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetDomain(v string) *ListCloudResourcesResponseBodyData {
	s.Domain = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetEnableHttps(v int32) *ListCloudResourcesResponseBodyData {
	s.EnableHttps = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetGmtCreate(v string) *ListCloudResourcesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetGmtModified(v string) *ListCloudResourcesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetId(v int64) *ListCloudResourcesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetInstanceId(v string) *ListCloudResourcesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetListenerId(v string) *ListCloudResourcesResponseBodyData {
	s.ListenerId = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetListenerPort(v string) *ListCloudResourcesResponseBodyData {
	s.ListenerPort = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetRegionId(v string) *ListCloudResourcesResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetStatus(v string) *ListCloudResourcesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetUseSsl(v int32) *ListCloudResourcesResponseBodyData {
	s.UseSsl = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) SetUserId(v int64) *ListCloudResourcesResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListCloudResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
