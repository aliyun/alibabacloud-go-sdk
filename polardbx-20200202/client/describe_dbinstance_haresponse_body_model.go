// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceHAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceHAResponseBodyData) *DescribeDBInstanceHAResponseBody
	GetData() *DescribeDBInstanceHAResponseBodyData
	SetMessage(v string) *DescribeDBInstanceHAResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDBInstanceHAResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDBInstanceHAResponseBody
	GetSuccess() *bool
}

type DescribeDBInstanceHAResponseBody struct {
	// The high availability deployment information of the instance.
	Data *DescribeDBInstanceHAResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// successs
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A6D328C-84B8-40DC-BF49-6C73984D7494
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDBInstanceHAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceHAResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHAResponseBody) GetData() *DescribeDBInstanceHAResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceHAResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDBInstanceHAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceHAResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDBInstanceHAResponseBody) SetData(v *DescribeDBInstanceHAResponseBodyData) *DescribeDBInstanceHAResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceHAResponseBody) SetMessage(v string) *DescribeDBInstanceHAResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBody) SetRequestId(v string) *DescribeDBInstanceHAResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBody) SetSuccess(v bool) *DescribeDBInstanceHAResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceHAResponseBodyData struct {
	// The zone ID of the primary node.
	//
	// example:
	//
	// cn-hangzhou-d
	PrimaryAzoneId *string `json:"PrimaryAzoneId,omitempty" xml:"PrimaryAzoneId,omitempty"`
	// The region ID of the primary node.
	//
	// example:
	//
	// cn-hangzhou
	PrimaryRegionId *string `json:"PrimaryRegionId,omitempty" xml:"PrimaryRegionId,omitempty"`
	// The zone ID of the secondary node.
	//
	// example:
	//
	// cn-hangzhou-b
	SecondaryAzoneId *string `json:"SecondaryAzoneId,omitempty" xml:"SecondaryAzoneId,omitempty"`
	// The region ID of the secondary node.
	//
	// example:
	//
	// cn-hangzhou
	SecondaryRegionId *string `json:"SecondaryRegionId,omitempty" xml:"SecondaryRegionId,omitempty"`
	// The high availability topology type of the instance.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
}

func (s DescribeDBInstanceHAResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceHAResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHAResponseBodyData) GetPrimaryAzoneId() *string {
	return s.PrimaryAzoneId
}

func (s *DescribeDBInstanceHAResponseBodyData) GetPrimaryRegionId() *string {
	return s.PrimaryRegionId
}

func (s *DescribeDBInstanceHAResponseBodyData) GetSecondaryAzoneId() *string {
	return s.SecondaryAzoneId
}

func (s *DescribeDBInstanceHAResponseBodyData) GetSecondaryRegionId() *string {
	return s.SecondaryRegionId
}

func (s *DescribeDBInstanceHAResponseBodyData) GetTopologyType() *string {
	return s.TopologyType
}

func (s *DescribeDBInstanceHAResponseBodyData) SetPrimaryAzoneId(v string) *DescribeDBInstanceHAResponseBodyData {
	s.PrimaryAzoneId = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBodyData) SetPrimaryRegionId(v string) *DescribeDBInstanceHAResponseBodyData {
	s.PrimaryRegionId = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBodyData) SetSecondaryAzoneId(v string) *DescribeDBInstanceHAResponseBodyData {
	s.SecondaryAzoneId = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBodyData) SetSecondaryRegionId(v string) *DescribeDBInstanceHAResponseBodyData {
	s.SecondaryRegionId = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBodyData) SetTopologyType(v string) *DescribeDBInstanceHAResponseBodyData {
	s.TopologyType = &v
	return s
}

func (s *DescribeDBInstanceHAResponseBodyData) Validate() error {
	return dara.Validate(s)
}
