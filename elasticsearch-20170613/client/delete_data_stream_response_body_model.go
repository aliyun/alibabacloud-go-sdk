// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataStreamResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteDataStreamResponseBody
	GetResult() *bool
}

type DeleteDataStreamResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteDataStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataStreamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataStreamResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteDataStreamResponseBody) SetRequestId(v string) *DeleteDataStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataStreamResponseBody) SetResult(v bool) *DeleteDataStreamResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteDataStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
