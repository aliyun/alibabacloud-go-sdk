// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDictResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteDictResponseBody
	GetResult() *bool
}

type DeleteDictResponseBody struct {
	// example:
	//
	// 2BF6B57E-5AAD-5389-80CD-E200BBF91FF9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteDictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDictResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDictResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteDictResponseBody) SetRequestId(v string) *DeleteDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDictResponseBody) SetResult(v bool) *DeleteDictResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteDictResponseBody) Validate() error {
	return dara.Validate(s)
}
