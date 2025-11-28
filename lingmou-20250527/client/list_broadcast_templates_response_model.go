// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBroadcastTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBroadcastTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBroadcastTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListBroadcastTemplatesResponseBody) *ListBroadcastTemplatesResponse
	GetBody() *ListBroadcastTemplatesResponseBody
}

type ListBroadcastTemplatesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBroadcastTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBroadcastTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBroadcastTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListBroadcastTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBroadcastTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBroadcastTemplatesResponse) GetBody() *ListBroadcastTemplatesResponseBody {
	return s.Body
}

func (s *ListBroadcastTemplatesResponse) SetHeaders(v map[string]*string) *ListBroadcastTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListBroadcastTemplatesResponse) SetStatusCode(v int32) *ListBroadcastTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBroadcastTemplatesResponse) SetBody(v *ListBroadcastTemplatesResponseBody) *ListBroadcastTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListBroadcastTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
