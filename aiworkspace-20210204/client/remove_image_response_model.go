// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveImageResponse
	GetStatusCode() *int32
	SetBody(v *RemoveImageResponseBody) *RemoveImageResponse
	GetBody() *RemoveImageResponseBody
}

type RemoveImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveImageResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageResponse) GoString() string {
	return s.String()
}

func (s *RemoveImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveImageResponse) GetBody() *RemoveImageResponseBody {
	return s.Body
}

func (s *RemoveImageResponse) SetHeaders(v map[string]*string) *RemoveImageResponse {
	s.Headers = v
	return s
}

func (s *RemoveImageResponse) SetStatusCode(v int32) *RemoveImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveImageResponse) SetBody(v *RemoveImageResponseBody) *RemoveImageResponse {
	s.Body = v
	return s
}

func (s *RemoveImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
