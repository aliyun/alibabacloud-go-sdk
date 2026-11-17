// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelApiResponse
	GetStatusCode() *int32
	SetBody(v *GetModelApiResponseBody) *GetModelApiResponse
	GetBody() *GetModelApiResponseBody
}

type GetModelApiResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelApiResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelApiResponse) GoString() string {
	return s.String()
}

func (s *GetModelApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelApiResponse) GetBody() *GetModelApiResponseBody {
	return s.Body
}

func (s *GetModelApiResponse) SetHeaders(v map[string]*string) *GetModelApiResponse {
	s.Headers = v
	return s
}

func (s *GetModelApiResponse) SetStatusCode(v int32) *GetModelApiResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelApiResponse) SetBody(v *GetModelApiResponseBody) *GetModelApiResponse {
	s.Body = v
	return s
}

func (s *GetModelApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
