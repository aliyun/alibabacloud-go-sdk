// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainSnapshotDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveDomainSnapshotDataResponseBody
	GetRequestId() *string
	SetSnapshotDataInfos(v *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos) *DescribeLiveDomainSnapshotDataResponseBody
	GetSnapshotDataInfos() *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos
}

type DescribeLiveDomainSnapshotDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B955107D-E658-4E77-B913-E0AC3D31693F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The daily statistics on the number of snapshots.
	SnapshotDataInfos *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos `json:"SnapshotDataInfos,omitempty" xml:"SnapshotDataInfos,omitempty" type:"Struct"`
}

func (s DescribeLiveDomainSnapshotDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainSnapshotDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainSnapshotDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainSnapshotDataResponseBody) GetSnapshotDataInfos() *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos {
	return s.SnapshotDataInfos
}

func (s *DescribeLiveDomainSnapshotDataResponseBody) SetRequestId(v string) *DescribeLiveDomainSnapshotDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataResponseBody) SetSnapshotDataInfos(v *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos) *DescribeLiveDomainSnapshotDataResponseBody {
	s.SnapshotDataInfos = v
	return s
}

func (s *DescribeLiveDomainSnapshotDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos struct {
	SnapshotDataInfo []*DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo `json:"SnapshotDataInfo,omitempty" xml:"SnapshotDataInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos) GetSnapshotDataInfo() []*DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo {
	return s.SnapshotDataInfo
}

func (s *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos) SetSnapshotDataInfo(v []*DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo) *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos {
	s.SnapshotDataInfo = v
	return s
}

func (s *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo struct {
	// The date.
	//
	// example:
	//
	// 20180209
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The total number of snapshots that were captured on the day.
	//
	// example:
	//
	// 110
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo) GetDate() *string {
	return s.Date
}

func (s *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo) SetDate(v string) *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo {
	s.Date = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo) SetTotal(v int32) *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo {
	s.Total = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataResponseBodySnapshotDataInfosSnapshotDataInfo) Validate() error {
	return dara.Validate(s)
}
