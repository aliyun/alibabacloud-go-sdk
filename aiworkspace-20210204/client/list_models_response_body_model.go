// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModels(v []*Model) *ListModelsResponseBody
	GetModels() []*Model
	SetRequestId(v string) *ListModelsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListModelsResponseBody
	GetTotalCount() *int64
}

type ListModelsResponseBody struct {
	// The models.
	Models []*Model `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of models.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBody) GetModels() []*Model {
	return s.Models
}

func (s *ListModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListModelsResponseBody) SetModels(v []*Model) *ListModelsResponseBody {
	s.Models = v
	return s
}

func (s *ListModelsResponseBody) SetRequestId(v string) *ListModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelsResponseBody) SetTotalCount(v int64) *ListModelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelsResponseBody) Validate() error {
	return dara.Validate(s)
}
