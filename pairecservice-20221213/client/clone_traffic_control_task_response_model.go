// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneTrafficControlTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneTrafficControlTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneTrafficControlTaskResponse
	GetStatusCode() *int32
	SetBody(v *CloneTrafficControlTaskResponseBody) *CloneTrafficControlTaskResponse
	GetBody() *CloneTrafficControlTaskResponseBody
}

type CloneTrafficControlTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneTrafficControlTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *CloneTrafficControlTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneTrafficControlTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneTrafficControlTaskResponse) GetBody() *CloneTrafficControlTaskResponseBody {
	return s.Body
}

func (s *CloneTrafficControlTaskResponse) SetHeaders(v map[string]*string) *CloneTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *CloneTrafficControlTaskResponse) SetStatusCode(v int32) *CloneTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneTrafficControlTaskResponse) SetBody(v *CloneTrafficControlTaskResponseBody) *CloneTrafficControlTaskResponse {
	s.Body = v
	return s
}

func (s *CloneTrafficControlTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
