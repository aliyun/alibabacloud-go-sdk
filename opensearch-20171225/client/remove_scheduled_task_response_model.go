// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveScheduledTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveScheduledTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveScheduledTaskResponse
	GetStatusCode() *int32
	SetBody(v *RemoveScheduledTaskResponseBody) *RemoveScheduledTaskResponse
	GetBody() *RemoveScheduledTaskResponseBody
}

type RemoveScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveScheduledTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *RemoveScheduledTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveScheduledTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveScheduledTaskResponse) GetBody() *RemoveScheduledTaskResponseBody {
	return s.Body
}

func (s *RemoveScheduledTaskResponse) SetHeaders(v map[string]*string) *RemoveScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *RemoveScheduledTaskResponse) SetStatusCode(v int32) *RemoveScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveScheduledTaskResponse) SetBody(v *RemoveScheduledTaskResponseBody) *RemoveScheduledTaskResponse {
	s.Body = v
	return s
}

func (s *RemoveScheduledTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
