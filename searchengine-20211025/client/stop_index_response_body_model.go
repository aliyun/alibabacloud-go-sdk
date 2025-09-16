// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopIndexResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *StopIndexResponseBody
	GetResult() map[string]interface{}
}

type StopIndexResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The index map.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StopIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopIndexResponseBody) GoString() string {
	return s.String()
}

func (s *StopIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopIndexResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *StopIndexResponseBody) SetRequestId(v string) *StopIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopIndexResponseBody) SetResult(v map[string]interface{}) *StopIndexResponseBody {
	s.Result = v
	return s
}

func (s *StopIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
