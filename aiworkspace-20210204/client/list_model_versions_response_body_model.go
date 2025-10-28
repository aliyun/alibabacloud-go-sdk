// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListModelVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListModelVersionsResponseBody
	GetTotalCount() *int64
	SetVersions(v []*ModelVersion) *ListModelVersionsResponseBody
	GetVersions() []*ModelVersion
}

type ListModelVersionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC***3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of model versions.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The model versions.
	Versions []*ModelVersion `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListModelVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelVersionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListModelVersionsResponseBody) GetVersions() []*ModelVersion {
	return s.Versions
}

func (s *ListModelVersionsResponseBody) SetRequestId(v string) *ListModelVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelVersionsResponseBody) SetTotalCount(v int64) *ListModelVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelVersionsResponseBody) SetVersions(v []*ModelVersion) *ListModelVersionsResponseBody {
	s.Versions = v
	return s
}

func (s *ListModelVersionsResponseBody) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
