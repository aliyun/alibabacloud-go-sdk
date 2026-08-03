// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInsightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableInsightResponseBody
	GetRequestId() *string
}

type DisableInsightResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4ABAEA6E-C740-5CE2-A003-643E5519****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableInsightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableInsightResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInsightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableInsightResponseBody) SetRequestId(v string) *DisableInsightResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableInsightResponseBody) Validate() error {
	return dara.Validate(s)
}
