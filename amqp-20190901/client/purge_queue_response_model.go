// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PurgeQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PurgeQueueResponse
	GetStatusCode() *int32
	SetBody(v *PurgeQueueResponseBody) *PurgeQueueResponse
	GetBody() *PurgeQueueResponseBody
}

type PurgeQueueResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurgeQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurgeQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s PurgeQueueResponse) GoString() string {
	return s.String()
}

func (s *PurgeQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurgeQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PurgeQueueResponse) GetBody() *PurgeQueueResponseBody {
	return s.Body
}

func (s *PurgeQueueResponse) SetHeaders(v map[string]*string) *PurgeQueueResponse {
	s.Headers = v
	return s
}

func (s *PurgeQueueResponse) SetStatusCode(v int32) *PurgeQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *PurgeQueueResponse) SetBody(v *PurgeQueueResponseBody) *PurgeQueueResponse {
	s.Body = v
	return s
}

func (s *PurgeQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
