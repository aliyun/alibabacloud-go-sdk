// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotAttackStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHoneypotAttackStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHoneypotAttackStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetHoneypotAttackStatisticsResponseBody) *GetHoneypotAttackStatisticsResponse
	GetBody() *GetHoneypotAttackStatisticsResponseBody
}

type GetHoneypotAttackStatisticsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHoneypotAttackStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHoneypotAttackStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotAttackStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetHoneypotAttackStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHoneypotAttackStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHoneypotAttackStatisticsResponse) GetBody() *GetHoneypotAttackStatisticsResponseBody {
	return s.Body
}

func (s *GetHoneypotAttackStatisticsResponse) SetHeaders(v map[string]*string) *GetHoneypotAttackStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetHoneypotAttackStatisticsResponse) SetStatusCode(v int32) *GetHoneypotAttackStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHoneypotAttackStatisticsResponse) SetBody(v *GetHoneypotAttackStatisticsResponseBody) *GetHoneypotAttackStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetHoneypotAttackStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
