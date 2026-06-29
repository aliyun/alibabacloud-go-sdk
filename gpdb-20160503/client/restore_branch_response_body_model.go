// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *RestoreBranchResponseBody
	GetBranchId() *string
	SetRequestId(v string) *RestoreBranchResponseBody
	GetRequestId() *string
}

type RestoreBranchResponseBody struct {
	// The branch ID that uniquely identifies a Supabase branch.
	//
	// example:
	//
	// br-xxxx
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreBranchResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreBranchResponseBody) GetBranchId() *string {
	return s.BranchId
}

func (s *RestoreBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreBranchResponseBody) SetBranchId(v string) *RestoreBranchResponseBody {
	s.BranchId = &v
	return s
}

func (s *RestoreBranchResponseBody) SetRequestId(v string) *RestoreBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreBranchResponseBody) Validate() error {
	return dara.Validate(s)
}
