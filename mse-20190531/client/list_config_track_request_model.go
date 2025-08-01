// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigTrackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListConfigTrackRequest
	GetAcceptLanguage() *string
	SetDataId(v string) *ListConfigTrackRequest
	GetDataId() *string
	SetEndTs(v int64) *ListConfigTrackRequest
	GetEndTs() *int64
	SetGroup(v string) *ListConfigTrackRequest
	GetGroup() *string
	SetInstanceId(v string) *ListConfigTrackRequest
	GetInstanceId() *string
	SetIp(v string) *ListConfigTrackRequest
	GetIp() *string
	SetNamespaceId(v string) *ListConfigTrackRequest
	GetNamespaceId() *string
	SetPageNum(v int64) *ListConfigTrackRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListConfigTrackRequest
	GetPageSize() *int64
	SetRequestPars(v string) *ListConfigTrackRequest
	GetRequestPars() *string
	SetReverse(v bool) *ListConfigTrackRequest
	GetReverse() *bool
	SetStartTs(v int64) *ListConfigTrackRequest
	GetStartTs() *int64
}

type ListConfigTrackRequest struct {
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
	// The ID of the configuration.
	//
	// example:
	//
	// ballot
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The end timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1657178373
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The name of the configuration group.
	//
	// example:
	//
	// DEFAULT_GROUP
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-i7m2ne****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the listener.
	//
	// example:
	//
	// 192.168.22.2
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// f3a510e2-df52-4fad-9815-42d8bc40****
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
	// Specifies whether to enable reverse ordering. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The start timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1671010148
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s ListConfigTrackRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConfigTrackRequest) GoString() string {
	return s.String()
}

func (s *ListConfigTrackRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListConfigTrackRequest) GetDataId() *string {
	return s.DataId
}

func (s *ListConfigTrackRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *ListConfigTrackRequest) GetGroup() *string {
	return s.Group
}

func (s *ListConfigTrackRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConfigTrackRequest) GetIp() *string {
	return s.Ip
}

func (s *ListConfigTrackRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListConfigTrackRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListConfigTrackRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListConfigTrackRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListConfigTrackRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *ListConfigTrackRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *ListConfigTrackRequest) SetAcceptLanguage(v string) *ListConfigTrackRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListConfigTrackRequest) SetDataId(v string) *ListConfigTrackRequest {
	s.DataId = &v
	return s
}

func (s *ListConfigTrackRequest) SetEndTs(v int64) *ListConfigTrackRequest {
	s.EndTs = &v
	return s
}

func (s *ListConfigTrackRequest) SetGroup(v string) *ListConfigTrackRequest {
	s.Group = &v
	return s
}

func (s *ListConfigTrackRequest) SetInstanceId(v string) *ListConfigTrackRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConfigTrackRequest) SetIp(v string) *ListConfigTrackRequest {
	s.Ip = &v
	return s
}

func (s *ListConfigTrackRequest) SetNamespaceId(v string) *ListConfigTrackRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListConfigTrackRequest) SetPageNum(v int64) *ListConfigTrackRequest {
	s.PageNum = &v
	return s
}

func (s *ListConfigTrackRequest) SetPageSize(v int64) *ListConfigTrackRequest {
	s.PageSize = &v
	return s
}

func (s *ListConfigTrackRequest) SetRequestPars(v string) *ListConfigTrackRequest {
	s.RequestPars = &v
	return s
}

func (s *ListConfigTrackRequest) SetReverse(v bool) *ListConfigTrackRequest {
	s.Reverse = &v
	return s
}

func (s *ListConfigTrackRequest) SetStartTs(v int64) *ListConfigTrackRequest {
	s.StartTs = &v
	return s
}

func (s *ListConfigTrackRequest) Validate() error {
	return dara.Validate(s)
}
