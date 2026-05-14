// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkCreateClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkCreateClientResponse
	GetStatusCode() *int32
	SetBody(v *ClinkCreateClientResponseBody) *ClinkCreateClientResponse
	GetBody() *ClinkCreateClientResponseBody
}

type ClinkCreateClientResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkCreateClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkCreateClientResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateClientResponse) GoString() string {
	return s.String()
}

func (s *ClinkCreateClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkCreateClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkCreateClientResponse) GetBody() *ClinkCreateClientResponseBody {
	return s.Body
}

func (s *ClinkCreateClientResponse) SetHeaders(v map[string]*string) *ClinkCreateClientResponse {
	s.Headers = v
	return s
}

func (s *ClinkCreateClientResponse) SetStatusCode(v int32) *ClinkCreateClientResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkCreateClientResponse) SetBody(v *ClinkCreateClientResponseBody) *ClinkCreateClientResponse {
	s.Body = v
	return s
}

func (s *ClinkCreateClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
