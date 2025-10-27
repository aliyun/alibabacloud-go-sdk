// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotAttackerPortraitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHoneypotAttackerPortraitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHoneypotAttackerPortraitResponse
	GetStatusCode() *int32
	SetBody(v *ListHoneypotAttackerPortraitResponseBody) *ListHoneypotAttackerPortraitResponse
	GetBody() *ListHoneypotAttackerPortraitResponseBody
}

type ListHoneypotAttackerPortraitResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHoneypotAttackerPortraitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHoneypotAttackerPortraitResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAttackerPortraitResponse) GoString() string {
	return s.String()
}

func (s *ListHoneypotAttackerPortraitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHoneypotAttackerPortraitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHoneypotAttackerPortraitResponse) GetBody() *ListHoneypotAttackerPortraitResponseBody {
	return s.Body
}

func (s *ListHoneypotAttackerPortraitResponse) SetHeaders(v map[string]*string) *ListHoneypotAttackerPortraitResponse {
	s.Headers = v
	return s
}

func (s *ListHoneypotAttackerPortraitResponse) SetStatusCode(v int32) *ListHoneypotAttackerPortraitResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHoneypotAttackerPortraitResponse) SetBody(v *ListHoneypotAttackerPortraitResponseBody) *ListHoneypotAttackerPortraitResponse {
	s.Body = v
	return s
}

func (s *ListHoneypotAttackerPortraitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
