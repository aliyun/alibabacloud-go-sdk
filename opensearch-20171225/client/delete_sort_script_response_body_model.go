// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSortScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSortScriptResponseBody
	GetRequestId() *string
}

type DeleteSortScriptResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSortScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSortScriptResponseBody) SetRequestId(v string) *DeleteSortScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSortScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
