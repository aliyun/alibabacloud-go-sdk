// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHaVipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHaVipResponseBody
	GetRequestId() *string
}

type DeleteHaVipResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C44F62BE-9CE7-4277-B117-69243F3988BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHaVipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHaVipResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHaVipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHaVipResponseBody) SetRequestId(v string) *DeleteHaVipResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHaVipResponseBody) Validate() error {
	return dara.Validate(s)
}
