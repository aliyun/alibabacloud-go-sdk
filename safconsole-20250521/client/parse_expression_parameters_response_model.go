// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseExpressionParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ParseExpressionParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ParseExpressionParametersResponse
	GetStatusCode() *int32
	SetBody(v *ParseExpressionParametersResponseBody) *ParseExpressionParametersResponse
	GetBody() *ParseExpressionParametersResponseBody
}

type ParseExpressionParametersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ParseExpressionParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ParseExpressionParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s ParseExpressionParametersResponse) GoString() string {
	return s.String()
}

func (s *ParseExpressionParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ParseExpressionParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ParseExpressionParametersResponse) GetBody() *ParseExpressionParametersResponseBody {
	return s.Body
}

func (s *ParseExpressionParametersResponse) SetHeaders(v map[string]*string) *ParseExpressionParametersResponse {
	s.Headers = v
	return s
}

func (s *ParseExpressionParametersResponse) SetStatusCode(v int32) *ParseExpressionParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ParseExpressionParametersResponse) SetBody(v *ParseExpressionParametersResponseBody) *ParseExpressionParametersResponse {
	s.Body = v
	return s
}

func (s *ParseExpressionParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
