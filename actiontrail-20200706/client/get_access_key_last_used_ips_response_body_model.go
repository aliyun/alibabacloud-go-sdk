// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIps(v []*GetAccessKeyLastUsedIpsResponseBodyIps) *GetAccessKeyLastUsedIpsResponseBody
	GetIps() []*GetAccessKeyLastUsedIpsResponseBodyIps
	SetNextToken(v string) *GetAccessKeyLastUsedIpsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetAccessKeyLastUsedIpsResponseBody
	GetRequestId() *string
}

type GetAccessKeyLastUsedIpsResponseBody struct {
	// The IP addresses.
	//
	// This parameter is required.
	Ips []*GetAccessKeyLastUsedIpsResponseBodyIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJhY2NvdW50IjoiMTQyNDM3OTU4NjM4NzE2MSIsImV2ZW50SWQiOiI3MkJDRTExRi02OTU3LTQ0NUItQjY0MC1CNEUyMkM4NUEwQzgiLCJsb2dJZCI6IjgyLTE0MjQzNzk1ODYzODcxNjEiLCJ0aW1lIjoxNjAyMzExNTQwMD****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 145318BE-DEE1-4C57-AA7C-5BE7D34A6AE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyLastUsedIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedIpsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedIpsResponseBody) GetIps() []*GetAccessKeyLastUsedIpsResponseBodyIps {
	return s.Ips
}

func (s *GetAccessKeyLastUsedIpsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetAccessKeyLastUsedIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessKeyLastUsedIpsResponseBody) SetIps(v []*GetAccessKeyLastUsedIpsResponseBodyIps) *GetAccessKeyLastUsedIpsResponseBody {
	s.Ips = v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBody) SetNextToken(v string) *GetAccessKeyLastUsedIpsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAccessKeyLastUsedIpsResponseBodyIps struct {
	// The event details.
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
	// The IP address.
	//
	// example:
	//
	// ``10.10.**.**``
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The event source.
	//
	// Valid values:
	//
	// 	- Internal: other events.
	//
	// 	- ManagementEvent: management events.
	//
	// 	- DataEvent: data events.
	//
	// example:
	//
	// ManagementEvent
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The timestamp when the IP address was used. Unit: milliseconds.
	//
	// example:
	//
	// 1657247532000
	UsedTimestamp *int64 `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
}

func (s GetAccessKeyLastUsedIpsResponseBodyIps) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedIpsResponseBodyIps) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) GetDetail() *string {
	return s.Detail
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) GetIp() *string {
	return s.Ip
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) GetSource() *string {
	return s.Source
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) GetUsedTimestamp() *int64 {
	return s.UsedTimestamp
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) SetDetail(v string) *GetAccessKeyLastUsedIpsResponseBodyIps {
	s.Detail = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) SetIp(v string) *GetAccessKeyLastUsedIpsResponseBodyIps {
	s.Ip = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) SetSource(v string) *GetAccessKeyLastUsedIpsResponseBodyIps {
	s.Source = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) SetUsedTimestamp(v int64) *GetAccessKeyLastUsedIpsResponseBodyIps {
	s.UsedTimestamp = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) Validate() error {
	return dara.Validate(s)
}
