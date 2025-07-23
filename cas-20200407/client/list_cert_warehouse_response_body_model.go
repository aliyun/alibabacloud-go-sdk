// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertWarehouseList(v []*ListCertWarehouseResponseBodyCertWarehouseList) *ListCertWarehouseResponseBody
	GetCertWarehouseList() []*ListCertWarehouseResponseBodyCertWarehouseList
	SetCurrentPage(v int64) *ListCertWarehouseResponseBody
	GetCurrentPage() *int64
	SetRequestId(v string) *ListCertWarehouseResponseBody
	GetRequestId() *string
	SetShowSize(v int64) *ListCertWarehouseResponseBody
	GetShowSize() *int64
	SetTotalCount(v int64) *ListCertWarehouseResponseBody
	GetTotalCount() *int64
}

type ListCertWarehouseResponseBody struct {
	// The certificate application repositories.
	CertWarehouseList []*ListCertWarehouseResponseBodyCertWarehouseList `json:"CertWarehouseList,omitempty" xml:"CertWarehouseList,omitempty" type:"Repeated"`
	// The page number of the returned page. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned per page. Default value: 50.
	//
	// example:
	//
	// 50
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCertWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertWarehouseResponseBody) GetCertWarehouseList() []*ListCertWarehouseResponseBodyCertWarehouseList {
	return s.CertWarehouseList
}

func (s *ListCertWarehouseResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListCertWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCertWarehouseResponseBody) GetShowSize() *int64 {
	return s.ShowSize
}

func (s *ListCertWarehouseResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCertWarehouseResponseBody) SetCertWarehouseList(v []*ListCertWarehouseResponseBodyCertWarehouseList) *ListCertWarehouseResponseBody {
	s.CertWarehouseList = v
	return s
}

func (s *ListCertWarehouseResponseBody) SetCurrentPage(v int64) *ListCertWarehouseResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCertWarehouseResponseBody) SetRequestId(v string) *ListCertWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertWarehouseResponseBody) SetShowSize(v int64) *ListCertWarehouseResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCertWarehouseResponseBody) SetTotalCount(v int64) *ListCertWarehouseResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCertWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCertWarehouseResponseBodyCertWarehouseList struct {
	// The timestamp when the certificate application repository expires. Unit: milliseconds.
	//
	// example:
	//
	// 1665819958000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID of the certificate application repository.
	//
	// example:
	//
	// 14dcc8afc7578e1f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the certificate application repository has expired. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsExpired *bool `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// The name of the certificate application repository.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The instance ID of the private CA.
	//
	// example:
	//
	// 14dcc8afc7578e1f
	PcaInstanceId *string `json:"PcaInstanceId,omitempty" xml:"PcaInstanceId,omitempty"`
	// The queries per second (QPS).
	//
	// example:
	//
	// 10
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The type of the certificate application repository. Valid values:
	//
	// 	- **ssl**: certificate application repository of SSL certificates
	//
	// 	- **uploadPCA**: certificate application repository of uploaded private certificates
	//
	// 	- **free**: certificate application repository of free certificates, available only on the China site (aliyun.com)
	//
	// 	- **aliyunPCA**: certificate application repository of private certificates purchased from Alibaba Cloud Private Certificate Authority (PCA), available only on the China site (aliyun.com)
	//
	// 	- **disable**: disabled certificate application repository
	//
	// example:
	//
	// aliyunPCA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the certificate application repository.
	//
	// example:
	//
	// 1
	WhId *int64 `json:"WhId,omitempty" xml:"WhId,omitempty"`
}

func (s ListCertWarehouseResponseBodyCertWarehouseList) String() string {
	return dara.Prettify(s)
}

func (s ListCertWarehouseResponseBodyCertWarehouseList) GoString() string {
	return s.String()
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) GetIsExpired() *bool {
	return s.IsExpired
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) GetName() *string {
	return s.Name
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) GetPcaInstanceId() *string {
	return s.PcaInstanceId
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) GetQps() *int64 {
	return s.Qps
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) GetType() *string {
	return s.Type
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) GetWhId() *int64 {
	return s.WhId
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetEndTime(v int64) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.EndTime = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetInstanceId(v string) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.InstanceId = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetIsExpired(v bool) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.IsExpired = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetName(v string) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.Name = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetPcaInstanceId(v string) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.PcaInstanceId = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetQps(v int64) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.Qps = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetType(v string) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.Type = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) SetWhId(v int64) *ListCertWarehouseResponseBodyCertWarehouseList {
	s.WhId = &v
	return s
}

func (s *ListCertWarehouseResponseBodyCertWarehouseList) Validate() error {
	return dara.Validate(s)
}
