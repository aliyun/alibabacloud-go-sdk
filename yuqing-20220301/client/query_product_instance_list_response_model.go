// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProductInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryProductInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryProductInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *QueryProductInstanceListResponseBody) *QueryProductInstanceListResponse
	GetBody() *QueryProductInstanceListResponseBody
}

type QueryProductInstanceListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProductInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProductInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryProductInstanceListResponse) GoString() string {
	return s.String()
}

func (s *QueryProductInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryProductInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryProductInstanceListResponse) GetBody() *QueryProductInstanceListResponseBody {
	return s.Body
}

func (s *QueryProductInstanceListResponse) SetHeaders(v map[string]*string) *QueryProductInstanceListResponse {
	s.Headers = v
	return s
}

func (s *QueryProductInstanceListResponse) SetStatusCode(v int32) *QueryProductInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProductInstanceListResponse) SetBody(v *QueryProductInstanceListResponseBody) *QueryProductInstanceListResponse {
	s.Body = v
	return s
}

func (s *QueryProductInstanceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
