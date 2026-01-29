// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProductListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryProductListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryProductListResponse
	GetStatusCode() *int32
	SetBody(v *QueryProductListResponseBody) *QueryProductListResponse
	GetBody() *QueryProductListResponseBody
}

type QueryProductListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProductListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProductListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryProductListResponse) GoString() string {
	return s.String()
}

func (s *QueryProductListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryProductListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryProductListResponse) GetBody() *QueryProductListResponseBody {
	return s.Body
}

func (s *QueryProductListResponse) SetHeaders(v map[string]*string) *QueryProductListResponse {
	s.Headers = v
	return s
}

func (s *QueryProductListResponse) SetStatusCode(v int32) *QueryProductListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProductListResponse) SetBody(v *QueryProductListResponseBody) *QueryProductListResponse {
	s.Body = v
	return s
}

func (s *QueryProductListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
