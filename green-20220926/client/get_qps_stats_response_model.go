// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQpsStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQpsStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQpsStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetQpsStatsResponseBody) *GetQpsStatsResponse
	GetBody() *GetQpsStatsResponseBody
}

type GetQpsStatsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQpsStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQpsStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQpsStatsResponse) GoString() string {
	return s.String()
}

func (s *GetQpsStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQpsStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQpsStatsResponse) GetBody() *GetQpsStatsResponseBody {
	return s.Body
}

func (s *GetQpsStatsResponse) SetHeaders(v map[string]*string) *GetQpsStatsResponse {
	s.Headers = v
	return s
}

func (s *GetQpsStatsResponse) SetStatusCode(v int32) *GetQpsStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQpsStatsResponse) SetBody(v *GetQpsStatsResponseBody) *GetQpsStatsResponse {
	s.Body = v
	return s
}

func (s *GetQpsStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
