// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BuildIndexResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *BuildIndexResponseBody
	GetResult() map[string]interface{}
}

type BuildIndexResponseBody struct {
	// id of request
	//
	// example:
	//
	// 407BFD91-DE7D-50BA-8F88-CDE52A3B5E46
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of clusters
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BuildIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BuildIndexResponseBody) GoString() string {
	return s.String()
}

func (s *BuildIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BuildIndexResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *BuildIndexResponseBody) SetRequestId(v string) *BuildIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuildIndexResponseBody) SetResult(v map[string]interface{}) *BuildIndexResponseBody {
	s.Result = v
	return s
}

func (s *BuildIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
