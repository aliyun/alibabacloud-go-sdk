// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetBranchResponseBody
	GetRequestId() *string
}

type ResetBranchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetBranchResponseBody) GoString() string {
	return s.String()
}

func (s *ResetBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetBranchResponseBody) SetRequestId(v string) *ResetBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetBranchResponseBody) Validate() error {
	return dara.Validate(s)
}
