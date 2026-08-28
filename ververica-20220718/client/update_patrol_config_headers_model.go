// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePatrolConfigHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdatePatrolConfigHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdatePatrolConfigHeaders
	GetWorkspace() *string
}

type UpdatePatrolConfigHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdatePatrolConfigHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatrolConfigHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePatrolConfigHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdatePatrolConfigHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdatePatrolConfigHeaders) SetCommonHeaders(v map[string]*string) *UpdatePatrolConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePatrolConfigHeaders) SetWorkspace(v string) *UpdatePatrolConfigHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdatePatrolConfigHeaders) Validate() error {
	return dara.Validate(s)
}
