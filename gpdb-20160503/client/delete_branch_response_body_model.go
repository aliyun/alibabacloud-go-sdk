// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBranchResponseBody
	GetRequestId() *string
}

type DeleteBranchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBranchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBranchResponseBody) SetRequestId(v string) *DeleteBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBranchResponseBody) Validate() error {
	return dara.Validate(s)
}
