// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaCoverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMediaCoverResponseBody
	GetRequestId() *string
}

type UpdateMediaCoverResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0DC39B9E-13D4-40BA-AE76-CFF9BD64239D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaCoverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaCoverResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaCoverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaCoverResponseBody) SetRequestId(v string) *UpdateMediaCoverResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaCoverResponseBody) Validate() error {
	return dara.Validate(s)
}
