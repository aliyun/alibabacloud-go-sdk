// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRayHistoryServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRayHistoryServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRayHistoryServerResponse
	GetStatusCode() *int32
	SetBody(v *StopRayHistoryServerResponseBody) *StopRayHistoryServerResponse
	GetBody() *StopRayHistoryServerResponseBody
}

type StopRayHistoryServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRayHistoryServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRayHistoryServerResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRayHistoryServerResponse) GoString() string {
	return s.String()
}

func (s *StopRayHistoryServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRayHistoryServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRayHistoryServerResponse) GetBody() *StopRayHistoryServerResponseBody {
	return s.Body
}

func (s *StopRayHistoryServerResponse) SetHeaders(v map[string]*string) *StopRayHistoryServerResponse {
	s.Headers = v
	return s
}

func (s *StopRayHistoryServerResponse) SetStatusCode(v int32) *StopRayHistoryServerResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRayHistoryServerResponse) SetBody(v *StopRayHistoryServerResponseBody) *StopRayHistoryServerResponse {
	s.Body = v
	return s
}

func (s *StopRayHistoryServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
