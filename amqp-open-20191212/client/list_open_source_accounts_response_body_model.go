// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenSourceAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListOpenSourceAccountsResponseBody
	GetCode() *int32
	SetData(v []*ListOpenSourceAccountsResponseBodyData) *ListOpenSourceAccountsResponseBody
	GetData() []*ListOpenSourceAccountsResponseBodyData
	SetMaxResults(v int32) *ListOpenSourceAccountsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListOpenSourceAccountsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListOpenSourceAccountsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListOpenSourceAccountsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOpenSourceAccountsResponseBody
	GetSuccess() *bool
}

type ListOpenSourceAccountsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListOpenSourceAccountsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token that marks the end of the current query. This token is passed as a parameter in the next call to continue pagination. Set this parameter to an empty string for the first call or when the last page is returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 021788F6-E50C-4BD6-9F80-66B0A19A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOpenSourceAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOpenSourceAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenSourceAccountsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListOpenSourceAccountsResponseBody) GetData() []*ListOpenSourceAccountsResponseBodyData {
	return s.Data
}

func (s *ListOpenSourceAccountsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOpenSourceAccountsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOpenSourceAccountsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOpenSourceAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOpenSourceAccountsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOpenSourceAccountsResponseBody) SetCode(v int32) *ListOpenSourceAccountsResponseBody {
	s.Code = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBody) SetData(v []*ListOpenSourceAccountsResponseBodyData) *ListOpenSourceAccountsResponseBody {
	s.Data = v
	return s
}

func (s *ListOpenSourceAccountsResponseBody) SetMaxResults(v int32) *ListOpenSourceAccountsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBody) SetMessage(v string) *ListOpenSourceAccountsResponseBody {
	s.Message = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBody) SetNextToken(v string) *ListOpenSourceAccountsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBody) SetRequestId(v string) *ListOpenSourceAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBody) SetSuccess(v bool) *ListOpenSourceAccountsResponseBody {
	s.Success = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBody) Validate() error {
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

type ListOpenSourceAccountsResponseBodyData struct {
	// The Alibaba Cloud UID.
	//
	// example:
	//
	// 12345678900123
	AliyunUserId *int64 `json:"AliyunUserId,omitempty" xml:"AliyunUserId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// amqp-abc***
	CInstanceId *string `json:"CInstanceId,omitempty" xml:"CInstanceId,omitempty"`
	// The time when the user was created.
	//
	// example:
	//
	// 1704067200000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// {}
	ExtraJson *string `json:"ExtraJson,omitempty" xml:"ExtraJson,omitempty"`
	// The hash algorithm.
	//
	// example:
	//
	// bcrypt
	HashingAlgorithm *string `json:"HashingAlgorithm,omitempty" xml:"HashingAlgorithm,omitempty"`
	// The quota configuration.
	//
	// example:
	//
	// null
	Limits *string `json:"Limits,omitempty" xml:"Limits,omitempty"`
	// The username.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The password hash.
	//
	// example:
	//
	// f6af6bbb91e947d988d191bfe01c9a9b
	PasswordHash *string `json:"PasswordHash,omitempty" xml:"PasswordHash,omitempty"`
	// The tags.
	//
	// example:
	//
	// test
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Indicates whether the password is a weak password. Valid values:
	//
	// - 0: No.
	//
	// - 1: Yes.
	WeakPassword *bool `json:"WeakPassword,omitempty" xml:"WeakPassword,omitempty"`
}

func (s ListOpenSourceAccountsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOpenSourceAccountsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOpenSourceAccountsResponseBodyData) GetAliyunUserId() *int64 {
	return s.AliyunUserId
}

func (s *ListOpenSourceAccountsResponseBodyData) GetCInstanceId() *string {
	return s.CInstanceId
}

func (s *ListOpenSourceAccountsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListOpenSourceAccountsResponseBodyData) GetExtraJson() *string {
	return s.ExtraJson
}

func (s *ListOpenSourceAccountsResponseBodyData) GetHashingAlgorithm() *string {
	return s.HashingAlgorithm
}

func (s *ListOpenSourceAccountsResponseBodyData) GetLimits() *string {
	return s.Limits
}

func (s *ListOpenSourceAccountsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListOpenSourceAccountsResponseBodyData) GetPasswordHash() *string {
	return s.PasswordHash
}

func (s *ListOpenSourceAccountsResponseBodyData) GetTags() *string {
	return s.Tags
}

func (s *ListOpenSourceAccountsResponseBodyData) GetWeakPassword() *bool {
	return s.WeakPassword
}

func (s *ListOpenSourceAccountsResponseBodyData) SetAliyunUserId(v int64) *ListOpenSourceAccountsResponseBodyData {
	s.AliyunUserId = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBodyData) SetCInstanceId(v string) *ListOpenSourceAccountsResponseBodyData {
	s.CInstanceId = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBodyData) SetCreateTime(v string) *ListOpenSourceAccountsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBodyData) SetExtraJson(v string) *ListOpenSourceAccountsResponseBodyData {
	s.ExtraJson = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBodyData) SetHashingAlgorithm(v string) *ListOpenSourceAccountsResponseBodyData {
	s.HashingAlgorithm = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBodyData) SetLimits(v string) *ListOpenSourceAccountsResponseBodyData {
	s.Limits = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBodyData) SetName(v string) *ListOpenSourceAccountsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBodyData) SetPasswordHash(v string) *ListOpenSourceAccountsResponseBodyData {
	s.PasswordHash = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBodyData) SetTags(v string) *ListOpenSourceAccountsResponseBodyData {
	s.Tags = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBodyData) SetWeakPassword(v bool) *ListOpenSourceAccountsResponseBodyData {
	s.WeakPassword = &v
	return s
}

func (s *ListOpenSourceAccountsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
