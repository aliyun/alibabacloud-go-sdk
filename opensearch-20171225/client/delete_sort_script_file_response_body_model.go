// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSortScriptFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSortScriptFileResponseBody
	GetRequestId() *string
}

type DeleteSortScriptFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSortScriptFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSortScriptFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSortScriptFileResponseBody) SetRequestId(v string) *DeleteSortScriptFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSortScriptFileResponseBody) Validate() error {
	return dara.Validate(s)
}
