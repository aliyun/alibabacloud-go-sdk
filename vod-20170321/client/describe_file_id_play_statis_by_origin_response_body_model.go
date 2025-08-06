// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileIdPlayStatisByOriginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilePlayStatisList(v []*DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) *DescribeFileIdPlayStatisByOriginResponseBody
	GetFilePlayStatisList() []*DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList
	SetHasNext(v bool) *DescribeFileIdPlayStatisByOriginResponseBody
	GetHasNext() *bool
	SetRequestId(v string) *DescribeFileIdPlayStatisByOriginResponseBody
	GetRequestId() *string
	SetScrollToken(v string) *DescribeFileIdPlayStatisByOriginResponseBody
	GetScrollToken() *string
}

type DescribeFileIdPlayStatisByOriginResponseBody struct {
	FilePlayStatisList []*DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList `json:"FilePlayStatisList,omitempty" xml:"FilePlayStatisList,omitempty" type:"Repeated"`
	HasNext            *bool                                                             `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	RequestId          *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScrollToken        *string                                                           `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
}

func (s DescribeFileIdPlayStatisByOriginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileIdPlayStatisByOriginResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileIdPlayStatisByOriginResponseBody) GetFilePlayStatisList() []*DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList {
	return s.FilePlayStatisList
}

func (s *DescribeFileIdPlayStatisByOriginResponseBody) GetHasNext() *bool {
	return s.HasNext
}

func (s *DescribeFileIdPlayStatisByOriginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFileIdPlayStatisByOriginResponseBody) GetScrollToken() *string {
	return s.ScrollToken
}

func (s *DescribeFileIdPlayStatisByOriginResponseBody) SetFilePlayStatisList(v []*DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) *DescribeFileIdPlayStatisByOriginResponseBody {
	s.FilePlayStatisList = v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponseBody) SetHasNext(v bool) *DescribeFileIdPlayStatisByOriginResponseBody {
	s.HasNext = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponseBody) SetRequestId(v string) *DescribeFileIdPlayStatisByOriginResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponseBody) SetScrollToken(v string) *DescribeFileIdPlayStatisByOriginResponseBody {
	s.ScrollToken = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList struct {
	FileId       *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Flux         *int64  `json:"Flux,omitempty" xml:"Flux,omitempty"`
	PlayCount    *int64  `json:"PlayCount,omitempty" xml:"PlayCount,omitempty"`
	RequestCount *int64  `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	StatisTime   *string `json:"StatisTime,omitempty" xml:"StatisTime,omitempty"`
}

func (s DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) GoString() string {
	return s.String()
}

func (s *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) GetFileId() *string {
	return s.FileId
}

func (s *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) GetFlux() *int64 {
	return s.Flux
}

func (s *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) GetPlayCount() *int64 {
	return s.PlayCount
}

func (s *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) GetRequestCount() *int64 {
	return s.RequestCount
}

func (s *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) GetStatisTime() *string {
	return s.StatisTime
}

func (s *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) SetFileId(v string) *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList {
	s.FileId = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) SetFlux(v int64) *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList {
	s.Flux = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) SetPlayCount(v int64) *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList {
	s.PlayCount = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) SetRequestCount(v int64) *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList {
	s.RequestCount = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) SetStatisTime(v string) *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList {
	s.StatisTime = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginResponseBodyFilePlayStatisList) Validate() error {
	return dara.Validate(s)
}
