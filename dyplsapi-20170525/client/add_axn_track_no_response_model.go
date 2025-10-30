// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAxnTrackNoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAxnTrackNoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAxnTrackNoResponse
	GetStatusCode() *int32
	SetBody(v *AddAxnTrackNoResponseBody) *AddAxnTrackNoResponse
	GetBody() *AddAxnTrackNoResponseBody
}

type AddAxnTrackNoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAxnTrackNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAxnTrackNoResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAxnTrackNoResponse) GoString() string {
	return s.String()
}

func (s *AddAxnTrackNoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAxnTrackNoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAxnTrackNoResponse) GetBody() *AddAxnTrackNoResponseBody {
	return s.Body
}

func (s *AddAxnTrackNoResponse) SetHeaders(v map[string]*string) *AddAxnTrackNoResponse {
	s.Headers = v
	return s
}

func (s *AddAxnTrackNoResponse) SetStatusCode(v int32) *AddAxnTrackNoResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAxnTrackNoResponse) SetBody(v *AddAxnTrackNoResponseBody) *AddAxnTrackNoResponse {
	s.Body = v
	return s
}

func (s *AddAxnTrackNoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
