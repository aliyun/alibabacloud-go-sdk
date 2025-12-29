// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSTBServiceProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListSTBServiceProvidersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSTBServiceProvidersResponseBody
	GetRequestId() *string
	SetResult(v map[string][]*string) *ListSTBServiceProvidersResponseBody
	GetResult() map[string][]*string
	SetStatusCode(v int32) *ListSTBServiceProvidersResponseBody
	GetStatusCode() *int32
}

type ListSTBServiceProvidersResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1036C376-7A37-5A73-BE8B-C6DB40107EC1
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string][]*string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListSTBServiceProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSTBServiceProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListSTBServiceProvidersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSTBServiceProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSTBServiceProvidersResponseBody) GetResult() map[string][]*string {
	return s.Result
}

func (s *ListSTBServiceProvidersResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSTBServiceProvidersResponseBody) SetMessage(v string) *ListSTBServiceProvidersResponseBody {
	s.Message = &v
	return s
}

func (s *ListSTBServiceProvidersResponseBody) SetRequestId(v string) *ListSTBServiceProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSTBServiceProvidersResponseBody) SetResult(v map[string][]*string) *ListSTBServiceProvidersResponseBody {
	s.Result = v
	return s
}

func (s *ListSTBServiceProvidersResponseBody) SetStatusCode(v int32) *ListSTBServiceProvidersResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ListSTBServiceProvidersResponseBody) Validate() error {
	return dara.Validate(s)
}
