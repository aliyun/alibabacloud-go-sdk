// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiModelCardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAiModelCardResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteAiModelCardResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAiModelCardResponseBody
	GetRequestId() *string
}

type DeleteAiModelCardResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAiModelCardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiModelCardResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAiModelCardResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAiModelCardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAiModelCardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAiModelCardResponseBody) SetCode(v string) *DeleteAiModelCardResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAiModelCardResponseBody) SetMessage(v string) *DeleteAiModelCardResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAiModelCardResponseBody) SetRequestId(v string) *DeleteAiModelCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAiModelCardResponseBody) Validate() error {
	return dara.Validate(s)
}
