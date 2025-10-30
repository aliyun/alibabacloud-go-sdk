// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneNoAByTrackNoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPhoneNoAByTrackNoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPhoneNoAByTrackNoResponse
	GetStatusCode() *int32
	SetBody(v *QueryPhoneNoAByTrackNoResponseBody) *QueryPhoneNoAByTrackNoResponse
	GetBody() *QueryPhoneNoAByTrackNoResponseBody
}

type QueryPhoneNoAByTrackNoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPhoneNoAByTrackNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPhoneNoAByTrackNoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoResponse) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPhoneNoAByTrackNoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPhoneNoAByTrackNoResponse) GetBody() *QueryPhoneNoAByTrackNoResponseBody {
	return s.Body
}

func (s *QueryPhoneNoAByTrackNoResponse) SetHeaders(v map[string]*string) *QueryPhoneNoAByTrackNoResponse {
	s.Headers = v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponse) SetStatusCode(v int32) *QueryPhoneNoAByTrackNoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponse) SetBody(v *QueryPhoneNoAByTrackNoResponseBody) *QueryPhoneNoAByTrackNoResponse {
	s.Body = v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
