// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScopeUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeScopeUsersResponseBody
	GetCode() *int32
	SetData(v []*DescribeScopeUsersResponseBodyData) *DescribeScopeUsersResponseBody
	GetData() []*DescribeScopeUsersResponseBodyData
	SetMessage(v string) *DescribeScopeUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeScopeUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeScopeUsersResponseBody
	GetSuccess() *bool
}

type DescribeScopeUsersResponseBody struct {
	// The HTTP status code that is returned.
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
	Data []*DescribeScopeUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeScopeUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScopeUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeScopeUsersResponseBody) GetData() []*DescribeScopeUsersResponseBodyData {
	return s.Data
}

func (s *DescribeScopeUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeScopeUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScopeUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeScopeUsersResponseBody) SetCode(v int32) *DescribeScopeUsersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetData(v []*DescribeScopeUsersResponseBodyData) *DescribeScopeUsersResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetMessage(v string) *DescribeScopeUsersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetRequestId(v string) *DescribeScopeUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetSuccess(v bool) *DescribeScopeUsersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeScopeUsersResponseBody) Validate() error {
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

type DescribeScopeUsersResponseBodyData struct {
	// The ID of the security information and event management (SIEM) user.
	//
	// example:
	//
	// 123456789****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 云code。  取值：
	//
	// - qcloud：腾讯云
	//
	// - hcloud：华为云
	//
	// example:
	//
	// qcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// An array consisting of the domain names that are protected by the WAF instance.
	//
	// example:
	//
	// [123.com, 456.com]
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// example:
	//
	// waf-cn-tl123ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 多云用户ID。
	//
	// example:
	//
	// 123456789****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// test001
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeScopeUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeScopeUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeScopeUsersResponseBodyData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DescribeScopeUsersResponseBodyData) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeScopeUsersResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeScopeUsersResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *DescribeScopeUsersResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *DescribeScopeUsersResponseBodyData) SetAliUid(v int64) *DescribeScopeUsersResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetCloudCode(v string) *DescribeScopeUsersResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetDomains(v []*string) *DescribeScopeUsersResponseBodyData {
	s.Domains = v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetInstanceId(v string) *DescribeScopeUsersResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetUserId(v string) *DescribeScopeUsersResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetUserName(v string) *DescribeScopeUsersResponseBodyData {
	s.UserName = &v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
