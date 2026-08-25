// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomResourceStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomResourceStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomResourceStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomResourceStatsResponseBody) *GetCustomResourceStatsResponse
	GetBody() *GetCustomResourceStatsResponseBody
}

type GetCustomResourceStatsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomResourceStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomResourceStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomResourceStatsResponse) GoString() string {
	return s.String()
}

func (s *GetCustomResourceStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomResourceStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomResourceStatsResponse) GetBody() *GetCustomResourceStatsResponseBody {
	return s.Body
}

func (s *GetCustomResourceStatsResponse) SetHeaders(v map[string]*string) *GetCustomResourceStatsResponse {
	s.Headers = v
	return s
}

func (s *GetCustomResourceStatsResponse) SetStatusCode(v int32) *GetCustomResourceStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomResourceStatsResponse) SetBody(v *GetCustomResourceStatsResponseBody) *GetCustomResourceStatsResponse {
	s.Body = v
	return s
}

func (s *GetCustomResourceStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
