// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAgentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *DeleteAgentResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *DeleteAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAgentResponseBody
	GetSuccess() *bool
}

type DeleteAgentResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAgentResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DeleteAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAgentResponseBody) SetCode(v string) *DeleteAgentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentResponseBody) SetHttpStatusCode(v int64) *DeleteAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAgentResponseBody) SetMessage(v string) *DeleteAgentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentResponseBody) SetRequestId(v string) *DeleteAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentResponseBody) SetSuccess(v bool) *DeleteAgentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
