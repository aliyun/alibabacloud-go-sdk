// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartCasterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartCasterResponseBody
	GetRequestId() *string
}

type RestartCasterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6EBD1AC4-C34D-4AE1-963E-B688A228BE31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartCasterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartCasterResponseBody) GoString() string {
	return s.String()
}

func (s *RestartCasterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartCasterResponseBody) SetRequestId(v string) *RestartCasterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartCasterResponseBody) Validate() error {
	return dara.Validate(s)
}
