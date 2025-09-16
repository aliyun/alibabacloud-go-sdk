// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopTaskResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *StopTaskResponseBody
	GetResult() map[string]interface{}
}

type StopTaskResponseBody struct {
	// id of request
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StopTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTaskResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *StopTaskResponseBody) SetRequestId(v string) *StopTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTaskResponseBody) SetResult(v map[string]interface{}) *StopTaskResponseBody {
	s.Result = v
	return s
}

func (s *StopTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
