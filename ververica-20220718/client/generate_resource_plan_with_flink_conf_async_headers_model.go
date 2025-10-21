// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateResourcePlanWithFlinkConfAsyncHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GenerateResourcePlanWithFlinkConfAsyncHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GenerateResourcePlanWithFlinkConfAsyncHeaders
	GetWorkspace() *string
}

type GenerateResourcePlanWithFlinkConfAsyncHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GenerateResourcePlanWithFlinkConfAsyncHeaders) String() string {
	return dara.Prettify(s)
}

func (s GenerateResourcePlanWithFlinkConfAsyncHeaders) GoString() string {
	return s.String()
}

func (s *GenerateResourcePlanWithFlinkConfAsyncHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GenerateResourcePlanWithFlinkConfAsyncHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GenerateResourcePlanWithFlinkConfAsyncHeaders) SetCommonHeaders(v map[string]*string) *GenerateResourcePlanWithFlinkConfAsyncHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncHeaders) SetWorkspace(v string) *GenerateResourcePlanWithFlinkConfAsyncHeaders {
	s.Workspace = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncHeaders) Validate() error {
	return dara.Validate(s)
}
