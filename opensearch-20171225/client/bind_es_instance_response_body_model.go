// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindEsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindEsInstanceResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *BindEsInstanceResponseBody
	GetResult() map[string]interface{}
}

type BindEsInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F5099063-6B86-F398-D843-905F9EFB683A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	//
	// example:
	//
	// []
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BindEsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindEsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *BindEsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindEsInstanceResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *BindEsInstanceResponseBody) SetRequestId(v string) *BindEsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindEsInstanceResponseBody) SetResult(v map[string]interface{}) *BindEsInstanceResponseBody {
	s.Result = v
	return s
}

func (s *BindEsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
