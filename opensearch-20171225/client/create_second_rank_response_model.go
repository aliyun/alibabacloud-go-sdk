// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecondRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecondRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecondRankResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecondRankResponseBody) *CreateSecondRankResponse
	GetBody() *CreateSecondRankResponseBody
}

type CreateSecondRankResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecondRankResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecondRankResponse) GoString() string {
	return s.String()
}

func (s *CreateSecondRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecondRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecondRankResponse) GetBody() *CreateSecondRankResponseBody {
	return s.Body
}

func (s *CreateSecondRankResponse) SetHeaders(v map[string]*string) *CreateSecondRankResponse {
	s.Headers = v
	return s
}

func (s *CreateSecondRankResponse) SetStatusCode(v int32) *CreateSecondRankResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecondRankResponse) SetBody(v *CreateSecondRankResponseBody) *CreateSecondRankResponse {
	s.Body = v
	return s
}

func (s *CreateSecondRankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
