// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSecretValueResponseBody
	GetRequestId() *string
}

type DeleteSecretValueResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSecretValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecretValueResponseBody) SetRequestId(v string) *DeleteSecretValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretValueResponseBody) Validate() error {
	return dara.Validate(s)
}
