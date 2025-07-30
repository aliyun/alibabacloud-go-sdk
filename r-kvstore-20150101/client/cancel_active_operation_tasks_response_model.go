// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelActiveOperationTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelActiveOperationTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelActiveOperationTasksResponse
	GetStatusCode() *int32
	SetBody(v *CancelActiveOperationTasksResponseBody) *CancelActiveOperationTasksResponse
	GetBody() *CancelActiveOperationTasksResponseBody
}

type CancelActiveOperationTasksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelActiveOperationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelActiveOperationTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelActiveOperationTasksResponse) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelActiveOperationTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelActiveOperationTasksResponse) GetBody() *CancelActiveOperationTasksResponseBody {
	return s.Body
}

func (s *CancelActiveOperationTasksResponse) SetHeaders(v map[string]*string) *CancelActiveOperationTasksResponse {
	s.Headers = v
	return s
}

func (s *CancelActiveOperationTasksResponse) SetStatusCode(v int32) *CancelActiveOperationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelActiveOperationTasksResponse) SetBody(v *CancelActiveOperationTasksResponseBody) *CancelActiveOperationTasksResponse {
	s.Body = v
	return s
}

func (s *CancelActiveOperationTasksResponse) Validate() error {
	return dara.Validate(s)
}
