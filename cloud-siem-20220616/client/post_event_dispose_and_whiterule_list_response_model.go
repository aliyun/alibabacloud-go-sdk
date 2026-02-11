// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostEventDisposeAndWhiteruleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostEventDisposeAndWhiteruleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostEventDisposeAndWhiteruleListResponse
	GetStatusCode() *int32
	SetBody(v *PostEventDisposeAndWhiteruleListResponseBody) *PostEventDisposeAndWhiteruleListResponse
	GetBody() *PostEventDisposeAndWhiteruleListResponseBody
}

type PostEventDisposeAndWhiteruleListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostEventDisposeAndWhiteruleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostEventDisposeAndWhiteruleListResponse) String() string {
	return dara.Prettify(s)
}

func (s PostEventDisposeAndWhiteruleListResponse) GoString() string {
	return s.String()
}

func (s *PostEventDisposeAndWhiteruleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostEventDisposeAndWhiteruleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostEventDisposeAndWhiteruleListResponse) GetBody() *PostEventDisposeAndWhiteruleListResponseBody {
	return s.Body
}

func (s *PostEventDisposeAndWhiteruleListResponse) SetHeaders(v map[string]*string) *PostEventDisposeAndWhiteruleListResponse {
	s.Headers = v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponse) SetStatusCode(v int32) *PostEventDisposeAndWhiteruleListResponse {
	s.StatusCode = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponse) SetBody(v *PostEventDisposeAndWhiteruleListResponseBody) *PostEventDisposeAndWhiteruleListResponse {
	s.Body = v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
