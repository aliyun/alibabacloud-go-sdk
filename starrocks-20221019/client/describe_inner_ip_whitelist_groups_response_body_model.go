// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInnerIpWhitelistGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeInnerIpWhitelistGroupsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeInnerIpWhitelistGroupsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeInnerIpWhitelistGroupsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeInnerIpWhitelistGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInnerIpWhitelistGroupsResponseBody
	GetSuccess() *bool
	SetData(v []*DescribeInnerIpWhitelistGroupsResponseBodyData) *DescribeInnerIpWhitelistGroupsResponseBody
	GetData() []*DescribeInnerIpWhitelistGroupsResponseBodyData
}

type DescribeInnerIpWhitelistGroupsResponseBody struct {
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Data    []*DescribeInnerIpWhitelistGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeInnerIpWhitelistGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInnerIpWhitelistGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) GetData() []*DescribeInnerIpWhitelistGroupsResponseBodyData {
	return s.Data
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) SetErrCode(v string) *DescribeInnerIpWhitelistGroupsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) SetErrMessage(v string) *DescribeInnerIpWhitelistGroupsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) SetHttpStatusCode(v int32) *DescribeInnerIpWhitelistGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) SetRequestId(v string) *DescribeInnerIpWhitelistGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) SetSuccess(v bool) *DescribeInnerIpWhitelistGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) SetData(v []*DescribeInnerIpWhitelistGroupsResponseBodyData) *DescribeInnerIpWhitelistGroupsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInnerIpWhitelistGroupsResponseBodyData struct {
	CidrIpList []*string `json:"CidrIpList,omitempty" xml:"CidrIpList,omitempty" type:"Repeated"`
	// example:
	//
	// test1
	InnerIpWhitelistGroupId *string `json:"InnerIpWhitelistGroupId,omitempty" xml:"InnerIpWhitelistGroupId,omitempty"`
}

func (s DescribeInnerIpWhitelistGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInnerIpWhitelistGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInnerIpWhitelistGroupsResponseBodyData) GetCidrIpList() []*string {
	return s.CidrIpList
}

func (s *DescribeInnerIpWhitelistGroupsResponseBodyData) GetInnerIpWhitelistGroupId() *string {
	return s.InnerIpWhitelistGroupId
}

func (s *DescribeInnerIpWhitelistGroupsResponseBodyData) SetCidrIpList(v []*string) *DescribeInnerIpWhitelistGroupsResponseBodyData {
	s.CidrIpList = v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsResponseBodyData) SetInnerIpWhitelistGroupId(v string) *DescribeInnerIpWhitelistGroupsResponseBodyData {
	s.InnerIpWhitelistGroupId = &v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
