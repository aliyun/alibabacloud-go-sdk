// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStaticsQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetLength(v int64) *ListStaticsQueryResponse
	GetLength() *int64
	SetSort(v string) *ListStaticsQueryResponse
	GetSort() *string
	SetStatics(v []*StaticsInfo) *ListStaticsQueryResponse
	GetStatics() []*StaticsInfo
}

type ListStaticsQueryResponse struct {
	Length  *int64         `json:"length,omitempty" xml:"length,omitempty"`
	Sort    *string        `json:"sort,omitempty" xml:"sort,omitempty"`
	Statics []*StaticsInfo `json:"statics,omitempty" xml:"statics,omitempty" type:"Repeated"`
}

func (s ListStaticsQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStaticsQueryResponse) GoString() string {
	return s.String()
}

func (s *ListStaticsQueryResponse) GetLength() *int64 {
	return s.Length
}

func (s *ListStaticsQueryResponse) GetSort() *string {
	return s.Sort
}

func (s *ListStaticsQueryResponse) GetStatics() []*StaticsInfo {
	return s.Statics
}

func (s *ListStaticsQueryResponse) SetLength(v int64) *ListStaticsQueryResponse {
	s.Length = &v
	return s
}

func (s *ListStaticsQueryResponse) SetSort(v string) *ListStaticsQueryResponse {
	s.Sort = &v
	return s
}

func (s *ListStaticsQueryResponse) SetStatics(v []*StaticsInfo) *ListStaticsQueryResponse {
	s.Statics = v
	return s
}

func (s *ListStaticsQueryResponse) Validate() error {
	return dara.Validate(s)
}
