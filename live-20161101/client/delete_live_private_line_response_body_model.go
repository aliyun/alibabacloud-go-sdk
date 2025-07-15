// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePrivateLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLivePrivateLineResponseBody
	GetRequestId() *string
}

type DeleteLivePrivateLineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7908F2FF-44F8-120F-9FD6-85AE4******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLivePrivateLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePrivateLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLivePrivateLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLivePrivateLineResponseBody) SetRequestId(v string) *DeleteLivePrivateLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLivePrivateLineResponseBody) Validate() error {
	return dara.Validate(s)
}
