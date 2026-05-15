// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiOutboundTaskProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiOutboundTaskProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetAiOutboundTaskProgressResponseBody) *GetAiOutboundTaskProgressResponse
	GetBody() *GetAiOutboundTaskProgressResponseBody
}

type GetAiOutboundTaskProgressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiOutboundTaskProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiOutboundTaskProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskProgressResponse) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiOutboundTaskProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiOutboundTaskProgressResponse) GetBody() *GetAiOutboundTaskProgressResponseBody {
	return s.Body
}

func (s *GetAiOutboundTaskProgressResponse) SetHeaders(v map[string]*string) *GetAiOutboundTaskProgressResponse {
	s.Headers = v
	return s
}

func (s *GetAiOutboundTaskProgressResponse) SetStatusCode(v int32) *GetAiOutboundTaskProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponse) SetBody(v *GetAiOutboundTaskProgressResponseBody) *GetAiOutboundTaskProgressResponse {
	s.Body = v
	return s
}

func (s *GetAiOutboundTaskProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
