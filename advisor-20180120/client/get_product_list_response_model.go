// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProductListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProductListResponse
	GetStatusCode() *int32
	SetBody(v *GetProductListResponseBody) *GetProductListResponse
	GetBody() *GetProductListResponseBody
}

type GetProductListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProductListResponse) GoString() string {
	return s.String()
}

func (s *GetProductListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProductListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProductListResponse) GetBody() *GetProductListResponseBody {
	return s.Body
}

func (s *GetProductListResponse) SetHeaders(v map[string]*string) *GetProductListResponse {
	s.Headers = v
	return s
}

func (s *GetProductListResponse) SetStatusCode(v int32) *GetProductListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductListResponse) SetBody(v *GetProductListResponseBody) *GetProductListResponse {
	s.Body = v
	return s
}

func (s *GetProductListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
