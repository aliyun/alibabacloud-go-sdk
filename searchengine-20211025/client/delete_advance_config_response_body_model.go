// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdvanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAdvanceConfigResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeleteAdvanceConfigResponseBody
	GetResult() map[string]interface{}
}

type DeleteAdvanceConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteAdvanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdvanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAdvanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAdvanceConfigResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeleteAdvanceConfigResponseBody) SetRequestId(v string) *DeleteAdvanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAdvanceConfigResponseBody) SetResult(v map[string]interface{}) *DeleteAdvanceConfigResponseBody {
	s.Result = v
	return s
}

func (s *DeleteAdvanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
