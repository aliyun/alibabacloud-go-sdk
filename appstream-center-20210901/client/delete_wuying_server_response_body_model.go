// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWuyingServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWuyingServerResponseBody
	GetRequestId() *string
}

type DeleteWuyingServerResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWuyingServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWuyingServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWuyingServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWuyingServerResponseBody) SetRequestId(v string) *DeleteWuyingServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWuyingServerResponseBody) Validate() error {
	return dara.Validate(s)
}
