// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDeviceInfoResponseBody
	GetCode() *string
	SetData(v *DescribeDeviceInfoResponseBodyData) *DescribeDeviceInfoResponseBody
	GetData() *DescribeDeviceInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeDeviceInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeDeviceInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDeviceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDeviceInfoResponseBody
	GetSuccess() *bool
}

type DescribeDeviceInfoResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeDeviceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDeviceInfoResponseBody) GetData() *DescribeDeviceInfoResponseBodyData {
	return s.Data
}

func (s *DescribeDeviceInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDeviceInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeviceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDeviceInfoResponseBody) SetCode(v string) *DescribeDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetData(v *DescribeDeviceInfoResponseBodyData) *DescribeDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetHttpStatusCode(v int32) *DescribeDeviceInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetMessage(v string) *DescribeDeviceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetRequestId(v string) *DescribeDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetSuccess(v bool) *DescribeDeviceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDeviceInfoResponseBodyData struct {
	AuthorizedCount    *int32 `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	DeviceCount        *int32 `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	DistributableCount *int32 `json:"DistributableCount,omitempty" xml:"DistributableCount,omitempty"`
}

func (s DescribeDeviceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBodyData) GetAuthorizedCount() *int32 {
	return s.AuthorizedCount
}

func (s *DescribeDeviceInfoResponseBodyData) GetDeviceCount() *int32 {
	return s.DeviceCount
}

func (s *DescribeDeviceInfoResponseBodyData) GetDistributableCount() *int32 {
	return s.DistributableCount
}

func (s *DescribeDeviceInfoResponseBodyData) SetAuthorizedCount(v int32) *DescribeDeviceInfoResponseBodyData {
	s.AuthorizedCount = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyData) SetDeviceCount(v int32) *DescribeDeviceInfoResponseBodyData {
	s.DeviceCount = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyData) SetDistributableCount(v int32) *DescribeDeviceInfoResponseBodyData {
	s.DistributableCount = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
