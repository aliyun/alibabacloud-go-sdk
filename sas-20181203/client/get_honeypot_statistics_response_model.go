// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHoneypotStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHoneypotStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetHoneypotStatisticsResponseBody) *GetHoneypotStatisticsResponse
	GetBody() *GetHoneypotStatisticsResponseBody
}

type GetHoneypotStatisticsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHoneypotStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHoneypotStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetHoneypotStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHoneypotStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHoneypotStatisticsResponse) GetBody() *GetHoneypotStatisticsResponseBody {
	return s.Body
}

func (s *GetHoneypotStatisticsResponse) SetHeaders(v map[string]*string) *GetHoneypotStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetHoneypotStatisticsResponse) SetStatusCode(v int32) *GetHoneypotStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHoneypotStatisticsResponse) SetBody(v *GetHoneypotStatisticsResponseBody) *GetHoneypotStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetHoneypotStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
