// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppSecretResponseBody
	GetRequestId() *string
}

type DeleteAppSecretResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppSecretResponseBody) SetRequestId(v string) *DeleteAppSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
