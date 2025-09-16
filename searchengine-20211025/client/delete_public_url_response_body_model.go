// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePublicUrlResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeletePublicUrlResponseBody
	GetResult() map[string]interface{}
}

type DeletePublicUrlResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F43E8AB4-419C-5F4C-90D6-615590DFAA3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeletePublicUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePublicUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePublicUrlResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeletePublicUrlResponseBody) SetRequestId(v string) *DeletePublicUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePublicUrlResponseBody) SetResult(v map[string]interface{}) *DeletePublicUrlResponseBody {
	s.Result = v
	return s
}

func (s *DeletePublicUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
