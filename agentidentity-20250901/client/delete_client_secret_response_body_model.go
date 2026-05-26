// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteClientSecretResponseBody
	GetRequestId() *string
}

type DeleteClientSecretResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClientSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClientSecretResponseBody) SetRequestId(v string) *DeleteClientSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClientSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
