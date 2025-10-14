// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyActiveOperationTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyActiveOperationTasksResponse
	GetStatusCode() *int32
	SetBody(v *ModifyActiveOperationTasksResponseBody) *ModifyActiveOperationTasksResponse
	GetBody() *ModifyActiveOperationTasksResponseBody
}

type ModifyActiveOperationTasksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyActiveOperationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyActiveOperationTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationTasksResponse) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyActiveOperationTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyActiveOperationTasksResponse) GetBody() *ModifyActiveOperationTasksResponseBody {
	return s.Body
}

func (s *ModifyActiveOperationTasksResponse) SetHeaders(v map[string]*string) *ModifyActiveOperationTasksResponse {
	s.Headers = v
	return s
}

func (s *ModifyActiveOperationTasksResponse) SetStatusCode(v int32) *ModifyActiveOperationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyActiveOperationTasksResponse) SetBody(v *ModifyActiveOperationTasksResponseBody) *ModifyActiveOperationTasksResponse {
	s.Body = v
	return s
}

func (s *ModifyActiveOperationTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
