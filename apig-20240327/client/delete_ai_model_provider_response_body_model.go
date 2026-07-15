// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiModelProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAiModelProviderResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteAiModelProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAiModelProviderResponseBody
	GetRequestId() *string
}

type DeleteAiModelProviderResponseBody struct {
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

func (s DeleteAiModelProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiModelProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAiModelProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAiModelProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAiModelProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAiModelProviderResponseBody) SetCode(v string) *DeleteAiModelProviderResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAiModelProviderResponseBody) SetMessage(v string) *DeleteAiModelProviderResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAiModelProviderResponseBody) SetRequestId(v string) *DeleteAiModelProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAiModelProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
