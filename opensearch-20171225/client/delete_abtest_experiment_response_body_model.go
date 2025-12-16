// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABTestExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteABTestExperimentResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeleteABTestExperimentResponseBody
	GetResult() map[string]interface{}
}

type DeleteABTestExperimentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteABTestExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABTestExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteABTestExperimentResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeleteABTestExperimentResponseBody) SetRequestId(v string) *DeleteABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteABTestExperimentResponseBody) SetResult(v map[string]interface{}) *DeleteABTestExperimentResponseBody {
	s.Result = v
	return s
}

func (s *DeleteABTestExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
