// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ActiveFlowLogResponseBody
	GetRequestId() *string
}

type ActiveFlowLogResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CBBE5EBF-69C1-4395-B36B-26B7605F87EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActiveFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActiveFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActiveFlowLogResponseBody) SetRequestId(v string) *ActiveFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
