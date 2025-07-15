// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversationDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConversationDetailResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetConversationDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetConversationDetailResponseBody
	GetMessage() *string
	SetPhrases(v []*GetConversationDetailResponseBodyPhrases) *GetConversationDetailResponseBody
	GetPhrases() []*GetConversationDetailResponseBodyPhrases
	SetRequestId(v string) *GetConversationDetailResponseBody
	GetRequestId() *string
}

type GetConversationDetailResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Phrases        []*GetConversationDetailResponseBodyPhrases `json:"Phrases,omitempty" xml:"Phrases,omitempty" type:"Repeated"`
	// example:
	//
	// 7E407F9B-A278-52A0-B193-3EE5471D7A87
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConversationDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConversationDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConversationDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetConversationDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConversationDetailResponseBody) GetPhrases() []*GetConversationDetailResponseBodyPhrases {
	return s.Phrases
}

func (s *GetConversationDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConversationDetailResponseBody) SetCode(v string) *GetConversationDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetConversationDetailResponseBody) SetHttpStatusCode(v int32) *GetConversationDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConversationDetailResponseBody) SetMessage(v string) *GetConversationDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetConversationDetailResponseBody) SetPhrases(v []*GetConversationDetailResponseBodyPhrases) *GetConversationDetailResponseBody {
	s.Phrases = v
	return s
}

func (s *GetConversationDetailResponseBody) SetRequestId(v string) *GetConversationDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConversationDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetConversationDetailResponseBodyPhrases struct {
	// example:
	//
	// 240
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 1280
	End      *int32 `json:"End,omitempty" xml:"End,omitempty"`
	Finished *bool  `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// example:
	//
	// agent@ccc-test
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Words    *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetConversationDetailResponseBodyPhrases) String() string {
	return dara.Prettify(s)
}

func (s GetConversationDetailResponseBodyPhrases) GoString() string {
	return s.String()
}

func (s *GetConversationDetailResponseBodyPhrases) GetBegin() *int32 {
	return s.Begin
}

func (s *GetConversationDetailResponseBodyPhrases) GetEnd() *int32 {
	return s.End
}

func (s *GetConversationDetailResponseBodyPhrases) GetFinished() *bool {
	return s.Finished
}

func (s *GetConversationDetailResponseBodyPhrases) GetIdentity() *string {
	return s.Identity
}

func (s *GetConversationDetailResponseBodyPhrases) GetRole() *string {
	return s.Role
}

func (s *GetConversationDetailResponseBodyPhrases) GetWords() *string {
	return s.Words
}

func (s *GetConversationDetailResponseBodyPhrases) SetBegin(v int32) *GetConversationDetailResponseBodyPhrases {
	s.Begin = &v
	return s
}

func (s *GetConversationDetailResponseBodyPhrases) SetEnd(v int32) *GetConversationDetailResponseBodyPhrases {
	s.End = &v
	return s
}

func (s *GetConversationDetailResponseBodyPhrases) SetFinished(v bool) *GetConversationDetailResponseBodyPhrases {
	s.Finished = &v
	return s
}

func (s *GetConversationDetailResponseBodyPhrases) SetIdentity(v string) *GetConversationDetailResponseBodyPhrases {
	s.Identity = &v
	return s
}

func (s *GetConversationDetailResponseBodyPhrases) SetRole(v string) *GetConversationDetailResponseBodyPhrases {
	s.Role = &v
	return s
}

func (s *GetConversationDetailResponseBodyPhrases) SetWords(v string) *GetConversationDetailResponseBodyPhrases {
	s.Words = &v
	return s
}

func (s *GetConversationDetailResponseBodyPhrases) Validate() error {
	return dara.Validate(s)
}
