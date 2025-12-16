// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompileSortScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CompileSortScriptResponseBody
	GetRequestId() *string
}

type CompileSortScriptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CompileSortScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompileSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *CompileSortScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompileSortScriptResponseBody) SetRequestId(v string) *CompileSortScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompileSortScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
