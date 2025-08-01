// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamingTrackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListNamingTrackRequest
	GetAcceptLanguage() *string
	SetEndTs(v int64) *ListNamingTrackRequest
	GetEndTs() *int64
	SetGroup(v string) *ListNamingTrackRequest
	GetGroup() *string
	SetInstanceId(v string) *ListNamingTrackRequest
	GetInstanceId() *string
	SetIp(v string) *ListNamingTrackRequest
	GetIp() *string
	SetNamespaceId(v string) *ListNamingTrackRequest
	GetNamespaceId() *string
	SetPageNum(v int64) *ListNamingTrackRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListNamingTrackRequest
	GetPageSize() *int64
	SetRequestPars(v string) *ListNamingTrackRequest
	GetRequestPars() *string
	SetReverse(v bool) *ListNamingTrackRequest
	GetReverse() *bool
	SetServiceName(v string) *ListNamingTrackRequest
	GetServiceName() *string
	SetStartTs(v int64) *ListNamingTrackRequest
	GetStartTs() *int64
}

type ListNamingTrackRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The end timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1665299698
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The group.
	//
	// example:
	//
	// group
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-tl32d*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 172.16.183.232
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cd4d3703-e2a6-46b5-85c6-4447e4f****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// Specifies whether to sort the query results in chronological order or reverse chronological order. Default value: `false`.
	//
	// 	- `true`: sorts the query results in reverse chronological order.
	//
	// 	- `false`: sorts the query results in chronological order.
	//
	// example:
	//
	// false
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// fpx-xms-baseinfo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The start timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1660238450
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s ListNamingTrackRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNamingTrackRequest) GoString() string {
	return s.String()
}

func (s *ListNamingTrackRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListNamingTrackRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *ListNamingTrackRequest) GetGroup() *string {
	return s.Group
}

func (s *ListNamingTrackRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNamingTrackRequest) GetIp() *string {
	return s.Ip
}

func (s *ListNamingTrackRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListNamingTrackRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListNamingTrackRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListNamingTrackRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListNamingTrackRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *ListNamingTrackRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListNamingTrackRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *ListNamingTrackRequest) SetAcceptLanguage(v string) *ListNamingTrackRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListNamingTrackRequest) SetEndTs(v int64) *ListNamingTrackRequest {
	s.EndTs = &v
	return s
}

func (s *ListNamingTrackRequest) SetGroup(v string) *ListNamingTrackRequest {
	s.Group = &v
	return s
}

func (s *ListNamingTrackRequest) SetInstanceId(v string) *ListNamingTrackRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNamingTrackRequest) SetIp(v string) *ListNamingTrackRequest {
	s.Ip = &v
	return s
}

func (s *ListNamingTrackRequest) SetNamespaceId(v string) *ListNamingTrackRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListNamingTrackRequest) SetPageNum(v int64) *ListNamingTrackRequest {
	s.PageNum = &v
	return s
}

func (s *ListNamingTrackRequest) SetPageSize(v int64) *ListNamingTrackRequest {
	s.PageSize = &v
	return s
}

func (s *ListNamingTrackRequest) SetRequestPars(v string) *ListNamingTrackRequest {
	s.RequestPars = &v
	return s
}

func (s *ListNamingTrackRequest) SetReverse(v bool) *ListNamingTrackRequest {
	s.Reverse = &v
	return s
}

func (s *ListNamingTrackRequest) SetServiceName(v string) *ListNamingTrackRequest {
	s.ServiceName = &v
	return s
}

func (s *ListNamingTrackRequest) SetStartTs(v int64) *ListNamingTrackRequest {
	s.StartTs = &v
	return s
}

func (s *ListNamingTrackRequest) Validate() error {
	return dara.Validate(s)
}
