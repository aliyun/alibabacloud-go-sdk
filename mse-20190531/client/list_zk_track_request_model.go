// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZkTrackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListZkTrackRequest
	GetAcceptLanguage() *string
	SetEndTs(v int64) *ListZkTrackRequest
	GetEndTs() *int64
	SetInstanceId(v string) *ListZkTrackRequest
	GetInstanceId() *string
	SetPageNum(v int64) *ListZkTrackRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListZkTrackRequest
	GetPageSize() *int64
	SetPath(v string) *ListZkTrackRequest
	GetPath() *string
	SetRequestPars(v string) *ListZkTrackRequest
	GetRequestPars() *string
	SetReverse(v bool) *ListZkTrackRequest
	GetReverse() *bool
	SetSessionId(v string) *ListZkTrackRequest
	GetSessionId() *string
	SetStartTs(v int64) *ListZkTrackRequest
	GetStartTs() *int64
}

type ListZkTrackRequest struct {
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
	// 1669619383
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-0ju2yq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The path.
	//
	// example:
	//
	// /path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The request parameters.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// Specifies whether to enable reverse ordering.
	//
	// example:
	//
	// false
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 0x301fdfbdbf00***
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The start timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1669618483
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s ListZkTrackRequest) String() string {
	return dara.Prettify(s)
}

func (s ListZkTrackRequest) GoString() string {
	return s.String()
}

func (s *ListZkTrackRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListZkTrackRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *ListZkTrackRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListZkTrackRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListZkTrackRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListZkTrackRequest) GetPath() *string {
	return s.Path
}

func (s *ListZkTrackRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListZkTrackRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *ListZkTrackRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListZkTrackRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *ListZkTrackRequest) SetAcceptLanguage(v string) *ListZkTrackRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListZkTrackRequest) SetEndTs(v int64) *ListZkTrackRequest {
	s.EndTs = &v
	return s
}

func (s *ListZkTrackRequest) SetInstanceId(v string) *ListZkTrackRequest {
	s.InstanceId = &v
	return s
}

func (s *ListZkTrackRequest) SetPageNum(v int64) *ListZkTrackRequest {
	s.PageNum = &v
	return s
}

func (s *ListZkTrackRequest) SetPageSize(v int64) *ListZkTrackRequest {
	s.PageSize = &v
	return s
}

func (s *ListZkTrackRequest) SetPath(v string) *ListZkTrackRequest {
	s.Path = &v
	return s
}

func (s *ListZkTrackRequest) SetRequestPars(v string) *ListZkTrackRequest {
	s.RequestPars = &v
	return s
}

func (s *ListZkTrackRequest) SetReverse(v bool) *ListZkTrackRequest {
	s.Reverse = &v
	return s
}

func (s *ListZkTrackRequest) SetSessionId(v string) *ListZkTrackRequest {
	s.SessionId = &v
	return s
}

func (s *ListZkTrackRequest) SetStartTs(v int64) *ListZkTrackRequest {
	s.StartTs = &v
	return s
}

func (s *ListZkTrackRequest) Validate() error {
	return dara.Validate(s)
}
