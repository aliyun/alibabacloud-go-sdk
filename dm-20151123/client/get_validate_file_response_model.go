// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidateFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetValidateFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetValidateFileResponse
	GetStatusCode() *int32
	SetBody(v *GetValidateFileResponseBody) *GetValidateFileResponse
	GetBody() *GetValidateFileResponseBody
}

type GetValidateFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetValidateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetValidateFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetValidateFileResponse) GoString() string {
	return s.String()
}

func (s *GetValidateFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetValidateFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetValidateFileResponse) GetBody() *GetValidateFileResponseBody {
	return s.Body
}

func (s *GetValidateFileResponse) SetHeaders(v map[string]*string) *GetValidateFileResponse {
	s.Headers = v
	return s
}

func (s *GetValidateFileResponse) SetStatusCode(v int32) *GetValidateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetValidateFileResponse) SetBody(v *GetValidateFileResponseBody) *GetValidateFileResponse {
	s.Body = v
	return s
}

func (s *GetValidateFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
