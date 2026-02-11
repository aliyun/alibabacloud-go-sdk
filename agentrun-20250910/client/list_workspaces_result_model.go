// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWorkspacesResult
	GetCode() *string
	SetData(v *ListWorkspacesOutput) *ListWorkspacesResult
	GetData() *ListWorkspacesOutput
	SetRequestId(v string) *ListWorkspacesResult
	GetRequestId() *string
}

type ListWorkspacesResult struct {
	Code      *string               `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ListWorkspacesOutput `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string               `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListWorkspacesResult) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResult) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResult) GetCode() *string {
	return s.Code
}

func (s *ListWorkspacesResult) GetData() *ListWorkspacesOutput {
	return s.Data
}

func (s *ListWorkspacesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspacesResult) SetCode(v string) *ListWorkspacesResult {
	s.Code = &v
	return s
}

func (s *ListWorkspacesResult) SetData(v *ListWorkspacesOutput) *ListWorkspacesResult {
	s.Data = v
	return s
}

func (s *ListWorkspacesResult) SetRequestId(v string) *ListWorkspacesResult {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
