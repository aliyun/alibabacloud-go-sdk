// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImportLogTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitImportLogTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitImportLogTasksResponse
	GetStatusCode() *int32
	SetBody(v *SubmitImportLogTasksResponseBody) *SubmitImportLogTasksResponse
	GetBody() *SubmitImportLogTasksResponseBody
}

type SubmitImportLogTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitImportLogTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitImportLogTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitImportLogTasksResponse) GoString() string {
	return s.String()
}

func (s *SubmitImportLogTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitImportLogTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitImportLogTasksResponse) GetBody() *SubmitImportLogTasksResponseBody {
	return s.Body
}

func (s *SubmitImportLogTasksResponse) SetHeaders(v map[string]*string) *SubmitImportLogTasksResponse {
	s.Headers = v
	return s
}

func (s *SubmitImportLogTasksResponse) SetStatusCode(v int32) *SubmitImportLogTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitImportLogTasksResponse) SetBody(v *SubmitImportLogTasksResponseBody) *SubmitImportLogTasksResponse {
	s.Body = v
	return s
}

func (s *SubmitImportLogTasksResponse) Validate() error {
	return dara.Validate(s)
}
