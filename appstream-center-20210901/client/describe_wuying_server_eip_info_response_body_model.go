// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWuyingServerEipInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEipInfoModel(v *DescribeWuyingServerEipInfoResponseBodyEipInfoModel) *DescribeWuyingServerEipInfoResponseBody
	GetEipInfoModel() *DescribeWuyingServerEipInfoResponseBodyEipInfoModel
	SetRequestId(v string) *DescribeWuyingServerEipInfoResponseBody
	GetRequestId() *string
}

type DescribeWuyingServerEipInfoResponseBody struct {
	EipInfoModel *DescribeWuyingServerEipInfoResponseBodyEipInfoModel `json:"EipInfoModel,omitempty" xml:"EipInfoModel,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWuyingServerEipInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWuyingServerEipInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWuyingServerEipInfoResponseBody) GetEipInfoModel() *DescribeWuyingServerEipInfoResponseBodyEipInfoModel {
	return s.EipInfoModel
}

func (s *DescribeWuyingServerEipInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWuyingServerEipInfoResponseBody) SetEipInfoModel(v *DescribeWuyingServerEipInfoResponseBodyEipInfoModel) *DescribeWuyingServerEipInfoResponseBody {
	s.EipInfoModel = v
	return s
}

func (s *DescribeWuyingServerEipInfoResponseBody) SetRequestId(v string) *DescribeWuyingServerEipInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWuyingServerEipInfoResponseBody) Validate() error {
	if s.EipInfoModel != nil {
		if err := s.EipInfoModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeWuyingServerEipInfoResponseBodyEipInfoModel struct {
	// example:
	//
	// 171.xxx.xxx.221
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// example:
	//
	// eni-bp174p2xxxxxbyh02ix
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// example:
	//
	// 6606/6607
	ServerPortRange *string `json:"ServerPortRange,omitempty" xml:"ServerPortRange,omitempty"`
}

func (s DescribeWuyingServerEipInfoResponseBodyEipInfoModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeWuyingServerEipInfoResponseBodyEipInfoModel) GoString() string {
	return s.String()
}

func (s *DescribeWuyingServerEipInfoResponseBodyEipInfoModel) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeWuyingServerEipInfoResponseBodyEipInfoModel) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeWuyingServerEipInfoResponseBodyEipInfoModel) GetServerPortRange() *string {
	return s.ServerPortRange
}

func (s *DescribeWuyingServerEipInfoResponseBodyEipInfoModel) SetIpAddress(v string) *DescribeWuyingServerEipInfoResponseBodyEipInfoModel {
	s.IpAddress = &v
	return s
}

func (s *DescribeWuyingServerEipInfoResponseBodyEipInfoModel) SetNetworkInterfaceId(v string) *DescribeWuyingServerEipInfoResponseBodyEipInfoModel {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeWuyingServerEipInfoResponseBodyEipInfoModel) SetServerPortRange(v string) *DescribeWuyingServerEipInfoResponseBodyEipInfoModel {
	s.ServerPortRange = &v
	return s
}

func (s *DescribeWuyingServerEipInfoResponseBodyEipInfoModel) Validate() error {
	return dara.Validate(s)
}
