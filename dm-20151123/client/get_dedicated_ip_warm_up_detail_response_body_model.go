// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDedicatedIpWarmUpDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v []*GetDedicatedIpWarmUpDetailResponseBodyDetail) *GetDedicatedIpWarmUpDetailResponseBody
	GetDetail() []*GetDedicatedIpWarmUpDetailResponseBodyDetail
	SetRequestId(v string) *GetDedicatedIpWarmUpDetailResponseBody
	GetRequestId() *string
}

type GetDedicatedIpWarmUpDetailResponseBody struct {
	Detail    []*GetDedicatedIpWarmUpDetailResponseBodyDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDedicatedIpWarmUpDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDedicatedIpWarmUpDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDedicatedIpWarmUpDetailResponseBody) GetDetail() []*GetDedicatedIpWarmUpDetailResponseBodyDetail {
	return s.Detail
}

func (s *GetDedicatedIpWarmUpDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDedicatedIpWarmUpDetailResponseBody) SetDetail(v []*GetDedicatedIpWarmUpDetailResponseBodyDetail) *GetDedicatedIpWarmUpDetailResponseBody {
	s.Detail = v
	return s
}

func (s *GetDedicatedIpWarmUpDetailResponseBody) SetRequestId(v string) *GetDedicatedIpWarmUpDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDedicatedIpWarmUpDetailResponseBody) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDedicatedIpWarmUpDetailResponseBodyDetail struct {
	DayMark       *int64  `json:"DayMark,omitempty" xml:"DayMark,omitempty"`
	DeliverCounts *int64  `json:"DeliverCounts,omitempty" xml:"DeliverCounts,omitempty"`
	Esp           *string `json:"Esp,omitempty" xml:"Esp,omitempty"`
	SendCounts    *int64  `json:"SendCounts,omitempty" xml:"SendCounts,omitempty"`
}

func (s GetDedicatedIpWarmUpDetailResponseBodyDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDedicatedIpWarmUpDetailResponseBodyDetail) GoString() string {
	return s.String()
}

func (s *GetDedicatedIpWarmUpDetailResponseBodyDetail) GetDayMark() *int64 {
	return s.DayMark
}

func (s *GetDedicatedIpWarmUpDetailResponseBodyDetail) GetDeliverCounts() *int64 {
	return s.DeliverCounts
}

func (s *GetDedicatedIpWarmUpDetailResponseBodyDetail) GetEsp() *string {
	return s.Esp
}

func (s *GetDedicatedIpWarmUpDetailResponseBodyDetail) GetSendCounts() *int64 {
	return s.SendCounts
}

func (s *GetDedicatedIpWarmUpDetailResponseBodyDetail) SetDayMark(v int64) *GetDedicatedIpWarmUpDetailResponseBodyDetail {
	s.DayMark = &v
	return s
}

func (s *GetDedicatedIpWarmUpDetailResponseBodyDetail) SetDeliverCounts(v int64) *GetDedicatedIpWarmUpDetailResponseBodyDetail {
	s.DeliverCounts = &v
	return s
}

func (s *GetDedicatedIpWarmUpDetailResponseBodyDetail) SetEsp(v string) *GetDedicatedIpWarmUpDetailResponseBodyDetail {
	s.Esp = &v
	return s
}

func (s *GetDedicatedIpWarmUpDetailResponseBodyDetail) SetSendCounts(v int64) *GetDedicatedIpWarmUpDetailResponseBodyDetail {
	s.SendCounts = &v
	return s
}

func (s *GetDedicatedIpWarmUpDetailResponseBodyDetail) Validate() error {
	return dara.Validate(s)
}
