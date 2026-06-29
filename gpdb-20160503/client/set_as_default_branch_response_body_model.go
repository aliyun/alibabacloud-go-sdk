// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAsDefaultBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetAsDefaultBranchResponseBody
	GetRequestId() *string
}

type SetAsDefaultBranchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAsDefaultBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAsDefaultBranchResponseBody) GoString() string {
	return s.String()
}

func (s *SetAsDefaultBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAsDefaultBranchResponseBody) SetRequestId(v string) *SetAsDefaultBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAsDefaultBranchResponseBody) Validate() error {
	return dara.Validate(s)
}
