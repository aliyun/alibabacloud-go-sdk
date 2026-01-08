// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFbInstagramPagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFbInstagramPagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFbInstagramPagesResponse
	GetStatusCode() *int32
	SetBody(v *GetFbInstagramPagesResponseBody) *GetFbInstagramPagesResponse
	GetBody() *GetFbInstagramPagesResponseBody
}

type GetFbInstagramPagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFbInstagramPagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFbInstagramPagesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFbInstagramPagesResponse) GoString() string {
	return s.String()
}

func (s *GetFbInstagramPagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFbInstagramPagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFbInstagramPagesResponse) GetBody() *GetFbInstagramPagesResponseBody {
	return s.Body
}

func (s *GetFbInstagramPagesResponse) SetHeaders(v map[string]*string) *GetFbInstagramPagesResponse {
	s.Headers = v
	return s
}

func (s *GetFbInstagramPagesResponse) SetStatusCode(v int32) *GetFbInstagramPagesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFbInstagramPagesResponse) SetBody(v *GetFbInstagramPagesResponseBody) *GetFbInstagramPagesResponse {
	s.Body = v
	return s
}

func (s *GetFbInstagramPagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
