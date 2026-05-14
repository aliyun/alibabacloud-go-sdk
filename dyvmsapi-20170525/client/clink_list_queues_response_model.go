// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListQueuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkListQueuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkListQueuesResponse
	GetStatusCode() *int32
	SetBody(v *ClinkListQueuesResponseBody) *ClinkListQueuesResponse
	GetBody() *ClinkListQueuesResponseBody
}

type ClinkListQueuesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkListQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkListQueuesResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkListQueuesResponse) GoString() string {
	return s.String()
}

func (s *ClinkListQueuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkListQueuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkListQueuesResponse) GetBody() *ClinkListQueuesResponseBody {
	return s.Body
}

func (s *ClinkListQueuesResponse) SetHeaders(v map[string]*string) *ClinkListQueuesResponse {
	s.Headers = v
	return s
}

func (s *ClinkListQueuesResponse) SetStatusCode(v int32) *ClinkListQueuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkListQueuesResponse) SetBody(v *ClinkListQueuesResponseBody) *ClinkListQueuesResponse {
	s.Body = v
	return s
}

func (s *ClinkListQueuesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
