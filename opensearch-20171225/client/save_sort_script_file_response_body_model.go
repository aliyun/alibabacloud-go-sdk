// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSortScriptFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSortScriptFileResponseBody
	GetRequestId() *string
}

type SaveSortScriptFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SaveSortScriptFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSortScriptFileResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSortScriptFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSortScriptFileResponseBody) SetRequestId(v string) *SaveSortScriptFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSortScriptFileResponseBody) Validate() error {
	return dara.Validate(s)
}
