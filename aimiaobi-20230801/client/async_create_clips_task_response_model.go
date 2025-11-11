// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsyncCreateClipsTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsyncCreateClipsTaskResponse
	GetStatusCode() *int32
	SetBody(v *AsyncCreateClipsTaskResponseBody) *AsyncCreateClipsTaskResponse
	GetBody() *AsyncCreateClipsTaskResponseBody
}

type AsyncCreateClipsTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncCreateClipsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsyncCreateClipsTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTaskResponse) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsyncCreateClipsTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsyncCreateClipsTaskResponse) GetBody() *AsyncCreateClipsTaskResponseBody {
	return s.Body
}

func (s *AsyncCreateClipsTaskResponse) SetHeaders(v map[string]*string) *AsyncCreateClipsTaskResponse {
	s.Headers = v
	return s
}

func (s *AsyncCreateClipsTaskResponse) SetStatusCode(v int32) *AsyncCreateClipsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AsyncCreateClipsTaskResponse) SetBody(v *AsyncCreateClipsTaskResponseBody) *AsyncCreateClipsTaskResponse {
	s.Body = v
	return s
}

func (s *AsyncCreateClipsTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
