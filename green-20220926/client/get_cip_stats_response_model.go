// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCipStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCipStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCipStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetCipStatsResponseBody) *GetCipStatsResponse
	GetBody() *GetCipStatsResponseBody
}

type GetCipStatsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCipStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCipStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsResponse) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCipStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCipStatsResponse) GetBody() *GetCipStatsResponseBody {
	return s.Body
}

func (s *GetCipStatsResponse) SetHeaders(v map[string]*string) *GetCipStatsResponse {
	s.Headers = v
	return s
}

func (s *GetCipStatsResponse) SetStatusCode(v int32) *GetCipStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCipStatsResponse) SetBody(v *GetCipStatsResponseBody) *GetCipStatsResponse {
	s.Body = v
	return s
}

func (s *GetCipStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
