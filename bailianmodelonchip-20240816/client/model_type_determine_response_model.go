// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelTypeDetermineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelTypeDetermineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelTypeDetermineResponse
	GetStatusCode() *int32
	SetBody(v *ModelTypeDetermineResponseBody) *ModelTypeDetermineResponse
	GetBody() *ModelTypeDetermineResponseBody
}

type ModelTypeDetermineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelTypeDetermineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelTypeDetermineResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelTypeDetermineResponse) GoString() string {
	return s.String()
}

func (s *ModelTypeDetermineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelTypeDetermineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelTypeDetermineResponse) GetBody() *ModelTypeDetermineResponseBody {
	return s.Body
}

func (s *ModelTypeDetermineResponse) SetHeaders(v map[string]*string) *ModelTypeDetermineResponse {
	s.Headers = v
	return s
}

func (s *ModelTypeDetermineResponse) SetStatusCode(v int32) *ModelTypeDetermineResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelTypeDetermineResponse) SetBody(v *ModelTypeDetermineResponseBody) *ModelTypeDetermineResponse {
	s.Body = v
	return s
}

func (s *ModelTypeDetermineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
