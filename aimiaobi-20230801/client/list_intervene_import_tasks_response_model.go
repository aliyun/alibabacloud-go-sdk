// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterveneImportTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInterveneImportTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInterveneImportTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListInterveneImportTasksResponseBody) *ListInterveneImportTasksResponse
	GetBody() *ListInterveneImportTasksResponseBody
}

type ListInterveneImportTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterveneImportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterveneImportTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneImportTasksResponse) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInterveneImportTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInterveneImportTasksResponse) GetBody() *ListInterveneImportTasksResponseBody {
	return s.Body
}

func (s *ListInterveneImportTasksResponse) SetHeaders(v map[string]*string) *ListInterveneImportTasksResponse {
	s.Headers = v
	return s
}

func (s *ListInterveneImportTasksResponse) SetStatusCode(v int32) *ListInterveneImportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterveneImportTasksResponse) SetBody(v *ListInterveneImportTasksResponseBody) *ListInterveneImportTasksResponse {
	s.Body = v
	return s
}

func (s *ListInterveneImportTasksResponse) Validate() error {
	return dara.Validate(s)
}
