// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *WorkspaceResult
	GetCode() *string
	SetData(v *Workspace) *WorkspaceResult
	GetData() *Workspace
	SetRequestId(v string) *WorkspaceResult
	GetRequestId() *string
}

type WorkspaceResult struct {
	Code      *string    `json:"code,omitempty" xml:"code,omitempty"`
	Data      *Workspace `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string    `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s WorkspaceResult) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceResult) GoString() string {
	return s.String()
}

func (s *WorkspaceResult) GetCode() *string {
	return s.Code
}

func (s *WorkspaceResult) GetData() *Workspace {
	return s.Data
}

func (s *WorkspaceResult) GetRequestId() *string {
	return s.RequestId
}

func (s *WorkspaceResult) SetCode(v string) *WorkspaceResult {
	s.Code = &v
	return s
}

func (s *WorkspaceResult) SetData(v *Workspace) *WorkspaceResult {
	s.Data = v
	return s
}

func (s *WorkspaceResult) SetRequestId(v string) *WorkspaceResult {
	s.RequestId = &v
	return s
}

func (s *WorkspaceResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
