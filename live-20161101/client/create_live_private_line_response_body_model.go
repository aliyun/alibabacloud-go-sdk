// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePrivateLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLivePrivateLineResponseBody
	GetRequestId() *string
}

type CreateLivePrivateLineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7908F2FF-44F8-120F-9FD6-85AE4B6C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLivePrivateLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePrivateLineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLivePrivateLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLivePrivateLineResponseBody) SetRequestId(v string) *CreateLivePrivateLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLivePrivateLineResponseBody) Validate() error {
	return dara.Validate(s)
}
