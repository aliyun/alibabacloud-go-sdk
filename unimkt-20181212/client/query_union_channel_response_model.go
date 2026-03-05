// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUnionChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUnionChannelResponse
	GetStatusCode() *int32
	SetBody(v *QueryUnionChannelResponseBody) *QueryUnionChannelResponse
	GetBody() *QueryUnionChannelResponseBody
}

type QueryUnionChannelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUnionChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUnionChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionChannelResponse) GoString() string {
	return s.String()
}

func (s *QueryUnionChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUnionChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUnionChannelResponse) GetBody() *QueryUnionChannelResponseBody {
	return s.Body
}

func (s *QueryUnionChannelResponse) SetHeaders(v map[string]*string) *QueryUnionChannelResponse {
	s.Headers = v
	return s
}

func (s *QueryUnionChannelResponse) SetStatusCode(v int32) *QueryUnionChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUnionChannelResponse) SetBody(v *QueryUnionChannelResponseBody) *QueryUnionChannelResponse {
	s.Body = v
	return s
}

func (s *QueryUnionChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
