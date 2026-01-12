// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryCollectionsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemoryCollectionsResult
	GetCode() *string
	SetData(v *ListMemoryCollectionsOutput) *ListMemoryCollectionsResult
	GetData() *ListMemoryCollectionsOutput
	SetRequestId(v string) *ListMemoryCollectionsResult
	GetRequestId() *string
}

type ListMemoryCollectionsResult struct {
	Code      *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ListMemoryCollectionsOutput `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMemoryCollectionsResult) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryCollectionsResult) GoString() string {
	return s.String()
}

func (s *ListMemoryCollectionsResult) GetCode() *string {
	return s.Code
}

func (s *ListMemoryCollectionsResult) GetData() *ListMemoryCollectionsOutput {
	return s.Data
}

func (s *ListMemoryCollectionsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemoryCollectionsResult) SetCode(v string) *ListMemoryCollectionsResult {
	s.Code = &v
	return s
}

func (s *ListMemoryCollectionsResult) SetData(v *ListMemoryCollectionsOutput) *ListMemoryCollectionsResult {
	s.Data = v
	return s
}

func (s *ListMemoryCollectionsResult) SetRequestId(v string) *ListMemoryCollectionsResult {
	s.RequestId = &v
	return s
}

func (s *ListMemoryCollectionsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
