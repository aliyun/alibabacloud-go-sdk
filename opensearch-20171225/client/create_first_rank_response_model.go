// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirstRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFirstRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFirstRankResponse
	GetStatusCode() *int32
	SetBody(v *CreateFirstRankResponseBody) *CreateFirstRankResponse
	GetBody() *CreateFirstRankResponseBody
}

type CreateFirstRankResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFirstRankResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFirstRankResponse) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFirstRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFirstRankResponse) GetBody() *CreateFirstRankResponseBody {
	return s.Body
}

func (s *CreateFirstRankResponse) SetHeaders(v map[string]*string) *CreateFirstRankResponse {
	s.Headers = v
	return s
}

func (s *CreateFirstRankResponse) SetStatusCode(v int32) *CreateFirstRankResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFirstRankResponse) SetBody(v *CreateFirstRankResponseBody) *CreateFirstRankResponse {
	s.Body = v
	return s
}

func (s *CreateFirstRankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
