// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStreamVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StreamVoiceRequest struct {
}

func (s StreamVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s StreamVoiceRequest) GoString() string {
	return s.String()
}

func (s *StreamVoiceRequest) Validate() error {
	return dara.Validate(s)
}
