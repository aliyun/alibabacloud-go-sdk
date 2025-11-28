// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeStatDeviceInfoResponseBody
	GetCode() *string
	SetData(v *DescribeStatDeviceInfoResponseBodyData) *DescribeStatDeviceInfoResponseBody
	GetData() *DescribeStatDeviceInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeStatDeviceInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeStatDeviceInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeStatDeviceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeStatDeviceInfoResponseBody
	GetSuccess() *bool
}

type DescribeStatDeviceInfoResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeStatDeviceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeStatDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStatDeviceInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeStatDeviceInfoResponseBody) GetData() *DescribeStatDeviceInfoResponseBodyData {
	return s.Data
}

func (s *DescribeStatDeviceInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeStatDeviceInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeStatDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStatDeviceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeStatDeviceInfoResponseBody) SetCode(v string) *DescribeStatDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStatDeviceInfoResponseBody) SetData(v *DescribeStatDeviceInfoResponseBodyData) *DescribeStatDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeStatDeviceInfoResponseBody) SetHttpStatusCode(v int32) *DescribeStatDeviceInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeStatDeviceInfoResponseBody) SetMessage(v string) *DescribeStatDeviceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStatDeviceInfoResponseBody) SetRequestId(v string) *DescribeStatDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStatDeviceInfoResponseBody) SetSuccess(v bool) *DescribeStatDeviceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeStatDeviceInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStatDeviceInfoResponseBodyData struct {
	BizChainList         []*DescribeStatDeviceInfoResponseBodyDataBizChainList `json:"BizChainList,omitempty" xml:"BizChainList,omitempty" type:"Repeated"`
	DistributableCount   *int32                                                `json:"DistributableCount,omitempty" xml:"DistributableCount,omitempty"`
	TotalAuthorizedCount *int32                                                `json:"TotalAuthorizedCount,omitempty" xml:"TotalAuthorizedCount,omitempty"`
}

func (s DescribeStatDeviceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeStatDeviceInfoResponseBodyData) GetBizChainList() []*DescribeStatDeviceInfoResponseBodyDataBizChainList {
	return s.BizChainList
}

func (s *DescribeStatDeviceInfoResponseBodyData) GetDistributableCount() *int32 {
	return s.DistributableCount
}

func (s *DescribeStatDeviceInfoResponseBodyData) GetTotalAuthorizedCount() *int32 {
	return s.TotalAuthorizedCount
}

func (s *DescribeStatDeviceInfoResponseBodyData) SetBizChainList(v []*DescribeStatDeviceInfoResponseBodyDataBizChainList) *DescribeStatDeviceInfoResponseBodyData {
	s.BizChainList = v
	return s
}

func (s *DescribeStatDeviceInfoResponseBodyData) SetDistributableCount(v int32) *DescribeStatDeviceInfoResponseBodyData {
	s.DistributableCount = &v
	return s
}

func (s *DescribeStatDeviceInfoResponseBodyData) SetTotalAuthorizedCount(v int32) *DescribeStatDeviceInfoResponseBodyData {
	s.TotalAuthorizedCount = &v
	return s
}

func (s *DescribeStatDeviceInfoResponseBodyData) Validate() error {
	if s.BizChainList != nil {
		for _, item := range s.BizChainList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStatDeviceInfoResponseBodyDataBizChainList struct {
	AuthorizedCount *int32  `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	BizChainName    *string `json:"BizChainName,omitempty" xml:"BizChainName,omitempty"`
}

func (s DescribeStatDeviceInfoResponseBodyDataBizChainList) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatDeviceInfoResponseBodyDataBizChainList) GoString() string {
	return s.String()
}

func (s *DescribeStatDeviceInfoResponseBodyDataBizChainList) GetAuthorizedCount() *int32 {
	return s.AuthorizedCount
}

func (s *DescribeStatDeviceInfoResponseBodyDataBizChainList) GetBizChainName() *string {
	return s.BizChainName
}

func (s *DescribeStatDeviceInfoResponseBodyDataBizChainList) SetAuthorizedCount(v int32) *DescribeStatDeviceInfoResponseBodyDataBizChainList {
	s.AuthorizedCount = &v
	return s
}

func (s *DescribeStatDeviceInfoResponseBodyDataBizChainList) SetBizChainName(v string) *DescribeStatDeviceInfoResponseBodyDataBizChainList {
	s.BizChainName = &v
	return s
}

func (s *DescribeStatDeviceInfoResponseBodyDataBizChainList) Validate() error {
	return dara.Validate(s)
}
