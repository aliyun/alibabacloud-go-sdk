// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAliasResponseBody
	GetRequestId() *string
}

type DeleteAliasResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4c8ae23f-3a42-6791-a4ba-1faa77831c28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAliasResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAliasResponseBody) SetRequestId(v string) *DeleteAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
