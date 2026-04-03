// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListToolsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListToolsResult
	GetCode() *string
	SetData(v *ListToolsOutputV2) *ListToolsResult
	GetData() *ListToolsOutputV2
	SetRequestId(v string) *ListToolsResult
	GetRequestId() *string
}

type ListToolsResult struct {
	Code      *string            `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ListToolsOutputV2 `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListToolsResult) String() string {
	return dara.Prettify(s)
}

func (s ListToolsResult) GoString() string {
	return s.String()
}

func (s *ListToolsResult) GetCode() *string {
	return s.Code
}

func (s *ListToolsResult) GetData() *ListToolsOutputV2 {
	return s.Data
}

func (s *ListToolsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListToolsResult) SetCode(v string) *ListToolsResult {
	s.Code = &v
	return s
}

func (s *ListToolsResult) SetData(v *ListToolsOutputV2) *ListToolsResult {
	s.Data = v
	return s
}

func (s *ListToolsResult) SetRequestId(v string) *ListToolsResult {
	s.RequestId = &v
	return s
}

func (s *ListToolsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
