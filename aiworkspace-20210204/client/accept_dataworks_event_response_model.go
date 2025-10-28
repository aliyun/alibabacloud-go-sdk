// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptDataworksEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcceptDataworksEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcceptDataworksEventResponse
	GetStatusCode() *int32
	SetBody(v *AcceptDataworksEventResponseBody) *AcceptDataworksEventResponse
	GetBody() *AcceptDataworksEventResponseBody
}

type AcceptDataworksEventResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptDataworksEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptDataworksEventResponse) String() string {
	return dara.Prettify(s)
}

func (s AcceptDataworksEventResponse) GoString() string {
	return s.String()
}

func (s *AcceptDataworksEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcceptDataworksEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcceptDataworksEventResponse) GetBody() *AcceptDataworksEventResponseBody {
	return s.Body
}

func (s *AcceptDataworksEventResponse) SetHeaders(v map[string]*string) *AcceptDataworksEventResponse {
	s.Headers = v
	return s
}

func (s *AcceptDataworksEventResponse) SetStatusCode(v int32) *AcceptDataworksEventResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptDataworksEventResponse) SetBody(v *AcceptDataworksEventResponseBody) *AcceptDataworksEventResponse {
	s.Body = v
	return s
}

func (s *AcceptDataworksEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
