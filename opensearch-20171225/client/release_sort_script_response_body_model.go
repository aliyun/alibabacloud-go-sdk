// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSortScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseSortScriptResponseBody
	GetRequestId() *string
}

type ReleaseSortScriptResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ReleaseSortScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseSortScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseSortScriptResponseBody) SetRequestId(v string) *ReleaseSortScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseSortScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
