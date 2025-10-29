// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAICPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAICPublicKeyResponseBody
	GetRequestId() *string
}

type DeleteAICPublicKeyResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// AAE90880-4970-4D81-A534-A6C0F3631F74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAICPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAICPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAICPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAICPublicKeyResponseBody) SetRequestId(v string) *DeleteAICPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAICPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
