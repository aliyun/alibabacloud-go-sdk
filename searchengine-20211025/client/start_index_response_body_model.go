// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartIndexResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *StartIndexResponseBody
	GetResult() map[string]interface{}
}

type StartIndexResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D39EE0F1-D7EF-5F46-B781-6BF4185308B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The index map.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StartIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartIndexResponseBody) GoString() string {
	return s.String()
}

func (s *StartIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartIndexResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *StartIndexResponseBody) SetRequestId(v string) *StartIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartIndexResponseBody) SetResult(v map[string]interface{}) *StartIndexResponseBody {
	s.Result = v
	return s
}

func (s *StartIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
