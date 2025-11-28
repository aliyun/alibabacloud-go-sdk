// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeStatInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEdgeStatInfoResponseBody
	GetCode() *string
	SetData(v *DescribeEdgeStatInfoResponseBodyData) *DescribeEdgeStatInfoResponseBody
	GetData() *DescribeEdgeStatInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeEdgeStatInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeEdgeStatInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEdgeStatInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEdgeStatInfoResponseBody
	GetSuccess() *bool
}

type DescribeEdgeStatInfoResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeEdgeStatInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEdgeStatInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeStatInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdgeStatInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEdgeStatInfoResponseBody) GetData() *DescribeEdgeStatInfoResponseBodyData {
	return s.Data
}

func (s *DescribeEdgeStatInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeEdgeStatInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEdgeStatInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEdgeStatInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEdgeStatInfoResponseBody) SetCode(v string) *DescribeEdgeStatInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEdgeStatInfoResponseBody) SetData(v *DescribeEdgeStatInfoResponseBodyData) *DescribeEdgeStatInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEdgeStatInfoResponseBody) SetHttpStatusCode(v int32) *DescribeEdgeStatInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeEdgeStatInfoResponseBody) SetMessage(v string) *DescribeEdgeStatInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEdgeStatInfoResponseBody) SetRequestId(v string) *DescribeEdgeStatInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEdgeStatInfoResponseBody) SetSuccess(v bool) *DescribeEdgeStatInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEdgeStatInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEdgeStatInfoResponseBodyData struct {
	TotalDeviceLicenseCount *int64 `json:"TotalDeviceLicenseCount,omitempty" xml:"TotalDeviceLicenseCount,omitempty"`
	UsedDeviceLicenseCount  *int64 `json:"UsedDeviceLicenseCount,omitempty" xml:"UsedDeviceLicenseCount,omitempty"`
}

func (s DescribeEdgeStatInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeStatInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEdgeStatInfoResponseBodyData) GetTotalDeviceLicenseCount() *int64 {
	return s.TotalDeviceLicenseCount
}

func (s *DescribeEdgeStatInfoResponseBodyData) GetUsedDeviceLicenseCount() *int64 {
	return s.UsedDeviceLicenseCount
}

func (s *DescribeEdgeStatInfoResponseBodyData) SetTotalDeviceLicenseCount(v int64) *DescribeEdgeStatInfoResponseBodyData {
	s.TotalDeviceLicenseCount = &v
	return s
}

func (s *DescribeEdgeStatInfoResponseBodyData) SetUsedDeviceLicenseCount(v int64) *DescribeEdgeStatInfoResponseBodyData {
	s.UsedDeviceLicenseCount = &v
	return s
}

func (s *DescribeEdgeStatInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
