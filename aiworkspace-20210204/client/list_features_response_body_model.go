// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatures(v []*string) *ListFeaturesResponseBody
	GetFeatures() []*string
	SetRequestId(v string) *ListFeaturesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListFeaturesResponseBody
	GetTotalCount() *int64
}

type ListFeaturesResponseBody struct {
	Features []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeaturesResponseBody) GetFeatures() []*string {
	return s.Features
}

func (s *ListFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFeaturesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFeaturesResponseBody) SetFeatures(v []*string) *ListFeaturesResponseBody {
	s.Features = v
	return s
}

func (s *ListFeaturesResponseBody) SetRequestId(v string) *ListFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeaturesResponseBody) SetTotalCount(v int64) *ListFeaturesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFeaturesResponseBody) Validate() error {
	return dara.Validate(s)
}
