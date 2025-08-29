// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTrafficControlTargetItemReportDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryTrafficControlTargetItemReportDetailResponseBody
	GetRequestId() *string
	SetTrafficControlTargetItemReportDetail(v *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) *QueryTrafficControlTargetItemReportDetailResponseBody
	GetTrafficControlTargetItemReportDetail() *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail
}

type QueryTrafficControlTargetItemReportDetailResponseBody struct {
	RequestId                            *string                                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficControlTargetItemReportDetail *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail `json:"TrafficControlTargetItemReportDetail,omitempty" xml:"TrafficControlTargetItemReportDetail,omitempty" type:"Struct"`
}

func (s QueryTrafficControlTargetItemReportDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBody) GetTrafficControlTargetItemReportDetail() *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail {
	return s.TrafficControlTargetItemReportDetail
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBody) SetRequestId(v string) *QueryTrafficControlTargetItemReportDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBody) SetTrafficControlTargetItemReportDetail(v *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) *QueryTrafficControlTargetItemReportDetailResponseBody {
	s.TrafficControlTargetItemReportDetail = v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail struct {
	ItemControlTailReportDetails []*QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails `json:"ItemControlTailReportDetails,omitempty" xml:"ItemControlTailReportDetails,omitempty" type:"Repeated"`
	ItemControlTopReportDetails  []*QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails  `json:"ItemControlTopReportDetails,omitempty" xml:"ItemControlTopReportDetails,omitempty" type:"Repeated"`
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) GetItemControlTailReportDetails() []*QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails {
	return s.ItemControlTailReportDetails
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) GetItemControlTopReportDetails() []*QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails {
	return s.ItemControlTopReportDetails
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) SetItemControlTailReportDetails(v []*QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail {
	s.ItemControlTailReportDetails = v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) SetItemControlTopReportDetails(v []*QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail {
	s.ItemControlTopReportDetails = v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetail) Validate() error {
	return dara.Validate(s)
}

type QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails struct {
	Features       map[string]interface{} `json:"Features,omitempty" xml:"Features,omitempty"`
	ItemId         *string                `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	TargetProgress *string                `json:"TargetProgress,omitempty" xml:"TargetProgress,omitempty"`
	TargetTraffic  *int64                 `json:"TargetTraffic,omitempty" xml:"TargetTraffic,omitempty"`
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) GetItemId() *string {
	return s.ItemId
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) GetTargetProgress() *string {
	return s.TargetProgress
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) GetTargetTraffic() *int64 {
	return s.TargetTraffic
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) SetFeatures(v map[string]interface{}) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails {
	s.Features = v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) SetItemId(v string) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails {
	s.ItemId = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) SetTargetProgress(v string) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails {
	s.TargetProgress = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) SetTargetTraffic(v int64) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails {
	s.TargetTraffic = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTailReportDetails) Validate() error {
	return dara.Validate(s)
}

type QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails struct {
	Features       map[string]interface{} `json:"Features,omitempty" xml:"Features,omitempty"`
	ItemId         *string                `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	TargetProgress *string                `json:"TargetProgress,omitempty" xml:"TargetProgress,omitempty"`
	TargetTraffic  *int64                 `json:"TargetTraffic,omitempty" xml:"TargetTraffic,omitempty"`
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) GetItemId() *string {
	return s.ItemId
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) GetTargetProgress() *string {
	return s.TargetProgress
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) GetTargetTraffic() *int64 {
	return s.TargetTraffic
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) SetFeatures(v map[string]interface{}) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails {
	s.Features = v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) SetItemId(v string) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails {
	s.ItemId = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) SetTargetProgress(v string) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails {
	s.TargetProgress = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) SetTargetTraffic(v int64) *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails {
	s.TargetTraffic = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponseBodyTrafficControlTargetItemReportDetailItemControlTopReportDetails) Validate() error {
	return dara.Validate(s)
}
