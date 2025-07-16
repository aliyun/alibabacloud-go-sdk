// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScenegroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpenConversationId(v string) *CreateScenegroupResponseBody
	GetOpenConversationId() *string
	SetRequestId(v string) *CreateScenegroupResponseBody
	GetRequestId() *string
}

type CreateScenegroupResponseBody struct {
	// example:
	//
	// cid1324wwwerxxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateScenegroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScenegroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScenegroupResponseBody) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *CreateScenegroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScenegroupResponseBody) SetOpenConversationId(v string) *CreateScenegroupResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *CreateScenegroupResponseBody) SetRequestId(v string) *CreateScenegroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScenegroupResponseBody) Validate() error {
	return dara.Validate(s)
}
