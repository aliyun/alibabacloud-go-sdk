// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFormDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFormDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFormDataResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFormDataResponseBody) *DeleteFormDataResponse
	GetBody() *DeleteFormDataResponseBody
}

type DeleteFormDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFormDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFormDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteFormDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFormDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFormDataResponse) GetBody() *DeleteFormDataResponseBody {
	return s.Body
}

func (s *DeleteFormDataResponse) SetHeaders(v map[string]*string) *DeleteFormDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteFormDataResponse) SetStatusCode(v int32) *DeleteFormDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFormDataResponse) SetBody(v *DeleteFormDataResponseBody) *DeleteFormDataResponse {
	s.Body = v
	return s
}

func (s *DeleteFormDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
