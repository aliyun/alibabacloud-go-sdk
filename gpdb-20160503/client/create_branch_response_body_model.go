// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *CreateBranchResponseBody
	GetBranchId() *string
	SetProjectId(v string) *CreateBranchResponseBody
	GetProjectId() *string
	SetRequestId(v string) *CreateBranchResponseBody
	GetRequestId() *string
}

type CreateBranchResponseBody struct {
	// The branch ID. This ID uniquely identifies a Supabase branch.
	//
	// example:
	//
	// br-xxxx
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The Supabase project ID that corresponds to the primary branch.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBranchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBody) GetBranchId() *string {
	return s.BranchId
}

func (s *CreateBranchResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBranchResponseBody) SetBranchId(v string) *CreateBranchResponseBody {
	s.BranchId = &v
	return s
}

func (s *CreateBranchResponseBody) SetProjectId(v string) *CreateBranchResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateBranchResponseBody) SetRequestId(v string) *CreateBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBranchResponseBody) Validate() error {
	return dara.Validate(s)
}
