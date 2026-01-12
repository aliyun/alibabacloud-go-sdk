// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMemoryCollectionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MemoryCollectionResult
	GetCode() *string
	SetData(v *MemoryCollection) *MemoryCollectionResult
	GetData() *MemoryCollection
	SetRequestId(v string) *MemoryCollectionResult
	GetRequestId() *string
}

type MemoryCollectionResult struct {
	Code      *string           `json:"code,omitempty" xml:"code,omitempty"`
	Data      *MemoryCollection `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string           `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s MemoryCollectionResult) String() string {
	return dara.Prettify(s)
}

func (s MemoryCollectionResult) GoString() string {
	return s.String()
}

func (s *MemoryCollectionResult) GetCode() *string {
	return s.Code
}

func (s *MemoryCollectionResult) GetData() *MemoryCollection {
	return s.Data
}

func (s *MemoryCollectionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MemoryCollectionResult) SetCode(v string) *MemoryCollectionResult {
	s.Code = &v
	return s
}

func (s *MemoryCollectionResult) SetData(v *MemoryCollection) *MemoryCollectionResult {
	s.Data = v
	return s
}

func (s *MemoryCollectionResult) SetRequestId(v string) *MemoryCollectionResult {
	s.RequestId = &v
	return s
}

func (s *MemoryCollectionResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
