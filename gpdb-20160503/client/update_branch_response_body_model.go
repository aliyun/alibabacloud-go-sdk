// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateBranchResponseBody
	GetRequestId() *string
}

type UpdateBranchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBranchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBranchResponseBody) SetRequestId(v string) *UpdateBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBranchResponseBody) Validate() error {
	return dara.Validate(s)
}
