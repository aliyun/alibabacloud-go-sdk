// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerScanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeContainerScanConfigResponseBodyData) *DescribeContainerScanConfigResponseBody
	GetData() *DescribeContainerScanConfigResponseBodyData
	SetHttpStatusCode(v int32) *DescribeContainerScanConfigResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeContainerScanConfigResponseBody
	GetRequestId() *string
}

type DescribeContainerScanConfigResponseBody struct {
	// The response parameters.
	Data *DescribeContainerScanConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 69BFFCDE-37D6-5A49-A8BC-BB03AC83****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerScanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerScanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerScanConfigResponseBody) GetData() *DescribeContainerScanConfigResponseBodyData {
	return s.Data
}

func (s *DescribeContainerScanConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeContainerScanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerScanConfigResponseBody) SetData(v *DescribeContainerScanConfigResponseBodyData) *DescribeContainerScanConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContainerScanConfigResponseBody) SetHttpStatusCode(v int32) *DescribeContainerScanConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeContainerScanConfigResponseBody) SetRequestId(v string) *DescribeContainerScanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerScanConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeContainerScanConfigResponseBodyData struct {
	// The total number of container applications in the cluster.
	//
	// example:
	//
	// 100
	AllCount *int32 `json:"AllCount,omitempty" xml:"AllCount,omitempty"`
	// The names of the container applications.
	//
	// example:
	//
	// [\\"alicloud-monitor-controller\\"]
	AppNames *string `json:"AppNames,omitempty" xml:"AppNames,omitempty"`
	// The number of selected container applications.
	//
	// example:
	//
	// 10
	ChooseCount *int32 `json:"ChooseCount,omitempty" xml:"ChooseCount,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c8ca91e0907d94efaba7fb0827eb9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// hhht-cluster-02
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s DescribeContainerScanConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerScanConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContainerScanConfigResponseBodyData) GetAllCount() *int32 {
	return s.AllCount
}

func (s *DescribeContainerScanConfigResponseBodyData) GetAppNames() *string {
	return s.AppNames
}

func (s *DescribeContainerScanConfigResponseBodyData) GetChooseCount() *int32 {
	return s.ChooseCount
}

func (s *DescribeContainerScanConfigResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerScanConfigResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeContainerScanConfigResponseBodyData) SetAllCount(v int32) *DescribeContainerScanConfigResponseBodyData {
	s.AllCount = &v
	return s
}

func (s *DescribeContainerScanConfigResponseBodyData) SetAppNames(v string) *DescribeContainerScanConfigResponseBodyData {
	s.AppNames = &v
	return s
}

func (s *DescribeContainerScanConfigResponseBodyData) SetChooseCount(v int32) *DescribeContainerScanConfigResponseBodyData {
	s.ChooseCount = &v
	return s
}

func (s *DescribeContainerScanConfigResponseBodyData) SetClusterId(v string) *DescribeContainerScanConfigResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerScanConfigResponseBodyData) SetClusterName(v string) *DescribeContainerScanConfigResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *DescribeContainerScanConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
