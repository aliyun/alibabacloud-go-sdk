// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChannelsResponse
	GetStatusCode() *int32
	SetBody(v *ListChannelsResponseBody) *ListChannelsResponse
	GetBody() *ListChannelsResponseBody
}

type ListChannelsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChannelsResponse) GetBody() *ListChannelsResponseBody {
	return s.Body
}

func (s *ListChannelsResponse) SetHeaders(v map[string]*string) *ListChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListChannelsResponse) SetStatusCode(v int32) *ListChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChannelsResponse) SetBody(v *ListChannelsResponseBody) *ListChannelsResponse {
	s.Body = v
	return s
}

func (s *ListChannelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
