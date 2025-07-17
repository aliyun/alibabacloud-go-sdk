// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackInstanceRefreshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackInstanceRefreshResponseBody
	GetRequestId() *string
}

type RollbackInstanceRefreshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B13527BF-1FBD-4334-A512-20F5E9D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackInstanceRefreshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackInstanceRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackInstanceRefreshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackInstanceRefreshResponseBody) SetRequestId(v string) *RollbackInstanceRefreshResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackInstanceRefreshResponseBody) Validate() error {
	return dara.Validate(s)
}
