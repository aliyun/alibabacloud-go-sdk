// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishValidateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FinishValidateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FinishValidateResponse
	GetStatusCode() *int32
	SetBody(v *FinishValidateResponseBody) *FinishValidateResponse
	GetBody() *FinishValidateResponseBody
}

type FinishValidateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishValidateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishValidateResponse) String() string {
	return dara.Prettify(s)
}

func (s FinishValidateResponse) GoString() string {
	return s.String()
}

func (s *FinishValidateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FinishValidateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FinishValidateResponse) GetBody() *FinishValidateResponseBody {
	return s.Body
}

func (s *FinishValidateResponse) SetHeaders(v map[string]*string) *FinishValidateResponse {
	s.Headers = v
	return s
}

func (s *FinishValidateResponse) SetStatusCode(v int32) *FinishValidateResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishValidateResponse) SetBody(v *FinishValidateResponseBody) *FinishValidateResponse {
	s.Body = v
	return s
}

func (s *FinishValidateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
