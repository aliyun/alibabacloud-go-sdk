// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDingTalkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDingTalkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDingTalkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDingTalkResponseBody) *DeleteDingTalkResponse
	GetBody() *DeleteDingTalkResponseBody
}

type DeleteDingTalkResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDingTalkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDingTalkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDingTalkResponse) GoString() string {
	return s.String()
}

func (s *DeleteDingTalkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDingTalkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDingTalkResponse) GetBody() *DeleteDingTalkResponseBody {
	return s.Body
}

func (s *DeleteDingTalkResponse) SetHeaders(v map[string]*string) *DeleteDingTalkResponse {
	s.Headers = v
	return s
}

func (s *DeleteDingTalkResponse) SetStatusCode(v int32) *DeleteDingTalkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDingTalkResponse) SetBody(v *DeleteDingTalkResponseBody) *DeleteDingTalkResponse {
	s.Body = v
	return s
}

func (s *DeleteDingTalkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
