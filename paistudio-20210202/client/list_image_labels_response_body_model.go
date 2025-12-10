// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*ListImageLabelsResponseBodyLabels) *ListImageLabelsResponseBody
	GetLabels() []*ListImageLabelsResponseBodyLabels
	SetTotalCount(v int64) *ListImageLabelsResponseBody
	GetTotalCount() *int64
	SetRequestId(v string) *ListImageLabelsResponseBody
	GetRequestId() *string
}

type ListImageLabelsResponseBody struct {
	Labels []*ListImageLabelsResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListImageLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageLabelsResponseBody) GetLabels() []*ListImageLabelsResponseBodyLabels {
	return s.Labels
}

func (s *ListImageLabelsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListImageLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageLabelsResponseBody) SetLabels(v []*ListImageLabelsResponseBodyLabels) *ListImageLabelsResponseBody {
	s.Labels = v
	return s
}

func (s *ListImageLabelsResponseBody) SetTotalCount(v int64) *ListImageLabelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImageLabelsResponseBody) SetRequestId(v string) *ListImageLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageLabelsResponseBody) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImageLabelsResponseBodyLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListImageLabelsResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s ListImageLabelsResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *ListImageLabelsResponseBodyLabels) GetKey() *string {
	return s.Key
}

func (s *ListImageLabelsResponseBodyLabels) GetValue() *string {
	return s.Value
}

func (s *ListImageLabelsResponseBodyLabels) SetKey(v string) *ListImageLabelsResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *ListImageLabelsResponseBodyLabels) SetValue(v string) *ListImageLabelsResponseBodyLabels {
	s.Value = &v
	return s
}

func (s *ListImageLabelsResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}
