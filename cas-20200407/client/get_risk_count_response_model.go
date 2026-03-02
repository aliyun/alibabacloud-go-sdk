// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRiskCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRiskCountResponse
	GetStatusCode() *int32
	SetBody(v *GetRiskCountResponseBody) *GetRiskCountResponse
	GetBody() *GetRiskCountResponseBody
}

type GetRiskCountResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRiskCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRiskCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRiskCountResponse) GoString() string {
	return s.String()
}

func (s *GetRiskCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRiskCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRiskCountResponse) GetBody() *GetRiskCountResponseBody {
	return s.Body
}

func (s *GetRiskCountResponse) SetHeaders(v map[string]*string) *GetRiskCountResponse {
	s.Headers = v
	return s
}

func (s *GetRiskCountResponse) SetStatusCode(v int32) *GetRiskCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRiskCountResponse) SetBody(v *GetRiskCountResponseBody) *GetRiskCountResponse {
	s.Body = v
	return s
}

func (s *GetRiskCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
