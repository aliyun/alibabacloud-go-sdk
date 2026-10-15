// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRewriteChannelTitleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RewriteChannelTitleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RewriteChannelTitleResponse
	GetStatusCode() *int32
	SetBody(v *RewriteChannelTitleResponseBody) *RewriteChannelTitleResponse
	GetBody() *RewriteChannelTitleResponseBody
}

type RewriteChannelTitleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RewriteChannelTitleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RewriteChannelTitleResponse) String() string {
	return dara.Prettify(s)
}

func (s RewriteChannelTitleResponse) GoString() string {
	return s.String()
}

func (s *RewriteChannelTitleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RewriteChannelTitleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RewriteChannelTitleResponse) GetBody() *RewriteChannelTitleResponseBody {
	return s.Body
}

func (s *RewriteChannelTitleResponse) SetHeaders(v map[string]*string) *RewriteChannelTitleResponse {
	s.Headers = v
	return s
}

func (s *RewriteChannelTitleResponse) SetStatusCode(v int32) *RewriteChannelTitleResponse {
	s.StatusCode = &v
	return s
}

func (s *RewriteChannelTitleResponse) SetBody(v *RewriteChannelTitleResponseBody) *RewriteChannelTitleResponse {
	s.Body = v
	return s
}

func (s *RewriteChannelTitleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
