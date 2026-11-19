// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStreamVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StreamVoiceResponseBody
	GetRequestId() *string
}

type StreamVoiceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A3677B53-C3AE-5F04-953E-AA332CABB87B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StreamVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StreamVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *StreamVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StreamVoiceResponseBody) SetRequestId(v string) *StreamVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StreamVoiceResponseBody) Validate() error {
	return dara.Validate(s)
}
