// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementTableVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecallManagementTableVersions(v []*ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) *ListRecallManagementTableVersionsResponseBody
	GetRecallManagementTableVersions() []*ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions
	SetRequestId(v string) *ListRecallManagementTableVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRecallManagementTableVersionsResponseBody
	GetTotalCount() *int64
}

type ListRecallManagementTableVersionsResponseBody struct {
	RecallManagementTableVersions []*ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions `json:"RecallManagementTableVersions,omitempty" xml:"RecallManagementTableVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecallManagementTableVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementTableVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecallManagementTableVersionsResponseBody) GetRecallManagementTableVersions() []*ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions {
	return s.RecallManagementTableVersions
}

func (s *ListRecallManagementTableVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecallManagementTableVersionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRecallManagementTableVersionsResponseBody) SetRecallManagementTableVersions(v []*ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) *ListRecallManagementTableVersionsResponseBody {
	s.RecallManagementTableVersions = v
	return s
}

func (s *ListRecallManagementTableVersionsResponseBody) SetRequestId(v string) *ListRecallManagementTableVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecallManagementTableVersionsResponseBody) SetTotalCount(v int64) *ListRecallManagementTableVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecallManagementTableVersionsResponseBody) Validate() error {
	if s.RecallManagementTableVersions != nil {
		for _, item := range s.RecallManagementTableVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions struct {
	// example:
	//
	// ds=20250701
	DataVersion *string `json:"DataVersion,omitempty" xml:"DataVersion,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132
	PublishEndTime *string `json:"PublishEndTime,omitempty" xml:"PublishEndTime,omitempty"`
	// example:
	//
	// 2021-12-15T22:24:33.132
	PublishStartTime *string `json:"PublishStartTime,omitempty" xml:"PublishStartTime,omitempty"`
	// example:
	//
	// 202507010000
	RecallManagementTableVersionId *string `json:"RecallManagementTableVersionId,omitempty" xml:"RecallManagementTableVersionId,omitempty"`
	// example:
	//
	// 1000
	SourceTableDataSize *int64 `json:"SourceTableDataSize,omitempty" xml:"SourceTableDataSize,omitempty"`
	// example:
	//
	// 100
	SourceTableRowCount *int64 `json:"SourceTableRowCount,omitempty" xml:"SourceTableRowCount,omitempty"`
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) GoString() string {
	return s.String()
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) GetDataVersion() *string {
	return s.DataVersion
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) GetPublishEndTime() *string {
	return s.PublishEndTime
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) GetPublishStartTime() *string {
	return s.PublishStartTime
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) GetRecallManagementTableVersionId() *string {
	return s.RecallManagementTableVersionId
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) GetSourceTableDataSize() *int64 {
	return s.SourceTableDataSize
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) GetSourceTableRowCount() *int64 {
	return s.SourceTableRowCount
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) GetStatus() *string {
	return s.Status
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) SetDataVersion(v string) *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions {
	s.DataVersion = &v
	return s
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) SetEffectiveTime(v string) *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions {
	s.EffectiveTime = &v
	return s
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) SetPublishEndTime(v string) *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions {
	s.PublishEndTime = &v
	return s
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) SetPublishStartTime(v string) *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions {
	s.PublishStartTime = &v
	return s
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) SetRecallManagementTableVersionId(v string) *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions {
	s.RecallManagementTableVersionId = &v
	return s
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) SetSourceTableDataSize(v int64) *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions {
	s.SourceTableDataSize = &v
	return s
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) SetSourceTableRowCount(v int64) *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions {
	s.SourceTableRowCount = &v
	return s
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) SetStatus(v string) *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions {
	s.Status = &v
	return s
}

func (s *ListRecallManagementTableVersionsResponseBodyRecallManagementTableVersions) Validate() error {
	return dara.Validate(s)
}
