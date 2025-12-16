// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABTestGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteABTestGroupResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeleteABTestGroupResponseBody
	GetResult() map[string]interface{}
}

type DeleteABTestGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteABTestGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABTestGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteABTestGroupResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeleteABTestGroupResponseBody) SetRequestId(v string) *DeleteABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteABTestGroupResponseBody) SetResult(v map[string]interface{}) *DeleteABTestGroupResponseBody {
	s.Result = v
	return s
}

func (s *DeleteABTestGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
