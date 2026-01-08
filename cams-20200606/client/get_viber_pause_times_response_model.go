// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViberPauseTimesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetViberPauseTimesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetViberPauseTimesResponse
	GetStatusCode() *int32
	SetBody(v *GetViberPauseTimesResponseBody) *GetViberPauseTimesResponse
	GetBody() *GetViberPauseTimesResponseBody
}

type GetViberPauseTimesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetViberPauseTimesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetViberPauseTimesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetViberPauseTimesResponse) GoString() string {
	return s.String()
}

func (s *GetViberPauseTimesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetViberPauseTimesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetViberPauseTimesResponse) GetBody() *GetViberPauseTimesResponseBody {
	return s.Body
}

func (s *GetViberPauseTimesResponse) SetHeaders(v map[string]*string) *GetViberPauseTimesResponse {
	s.Headers = v
	return s
}

func (s *GetViberPauseTimesResponse) SetStatusCode(v int32) *GetViberPauseTimesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetViberPauseTimesResponse) SetBody(v *GetViberPauseTimesResponseBody) *GetViberPauseTimesResponse {
	s.Body = v
	return s
}

func (s *GetViberPauseTimesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
