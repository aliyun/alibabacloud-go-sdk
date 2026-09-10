// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExecuteSqlConversionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExecuteSqlConversionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExecuteSqlConversionResponse
	GetStatusCode() *int32
	SetBody(v *CreateExecuteSqlConversionResponseBody) *CreateExecuteSqlConversionResponse
	GetBody() *CreateExecuteSqlConversionResponseBody
}

type CreateExecuteSqlConversionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExecuteSqlConversionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExecuteSqlConversionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExecuteSqlConversionResponse) GoString() string {
	return s.String()
}

func (s *CreateExecuteSqlConversionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExecuteSqlConversionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExecuteSqlConversionResponse) GetBody() *CreateExecuteSqlConversionResponseBody {
	return s.Body
}

func (s *CreateExecuteSqlConversionResponse) SetHeaders(v map[string]*string) *CreateExecuteSqlConversionResponse {
	s.Headers = v
	return s
}

func (s *CreateExecuteSqlConversionResponse) SetStatusCode(v int32) *CreateExecuteSqlConversionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExecuteSqlConversionResponse) SetBody(v *CreateExecuteSqlConversionResponseBody) *CreateExecuteSqlConversionResponse {
	s.Body = v
	return s
}

func (s *CreateExecuteSqlConversionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
