// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthConfigByResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthConfigByResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthConfigByResourceResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthConfigByResourceResponseBody) *ListAuthConfigByResourceResponse
	GetBody() *ListAuthConfigByResourceResponseBody
}

type ListAuthConfigByResourceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthConfigByResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthConfigByResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthConfigByResourceResponse) GoString() string {
	return s.String()
}

func (s *ListAuthConfigByResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthConfigByResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthConfigByResourceResponse) GetBody() *ListAuthConfigByResourceResponseBody {
	return s.Body
}

func (s *ListAuthConfigByResourceResponse) SetHeaders(v map[string]*string) *ListAuthConfigByResourceResponse {
	s.Headers = v
	return s
}

func (s *ListAuthConfigByResourceResponse) SetStatusCode(v int32) *ListAuthConfigByResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthConfigByResourceResponse) SetBody(v *ListAuthConfigByResourceResponseBody) *ListAuthConfigByResourceResponse {
	s.Body = v
	return s
}

func (s *ListAuthConfigByResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
