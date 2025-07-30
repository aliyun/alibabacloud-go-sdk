// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationClientSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApplicationClientSecretResponseBody
	GetRequestId() *string
}

type DeleteApplicationClientSecretResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationClientSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationClientSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationClientSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationClientSecretResponseBody) SetRequestId(v string) *DeleteApplicationClientSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationClientSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
