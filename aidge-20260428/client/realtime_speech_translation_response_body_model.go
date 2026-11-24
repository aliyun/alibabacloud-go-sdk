// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRealtimeSpeechTranslationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RealtimeSpeechTranslationResponseBody
	GetRequestId() *string
}

type RealtimeSpeechTranslationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// uuid
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RealtimeSpeechTranslationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RealtimeSpeechTranslationResponseBody) GoString() string {
	return s.String()
}

func (s *RealtimeSpeechTranslationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RealtimeSpeechTranslationResponseBody) SetRequestId(v string) *RealtimeSpeechTranslationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RealtimeSpeechTranslationResponseBody) Validate() error {
	return dara.Validate(s)
}
