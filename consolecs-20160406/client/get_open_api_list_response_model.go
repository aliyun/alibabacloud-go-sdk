// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenApiListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpenApiListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpenApiListResponse
	GetStatusCode() *int32
	SetBody(v *GetOpenApiListResponseBody) *GetOpenApiListResponse
	GetBody() *GetOpenApiListResponseBody
}

type GetOpenApiListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpenApiListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpenApiListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpenApiListResponse) GoString() string {
	return s.String()
}

func (s *GetOpenApiListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpenApiListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpenApiListResponse) GetBody() *GetOpenApiListResponseBody {
	return s.Body
}

func (s *GetOpenApiListResponse) SetHeaders(v map[string]*string) *GetOpenApiListResponse {
	s.Headers = v
	return s
}

func (s *GetOpenApiListResponse) SetStatusCode(v int32) *GetOpenApiListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenApiListResponse) SetBody(v *GetOpenApiListResponseBody) *GetOpenApiListResponse {
	s.Body = v
	return s
}

func (s *GetOpenApiListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
