// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPatrolConfigHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetPatrolConfigHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetPatrolConfigHeaders
	GetWorkspace() *string
}

type GetPatrolConfigHeaders struct {
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

func (s GetPatrolConfigHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolConfigHeaders) GoString() string {
	return s.String()
}

func (s *GetPatrolConfigHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetPatrolConfigHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetPatrolConfigHeaders) SetCommonHeaders(v map[string]*string) *GetPatrolConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPatrolConfigHeaders) SetWorkspace(v string) *GetPatrolConfigHeaders {
	s.Workspace = &v
	return s
}

func (s *GetPatrolConfigHeaders) Validate() error {
	return dara.Validate(s)
}
