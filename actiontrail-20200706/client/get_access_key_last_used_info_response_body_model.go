// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyId(v string) *GetAccessKeyLastUsedInfoResponseBody
	GetAccessKeyId() *string
	SetAccountId(v string) *GetAccessKeyLastUsedInfoResponseBody
	GetAccountId() *string
	SetAccountType(v string) *GetAccessKeyLastUsedInfoResponseBody
	GetAccountType() *string
	SetDetail(v string) *GetAccessKeyLastUsedInfoResponseBody
	GetDetail() *string
	SetOwnerId(v string) *GetAccessKeyLastUsedInfoResponseBody
	GetOwnerId() *string
	SetRequestId(v string) *GetAccessKeyLastUsedInfoResponseBody
	GetRequestId() *string
	SetServiceName(v string) *GetAccessKeyLastUsedInfoResponseBody
	GetServiceName() *string
	SetServiceNameCn(v string) *GetAccessKeyLastUsedInfoResponseBody
	GetServiceNameCn() *string
	SetServiceNameEn(v string) *GetAccessKeyLastUsedInfoResponseBody
	GetServiceNameEn() *string
	SetSource(v string) *GetAccessKeyLastUsedInfoResponseBody
	GetSource() *string
	SetUsedTimestamp(v int64) *GetAccessKeyLastUsedInfoResponseBody
	GetUsedTimestamp() *int64
	SetUserName(v string) *GetAccessKeyLastUsedInfoResponseBody
	GetUserName() *string
}

type GetAccessKeyLastUsedInfoResponseBody struct {
	// The AccessKey ID.
	//
	// example:
	//
	// LTAI****************
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 104758519118****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The type of the account to which the AccessKey pair belongs.
	//
	// example:
	//
	// ram-user
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The details about the event.
	//
	// example:
	//
	// {
	//
	//   "eventId": "239EB588-CD24-522E-B0B5-174A1A58****",
	//
	//   "eventVersion": 1,
	//
	//   "eventSource": "ecs.cn-hangzhou.aliyuncs.com",
	//
	//   "sourceIpAddress": "``10.10.**.**``",
	//
	//   "eventType": "ApiCall",
	//
	//   "userIdentity": {
	//
	//     "accountId": "104758519118****",
	//
	//     "principalId": "24549429003625****",
	//
	//     "type": "ram-user",
	//
	//     "userName": "alice"
	//
	//   },
	//
	//   "serviceName": "Ecs",
	//
	//   "apiVersion": "2016-01-20",
	//
	//   "requestId": "239EB588-CD24-522E-B0B5-174A1A588BE0",
	//
	//   "eventTime": "2021-08-05T09:21:32Z",
	//
	//   "isGlobal": false,
	//
	//   "acsRegion": "cn-hangzhou",
	//
	//   "eventName": "DescribeInstances"
	//
	// }
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the account to which the AccessKey pair belongs.
	//
	// example:
	//
	// 24549429003625****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 239EB588-CD24-522E-B0B5-174A1A588BE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud service that was last accessed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The Chinese name of the Alibaba Cloud service that was last accessed.
	//
	// example:
	//
	// Elastic Compute Service (ECS)
	ServiceNameCn *string `json:"ServiceNameCn,omitempty" xml:"ServiceNameCn,omitempty"`
	// The English name of the Alibaba Cloud service that was last accessed.
	//
	// example:
	//
	// Elastic Compute Service
	ServiceNameEn *string `json:"ServiceNameEn,omitempty" xml:"ServiceNameEn,omitempty"`
	// The event source.
	//
	// example:
	//
	// ManagementEvent
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The timestamp when the AccessKey pair was last called.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1657247532000
	UsedTimestamp *int64 `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
	// The name of the account to which the AccessKey pair belongs.
	//
	// If the value of the AccountType parameter is root-account, the value of the UserName parameter is root. If the value of the AccountType parameter is ram-user, the value of the UserName parameter is the name of a RAM user.
	//
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetAccessKeyLastUsedInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetAccountType() *string {
	return s.AccountType
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetDetail() *string {
	return s.Detail
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetServiceNameCn() *string {
	return s.ServiceNameCn
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetServiceNameEn() *string {
	return s.ServiceNameEn
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetSource() *string {
	return s.Source
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetUsedTimestamp() *int64 {
	return s.UsedTimestamp
}

func (s *GetAccessKeyLastUsedInfoResponseBody) GetUserName() *string {
	return s.UserName
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetAccessKeyId(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetAccountId(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetAccountType(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.AccountType = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetDetail(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.Detail = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetOwnerId(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetServiceName(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetServiceNameCn(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.ServiceNameCn = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetServiceNameEn(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.ServiceNameEn = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetSource(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.Source = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetUsedTimestamp(v int64) *GetAccessKeyLastUsedInfoResponseBody {
	s.UsedTimestamp = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetUserName(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.UserName = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
