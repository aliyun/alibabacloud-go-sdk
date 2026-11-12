// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountAccessIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAccountAccessIdResponseBody
	GetCode() *int32
	SetData(v []*ListAccountAccessIdResponseBodyData) *ListAccountAccessIdResponseBody
	GetData() []*ListAccountAccessIdResponseBodyData
	SetMessage(v string) *ListAccountAccessIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAccountAccessIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAccountAccessIdResponseBody
	GetSuccess() *bool
}

type ListAccountAccessIdResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data []*ListAccountAccessIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message description.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAccountAccessIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccountAccessIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountAccessIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAccountAccessIdResponseBody) GetData() []*ListAccountAccessIdResponseBodyData {
	return s.Data
}

func (s *ListAccountAccessIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAccountAccessIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccountAccessIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAccountAccessIdResponseBody) SetCode(v int32) *ListAccountAccessIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListAccountAccessIdResponseBody) SetData(v []*ListAccountAccessIdResponseBodyData) *ListAccountAccessIdResponseBody {
	s.Data = v
	return s
}

func (s *ListAccountAccessIdResponseBody) SetMessage(v string) *ListAccountAccessIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListAccountAccessIdResponseBody) SetRequestId(v string) *ListAccountAccessIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountAccessIdResponseBody) SetSuccess(v bool) *ListAccountAccessIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListAccountAccessIdResponseBody) Validate() error {
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

type ListAccountAccessIdResponseBodyData struct {
	// The bound AccessKeyId.
	//
	// example:
	//
	// ABCXXXXXXXX
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The MD5 value of the multi-cloud AccessKeyId.
	//
	// example:
	//
	// abcXXXXXXXX
	AccessIdMd5 *string `json:"AccessIdMd5,omitempty" xml:"AccessIdMd5,omitempty"`
	// The cloud account ID.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The account information of the multi-cloud AccessKeyId. The format is Alibaba Cloud account ID|Alibaba Cloud account name|multi-cloud AccessKeyId.
	//
	// example:
	//
	// 123xxxxxx|xxxx|ABCXXXXX
	AccountStr *string `json:"AccountStr,omitempty" xml:"AccountStr,omitempty"`
	// Indicates whether the AccessKeyId is bound to threat detection and response. Valid values:
	//
	//  - 0: Not bound.
	//
	//  - 1: Bound.
	//
	// example:
	//
	// 1
	Bound *int32 `json:"Bound,omitempty" xml:"Bound,omitempty"`
	// The code of the multi-cloud provider.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The Alibaba Cloud account ID that corresponds to the multi-cloud AccessKeyId.
	//
	// example:
	//
	// ABCXXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListAccountAccessIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAccountAccessIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAccountAccessIdResponseBodyData) GetAccessId() *string {
	return s.AccessId
}

func (s *ListAccountAccessIdResponseBodyData) GetAccessIdMd5() *string {
	return s.AccessIdMd5
}

func (s *ListAccountAccessIdResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *ListAccountAccessIdResponseBodyData) GetAccountStr() *string {
	return s.AccountStr
}

func (s *ListAccountAccessIdResponseBodyData) GetBound() *int32 {
	return s.Bound
}

func (s *ListAccountAccessIdResponseBodyData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListAccountAccessIdResponseBodyData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *ListAccountAccessIdResponseBodyData) SetAccessId(v string) *ListAccountAccessIdResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetAccessIdMd5(v string) *ListAccountAccessIdResponseBodyData {
	s.AccessIdMd5 = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetAccountId(v string) *ListAccountAccessIdResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetAccountStr(v string) *ListAccountAccessIdResponseBodyData {
	s.AccountStr = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetBound(v int32) *ListAccountAccessIdResponseBodyData {
	s.Bound = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetCloudCode(v string) *ListAccountAccessIdResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetSubUserId(v int64) *ListAccountAccessIdResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
