// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartWuyingServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartWuyingServerResponseBody
	GetRequestId() *string
}

type RestartWuyingServerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartWuyingServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartWuyingServerResponseBody) GoString() string {
	return s.String()
}

func (s *RestartWuyingServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartWuyingServerResponseBody) SetRequestId(v string) *RestartWuyingServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartWuyingServerResponseBody) Validate() error {
	return dara.Validate(s)
}
