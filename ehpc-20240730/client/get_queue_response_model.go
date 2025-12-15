// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueueResponse
	GetStatusCode() *int32
	SetBody(v *GetQueueResponseBody) *GetQueueResponse
	GetBody() *GetQueueResponseBody
}

type GetQueueResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueueResponse) GoString() string {
	return s.String()
}

func (s *GetQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueueResponse) GetBody() *GetQueueResponseBody {
	return s.Body
}

func (s *GetQueueResponse) SetHeaders(v map[string]*string) *GetQueueResponse {
	s.Headers = v
	return s
}

func (s *GetQueueResponse) SetStatusCode(v int32) *GetQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueueResponse) SetBody(v *GetQueueResponseBody) *GetQueueResponse {
	s.Body = v
	return s
}

func (s *GetQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
