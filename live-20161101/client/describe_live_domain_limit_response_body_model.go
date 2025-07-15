// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveDomainLimitList(v *DescribeLiveDomainLimitResponseBodyLiveDomainLimitList) *DescribeLiveDomainLimitResponseBody
	GetLiveDomainLimitList() *DescribeLiveDomainLimitResponseBodyLiveDomainLimitList
	SetRequestId(v string) *DescribeLiveDomainLimitResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainLimitResponseBody struct {
	// The limits.
	LiveDomainLimitList *DescribeLiveDomainLimitResponseBodyLiveDomainLimitList `json:"LiveDomainLimitList,omitempty" xml:"LiveDomainLimitList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A3136B58-5876-4168-83CA-B562781981A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLimitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLimitResponseBody) GetLiveDomainLimitList() *DescribeLiveDomainLimitResponseBodyLiveDomainLimitList {
	return s.LiveDomainLimitList
}

func (s *DescribeLiveDomainLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainLimitResponseBody) SetLiveDomainLimitList(v *DescribeLiveDomainLimitResponseBodyLiveDomainLimitList) *DescribeLiveDomainLimitResponseBody {
	s.LiveDomainLimitList = v
	return s
}

func (s *DescribeLiveDomainLimitResponseBody) SetRequestId(v string) *DescribeLiveDomainLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainLimitResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainLimitResponseBodyLiveDomainLimitList struct {
	LiveDomainLimit []*DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit `json:"LiveDomainLimit,omitempty" xml:"LiveDomainLimit,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainLimitResponseBodyLiveDomainLimitList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLimitResponseBodyLiveDomainLimitList) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitList) GetLiveDomainLimit() []*DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit {
	return s.LiveDomainLimit
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitList) SetLiveDomainLimit(v []*DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) *DescribeLiveDomainLimitResponseBodyLiveDomainLimitList {
	s.LiveDomainLimit = v
	return s
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit struct {
	// The current number of ingested streams.
	//
	// example:
	//
	// 10
	CurrentNum *int32 `json:"CurrentNum,omitempty" xml:"CurrentNum,omitempty"`
	// The current number of transcoded streams.
	//
	// example:
	//
	// 10
	CurrentTranscodeNum *int32 `json:"CurrentTranscodeNum,omitempty" xml:"CurrentTranscodeNum,omitempty"`
	// The current number of streams relayed from the live center.
	//
	// example:
	//
	// 10
	CurrentTransferNum *int32 `json:"CurrentTransferNum,omitempty" xml:"CurrentTransferNum,omitempty"`
	// The name of the queried main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The maximum number of ingested streams.
	//
	// example:
	//
	// 100
	LimitNum *int32 `json:"LimitNum,omitempty" xml:"LimitNum,omitempty"`
	// The maximum number of transcoded streams.
	//
	// example:
	//
	// 50
	LimitTranscodeNum *int32 `json:"LimitTranscodeNum,omitempty" xml:"LimitTranscodeNum,omitempty"`
	// The maximum allowed number of streams relayed from the live center.
	//
	// example:
	//
	// 50
	LimitTransferNum *int32 `json:"LimitTransferNum,omitempty" xml:"LimitTransferNum,omitempty"`
}

func (s DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) GetCurrentNum() *int32 {
	return s.CurrentNum
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) GetCurrentTranscodeNum() *int32 {
	return s.CurrentTranscodeNum
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) GetCurrentTransferNum() *int32 {
	return s.CurrentTransferNum
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) GetLimitNum() *int32 {
	return s.LimitNum
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) GetLimitTranscodeNum() *int32 {
	return s.LimitTranscodeNum
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) GetLimitTransferNum() *int32 {
	return s.LimitTransferNum
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) SetCurrentNum(v int32) *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit {
	s.CurrentNum = &v
	return s
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) SetCurrentTranscodeNum(v int32) *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit {
	s.CurrentTranscodeNum = &v
	return s
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) SetCurrentTransferNum(v int32) *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit {
	s.CurrentTransferNum = &v
	return s
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) SetDomainName(v string) *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) SetLimitNum(v int32) *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit {
	s.LimitNum = &v
	return s
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) SetLimitTranscodeNum(v int32) *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit {
	s.LimitTranscodeNum = &v
	return s
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) SetLimitTransferNum(v int32) *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit {
	s.LimitTransferNum = &v
	return s
}

func (s *DescribeLiveDomainLimitResponseBodyLiveDomainLimitListLiveDomainLimit) Validate() error {
	return dara.Validate(s)
}
