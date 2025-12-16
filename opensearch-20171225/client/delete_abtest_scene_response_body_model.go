// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABTestSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteABTestSceneResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeleteABTestSceneResponseBody
	GetResult() map[string]interface{}
}

type DeleteABTestSceneResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteABTestSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABTestSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteABTestSceneResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeleteABTestSceneResponseBody) SetRequestId(v string) *DeleteABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteABTestSceneResponseBody) SetResult(v map[string]interface{}) *DeleteABTestSceneResponseBody {
	s.Result = v
	return s
}

func (s *DeleteABTestSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
