// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIndices(v *ListIndicesResponseBodyIndices) *ListIndicesResponseBody
	GetIndices() *ListIndicesResponseBodyIndices
	SetMessage(v string) *ListIndicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIndicesResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListIndicesResponseBody
	GetStatus() *string
}

type ListIndicesResponseBody struct {
	Indices *ListIndicesResponseBodyIndices `json:"Indices,omitempty" xml:"Indices,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIndicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIndicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndicesResponseBody) GetIndices() *ListIndicesResponseBodyIndices {
	return s.Indices
}

func (s *ListIndicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIndicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIndicesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListIndicesResponseBody) SetIndices(v *ListIndicesResponseBodyIndices) *ListIndicesResponseBody {
	s.Indices = v
	return s
}

func (s *ListIndicesResponseBody) SetMessage(v string) *ListIndicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListIndicesResponseBody) SetRequestId(v string) *ListIndicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndicesResponseBody) SetStatus(v string) *ListIndicesResponseBody {
	s.Status = &v
	return s
}

func (s *ListIndicesResponseBody) Validate() error {
	if s.Indices != nil {
		if err := s.Indices.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIndicesResponseBodyIndices struct {
	Indices []*ListIndicesResponseBodyIndicesIndices `json:"Indices,omitempty" xml:"Indices,omitempty" type:"Repeated"`
}

func (s ListIndicesResponseBodyIndices) String() string {
	return dara.Prettify(s)
}

func (s ListIndicesResponseBodyIndices) GoString() string {
	return s.String()
}

func (s *ListIndicesResponseBodyIndices) GetIndices() []*ListIndicesResponseBodyIndicesIndices {
	return s.Indices
}

func (s *ListIndicesResponseBodyIndices) SetIndices(v []*ListIndicesResponseBodyIndicesIndices) *ListIndicesResponseBodyIndices {
	s.Indices = v
	return s
}

func (s *ListIndicesResponseBodyIndices) Validate() error {
	if s.Indices != nil {
		for _, item := range s.Indices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIndicesResponseBodyIndicesIndices struct {
	// example:
	//
	// testcollection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// example:
	//
	// CREATE INDEX testindex ON mynamespace. testcollection
	IndexDef *string `json:"IndexDef,omitempty" xml:"IndexDef,omitempty"`
	// example:
	//
	// testindex
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListIndicesResponseBodyIndicesIndices) String() string {
	return dara.Prettify(s)
}

func (s ListIndicesResponseBodyIndicesIndices) GoString() string {
	return s.String()
}

func (s *ListIndicesResponseBodyIndicesIndices) GetCollection() *string {
	return s.Collection
}

func (s *ListIndicesResponseBodyIndicesIndices) GetIndexDef() *string {
	return s.IndexDef
}

func (s *ListIndicesResponseBodyIndicesIndices) GetIndexName() *string {
	return s.IndexName
}

func (s *ListIndicesResponseBodyIndicesIndices) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIndicesResponseBodyIndicesIndices) SetCollection(v string) *ListIndicesResponseBodyIndicesIndices {
	s.Collection = &v
	return s
}

func (s *ListIndicesResponseBodyIndicesIndices) SetIndexDef(v string) *ListIndicesResponseBodyIndicesIndices {
	s.IndexDef = &v
	return s
}

func (s *ListIndicesResponseBodyIndicesIndices) SetIndexName(v string) *ListIndicesResponseBodyIndicesIndices {
	s.IndexName = &v
	return s
}

func (s *ListIndicesResponseBodyIndicesIndices) SetNamespace(v string) *ListIndicesResponseBodyIndicesIndices {
	s.Namespace = &v
	return s
}

func (s *ListIndicesResponseBodyIndicesIndices) Validate() error {
	return dara.Validate(s)
}
