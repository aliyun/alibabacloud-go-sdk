// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileIdPlayStatisByEdgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilePlayStatisList(v []*DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) *DescribeFileIdPlayStatisByEdgeResponseBody
	GetFilePlayStatisList() []*DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList
	SetHasNext(v bool) *DescribeFileIdPlayStatisByEdgeResponseBody
	GetHasNext() *bool
	SetRequestId(v string) *DescribeFileIdPlayStatisByEdgeResponseBody
	GetRequestId() *string
	SetScrollToken(v string) *DescribeFileIdPlayStatisByEdgeResponseBody
	GetScrollToken() *string
}

type DescribeFileIdPlayStatisByEdgeResponseBody struct {
	FilePlayStatisList []*DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList `json:"FilePlayStatisList,omitempty" xml:"FilePlayStatisList,omitempty" type:"Repeated"`
	HasNext            *bool                                                           `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	RequestId          *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScrollToken        *string                                                         `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
}

func (s DescribeFileIdPlayStatisByEdgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileIdPlayStatisByEdgeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBody) GetFilePlayStatisList() []*DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList {
	return s.FilePlayStatisList
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBody) GetHasNext() *bool {
	return s.HasNext
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBody) GetScrollToken() *string {
	return s.ScrollToken
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBody) SetFilePlayStatisList(v []*DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) *DescribeFileIdPlayStatisByEdgeResponseBody {
	s.FilePlayStatisList = v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBody) SetHasNext(v bool) *DescribeFileIdPlayStatisByEdgeResponseBody {
	s.HasNext = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBody) SetRequestId(v string) *DescribeFileIdPlayStatisByEdgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBody) SetScrollToken(v string) *DescribeFileIdPlayStatisByEdgeResponseBody {
	s.ScrollToken = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList struct {
	FileId       *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Flux         *int64  `json:"Flux,omitempty" xml:"Flux,omitempty"`
	PlayCount    *int64  `json:"PlayCount,omitempty" xml:"PlayCount,omitempty"`
	RequestCount *int64  `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	StatisTime   *string `json:"StatisTime,omitempty" xml:"StatisTime,omitempty"`
}

func (s DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) GoString() string {
	return s.String()
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) GetFileId() *string {
	return s.FileId
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) GetFlux() *int64 {
	return s.Flux
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) GetPlayCount() *int64 {
	return s.PlayCount
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) GetRequestCount() *int64 {
	return s.RequestCount
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) GetStatisTime() *string {
	return s.StatisTime
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) SetFileId(v string) *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList {
	s.FileId = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) SetFlux(v int64) *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList {
	s.Flux = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) SetPlayCount(v int64) *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList {
	s.PlayCount = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) SetRequestCount(v int64) *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList {
	s.RequestCount = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) SetStatisTime(v string) *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList {
	s.StatisTime = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeResponseBodyFilePlayStatisList) Validate() error {
	return dara.Validate(s)
}
