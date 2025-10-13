// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlgorithmVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlgorithmVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetAlgorithmVersionResponseBody) *GetAlgorithmVersionResponse
	GetBody() *GetAlgorithmVersionResponseBody
}

type GetAlgorithmVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlgorithmVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlgorithmVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmVersionResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlgorithmVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlgorithmVersionResponse) GetBody() *GetAlgorithmVersionResponseBody {
	return s.Body
}

func (s *GetAlgorithmVersionResponse) SetHeaders(v map[string]*string) *GetAlgorithmVersionResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmVersionResponse) SetStatusCode(v int32) *GetAlgorithmVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgorithmVersionResponse) SetBody(v *GetAlgorithmVersionResponseBody) *GetAlgorithmVersionResponse {
	s.Body = v
	return s
}

func (s *GetAlgorithmVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
