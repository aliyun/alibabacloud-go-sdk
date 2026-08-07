// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGuardLogStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGuardLogStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGuardLogStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetGuardLogStatsResponseBody) *GetGuardLogStatsResponse
	GetBody() *GetGuardLogStatsResponseBody
}

type GetGuardLogStatsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGuardLogStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGuardLogStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGuardLogStatsResponse) GoString() string {
	return s.String()
}

func (s *GetGuardLogStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGuardLogStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGuardLogStatsResponse) GetBody() *GetGuardLogStatsResponseBody {
	return s.Body
}

func (s *GetGuardLogStatsResponse) SetHeaders(v map[string]*string) *GetGuardLogStatsResponse {
	s.Headers = v
	return s
}

func (s *GetGuardLogStatsResponse) SetStatusCode(v int32) *GetGuardLogStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGuardLogStatsResponse) SetBody(v *GetGuardLogStatsResponseBody) *GetGuardLogStatsResponse {
	s.Body = v
	return s
}

func (s *GetGuardLogStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
