// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserInRecycleBinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserInRecycleBinResponseBody
	GetRequestId() *string
}

type DeleteUserInRecycleBinResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserInRecycleBinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserInRecycleBinResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserInRecycleBinResponseBody) SetRequestId(v string) *DeleteUserInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserInRecycleBinResponseBody) Validate() error {
	return dara.Validate(s)
}
