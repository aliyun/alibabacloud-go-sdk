// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyMaterialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyMaterialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyMaterialResponse
	GetStatusCode() *int32
	SetBody(v *VerifyMaterialResponseBody) *VerifyMaterialResponse
	GetBody() *VerifyMaterialResponseBody
}

type VerifyMaterialResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyMaterialResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyMaterialResponse) GoString() string {
	return s.String()
}

func (s *VerifyMaterialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyMaterialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyMaterialResponse) GetBody() *VerifyMaterialResponseBody {
	return s.Body
}

func (s *VerifyMaterialResponse) SetHeaders(v map[string]*string) *VerifyMaterialResponse {
	s.Headers = v
	return s
}

func (s *VerifyMaterialResponse) SetStatusCode(v int32) *VerifyMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyMaterialResponse) SetBody(v *VerifyMaterialResponseBody) *VerifyMaterialResponse {
	s.Body = v
	return s
}

func (s *VerifyMaterialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
