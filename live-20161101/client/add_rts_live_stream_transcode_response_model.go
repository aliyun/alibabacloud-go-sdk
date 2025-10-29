// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRtsLiveStreamTranscodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRtsLiveStreamTranscodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRtsLiveStreamTranscodeResponse
	GetStatusCode() *int32
	SetBody(v *AddRtsLiveStreamTranscodeResponseBody) *AddRtsLiveStreamTranscodeResponse
	GetBody() *AddRtsLiveStreamTranscodeResponseBody
}

type AddRtsLiveStreamTranscodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRtsLiveStreamTranscodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRtsLiveStreamTranscodeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRtsLiveStreamTranscodeResponse) GoString() string {
	return s.String()
}

func (s *AddRtsLiveStreamTranscodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRtsLiveStreamTranscodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRtsLiveStreamTranscodeResponse) GetBody() *AddRtsLiveStreamTranscodeResponseBody {
	return s.Body
}

func (s *AddRtsLiveStreamTranscodeResponse) SetHeaders(v map[string]*string) *AddRtsLiveStreamTranscodeResponse {
	s.Headers = v
	return s
}

func (s *AddRtsLiveStreamTranscodeResponse) SetStatusCode(v int32) *AddRtsLiveStreamTranscodeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeResponse) SetBody(v *AddRtsLiveStreamTranscodeResponseBody) *AddRtsLiveStreamTranscodeResponse {
	s.Body = v
	return s
}

func (s *AddRtsLiveStreamTranscodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
