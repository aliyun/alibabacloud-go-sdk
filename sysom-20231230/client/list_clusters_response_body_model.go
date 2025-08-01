// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListClustersResponseBody
	GetRequestId() *string
	SetCode(v string) *ListClustersResponseBody
	GetCode() *string
	SetData(v []*ListClustersResponseBodyData) *ListClustersResponseBody
	GetData() []*ListClustersResponseBodyData
	SetMessage(v string) *ListClustersResponseBody
	GetMessage() *string
	SetTotal(v int64) *ListClustersResponseBody
	GetTotal() *int64
}

type ListClustersResponseBody struct {
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListClustersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 64
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClustersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListClustersResponseBody) GetData() []*ListClustersResponseBodyData {
	return s.Data
}

func (s *ListClustersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClustersResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetCode(v string) *ListClustersResponseBody {
	s.Code = &v
	return s
}

func (s *ListClustersResponseBody) SetData(v []*ListClustersResponseBodyData) *ListClustersResponseBody {
	s.Data = v
	return s
}

func (s *ListClustersResponseBody) SetMessage(v string) *ListClustersResponseBody {
	s.Message = &v
	return s
}

func (s *ListClustersResponseBody) SetTotal(v int64) *ListClustersResponseBody {
	s.Total = &v
	return s
}

func (s *ListClustersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyData struct {
	// example:
	//
	// c666d4774f0e2440b979bf917bf100e40
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// Running
	ClusterStatus *string `json:"cluster_status,omitempty" xml:"cluster_status,omitempty"`
	// example:
	//
	// ACK
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// example:
	//
	// 2024-12-25T15:08:19
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// example:
	//
	// 5389fba5-92a1-4ff4-9b26-773b97828144
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// auto-name-sbvCT
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 2024-12-25T15:08:19
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s ListClustersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClustersResponseBodyData) GetClusterStatus() *string {
	return s.ClusterStatus
}

func (s *ListClustersResponseBodyData) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClustersResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListClustersResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListClustersResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListClustersResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *ListClustersResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListClustersResponseBodyData) SetClusterId(v string) *ListClustersResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyData) SetClusterStatus(v string) *ListClustersResponseBodyData {
	s.ClusterStatus = &v
	return s
}

func (s *ListClustersResponseBodyData) SetClusterType(v string) *ListClustersResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *ListClustersResponseBodyData) SetCreatedAt(v string) *ListClustersResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListClustersResponseBodyData) SetId(v string) *ListClustersResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListClustersResponseBodyData) SetName(v string) *ListClustersResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyData) SetRegion(v string) *ListClustersResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListClustersResponseBodyData) SetUpdatedAt(v string) *ListClustersResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListClustersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
