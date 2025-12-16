// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSearchStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSearchStrategyResponse
	GetStatusCode() *int32
	SetBody(v *CreateSearchStrategyResponseBody) *CreateSearchStrategyResponse
	GetBody() *CreateSearchStrategyResponseBody
}

type CreateSearchStrategyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSearchStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateSearchStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSearchStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSearchStrategyResponse) GetBody() *CreateSearchStrategyResponseBody {
	return s.Body
}

func (s *CreateSearchStrategyResponse) SetHeaders(v map[string]*string) *CreateSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateSearchStrategyResponse) SetStatusCode(v int32) *CreateSearchStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSearchStrategyResponse) SetBody(v *CreateSearchStrategyResponseBody) *CreateSearchStrategyResponse {
	s.Body = v
	return s
}

func (s *CreateSearchStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
