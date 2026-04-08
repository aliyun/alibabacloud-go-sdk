// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMmsTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMmsTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMmsTimerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMmsTimerResponseBody) *DeleteMmsTimerResponse
	GetBody() *DeleteMmsTimerResponseBody
}

type DeleteMmsTimerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMmsTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMmsTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMmsTimerResponse) GoString() string {
	return s.String()
}

func (s *DeleteMmsTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMmsTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMmsTimerResponse) GetBody() *DeleteMmsTimerResponseBody {
	return s.Body
}

func (s *DeleteMmsTimerResponse) SetHeaders(v map[string]*string) *DeleteMmsTimerResponse {
	s.Headers = v
	return s
}

func (s *DeleteMmsTimerResponse) SetStatusCode(v int32) *DeleteMmsTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMmsTimerResponse) SetBody(v *DeleteMmsTimerResponseBody) *DeleteMmsTimerResponse {
	s.Body = v
	return s
}

func (s *DeleteMmsTimerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
