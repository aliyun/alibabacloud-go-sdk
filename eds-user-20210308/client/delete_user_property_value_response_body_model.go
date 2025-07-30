// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPropertyValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserPropertyValueResponseBody
	GetRequestId() *string
}

type DeleteUserPropertyValueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D6C62E40-F937-5803-B008-92E813399BA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserPropertyValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserPropertyValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserPropertyValueResponseBody) SetRequestId(v string) *DeleteUserPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserPropertyValueResponseBody) Validate() error {
	return dara.Validate(s)
}
