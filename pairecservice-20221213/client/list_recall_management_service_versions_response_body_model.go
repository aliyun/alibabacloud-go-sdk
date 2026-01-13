// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementServiceVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecallManagementServiceVersions(v []*ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) *ListRecallManagementServiceVersionsResponseBody
	GetRecallManagementServiceVersions() []*ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions
	SetRequestId(v string) *ListRecallManagementServiceVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRecallManagementServiceVersionsResponseBody
	GetTotalCount() *string
}

type ListRecallManagementServiceVersionsResponseBody struct {
	RecallManagementServiceVersions []*ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions `json:"RecallManagementServiceVersions,omitempty" xml:"RecallManagementServiceVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecallManagementServiceVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementServiceVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecallManagementServiceVersionsResponseBody) GetRecallManagementServiceVersions() []*ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions {
	return s.RecallManagementServiceVersions
}

func (s *ListRecallManagementServiceVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecallManagementServiceVersionsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRecallManagementServiceVersionsResponseBody) SetRecallManagementServiceVersions(v []*ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) *ListRecallManagementServiceVersionsResponseBody {
	s.RecallManagementServiceVersions = v
	return s
}

func (s *ListRecallManagementServiceVersionsResponseBody) SetRequestId(v string) *ListRecallManagementServiceVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecallManagementServiceVersionsResponseBody) SetTotalCount(v string) *ListRecallManagementServiceVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecallManagementServiceVersionsResponseBody) Validate() error {
	if s.RecallManagementServiceVersions != nil {
		for _, item := range s.RecallManagementServiceVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions struct {
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// false
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// version-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	RecallManagementServiceVersionId *string `json:"RecallManagementServiceVersionId,omitempty" xml:"RecallManagementServiceVersionId,omitempty"`
}

func (s ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) GoString() string {
	return s.String()
}

func (s *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) GetIsDefault() *string {
	return s.IsDefault
}

func (s *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) GetName() *string {
	return s.Name
}

func (s *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) GetRecallManagementServiceVersionId() *string {
	return s.RecallManagementServiceVersionId
}

func (s *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) SetGmtCreateTime(v string) *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions {
	s.GmtCreateTime = &v
	return s
}

func (s *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) SetGmtModifiedTime(v string) *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) SetIsDefault(v string) *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions {
	s.IsDefault = &v
	return s
}

func (s *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) SetName(v string) *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions {
	s.Name = &v
	return s
}

func (s *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) SetRecallManagementServiceVersionId(v string) *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions {
	s.RecallManagementServiceVersionId = &v
	return s
}

func (s *ListRecallManagementServiceVersionsResponseBodyRecallManagementServiceVersions) Validate() error {
	return dara.Validate(s)
}
