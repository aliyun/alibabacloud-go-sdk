// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserAuthConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserAuthConfigResponseBody
	GetRequestId() *string
}

type DeleteUserAuthConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6C38D3F9-B340-5230-B108-77E675452733
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserAuthConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserAuthConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserAuthConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserAuthConfigResponseBody) SetRequestId(v string) *DeleteUserAuthConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserAuthConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
