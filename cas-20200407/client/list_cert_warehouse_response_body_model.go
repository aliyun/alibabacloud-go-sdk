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
	// The list of certificate repositories.
	CertWarehouseList []*ListCertWarehouseResponseBodyCertWarehouseList `json:"CertWarehouseList,omitempty" xml:"CertWarehouseList,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries per page. Default value: 50.
	//
	// example:
	//
	// 50
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries.
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
	if s.CertWarehouseList != nil {
		for _, item := range s.CertWarehouseList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCertWarehouseResponseBodyCertWarehouseList struct {
	// The expiration time, in timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1665819958000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The sales instance.
	//
	// example:
	//
	// 14dcc8afc7578e1f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the repository has expired. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// false
	IsExpired *bool `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// The repository name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The PCA instance.
	//
	// example:
	//
	// 14dcc8afc7578e1f
	PcaInstanceId *string `json:"PcaInstanceId,omitempty" xml:"PcaInstanceId,omitempty"`
	// Qps。
	//
	// example:
	//
	// 10
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The repository type. Valid values:
	//
	// - **uploadCA**: an uploaded CA certificate that contains a complete certificate chain.
	//
	// - **uploadPCA**: an uploaded certificate, including a self-signed certificate, a certificate issued by a third party, or a certificate issued by Alibaba Cloud.
	//
	// - **aliyunPCA**: an Alibaba Cloud PCA certificate.
	//
	// example:
	//
	// aliyunPCA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The repository ID.
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
