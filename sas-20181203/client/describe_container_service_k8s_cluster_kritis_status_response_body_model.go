// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerServiceK8sClusterKritisStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKritisStatus(v *DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus) *DescribeContainerServiceK8sClusterKritisStatusResponseBody
	GetKritisStatus() *DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus
	SetRequestId(v string) *DescribeContainerServiceK8sClusterKritisStatusResponseBody
	GetRequestId() *string
}

type DescribeContainerServiceK8sClusterKritisStatusResponseBody struct {
	// The Kritis status of the ACK cluster.
	KritisStatus *DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus `json:"KritisStatus,omitempty" xml:"KritisStatus,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578AB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerServiceK8sClusterKritisStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClusterKritisStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponseBody) GetKritisStatus() *DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus {
	return s.KritisStatus
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponseBody) SetKritisStatus(v *DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus) *DescribeContainerServiceK8sClusterKritisStatusResponseBody {
	s.KritisStatus = v
	return s
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponseBody) SetRequestId(v string) *DescribeContainerServiceK8sClusterKritisStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus struct {
	// Indicates whether Kritis is installed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Install *bool `json:"Install,omitempty" xml:"Install,omitempty"`
}

func (s DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus) GetInstall() *bool {
	return s.Install
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus) SetInstall(v bool) *DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus {
	s.Install = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponseBodyKritisStatus) Validate() error {
	return dara.Validate(s)
}
