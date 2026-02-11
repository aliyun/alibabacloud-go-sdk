// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWafScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeWafScopeResponseBody
	GetCode() *int32
	SetData(v []*DescribeWafScopeResponseBodyData) *DescribeWafScopeResponseBody
	GetData() []*DescribeWafScopeResponseBodyData
	SetMessage(v string) *DescribeWafScopeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeWafScopeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeWafScopeResponseBody
	GetSuccess() *bool
}

type DescribeWafScopeResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeWafScopeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeWafScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafScopeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeWafScopeResponseBody) GetData() []*DescribeWafScopeResponseBodyData {
	return s.Data
}

func (s *DescribeWafScopeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeWafScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWafScopeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeWafScopeResponseBody) SetCode(v int32) *DescribeWafScopeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWafScopeResponseBody) SetData(v []*DescribeWafScopeResponseBodyData) *DescribeWafScopeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWafScopeResponseBody) SetMessage(v string) *DescribeWafScopeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWafScopeResponseBody) SetRequestId(v string) *DescribeWafScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWafScopeResponseBody) SetSuccess(v bool) *DescribeWafScopeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeWafScopeResponseBody) Validate() error {
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

type DescribeWafScopeResponseBodyData struct {
	// The ID of the Alibaba Cloud account in SIEM.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The domain names that are protected by the WAF instance.
	//
	// example:
	//
	// [123.com, 456.com]
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the WAF instance.
	//
	// example:
	//
	// waf-cn-tl123ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeWafScopeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafScopeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeResponseBodyData) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *DescribeWafScopeResponseBodyData) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeWafScopeResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeWafScopeResponseBodyData) SetAliuid(v int64) *DescribeWafScopeResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *DescribeWafScopeResponseBodyData) SetDomains(v []*string) *DescribeWafScopeResponseBodyData {
	s.Domains = v
	return s
}

func (s *DescribeWafScopeResponseBodyData) SetInstanceId(v string) *DescribeWafScopeResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeWafScopeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
