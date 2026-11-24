// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRealtimeSpeechTranslationRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RealtimeSpeechTranslationRequest struct {
}

func (s RealtimeSpeechTranslationRequest) String() string {
	return dara.Prettify(s)
}

func (s RealtimeSpeechTranslationRequest) GoString() string {
	return s.String()
}

func (s *RealtimeSpeechTranslationRequest) Validate() error {
	return dara.Validate(s)
}
