// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCorpTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCorpTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCorpTasksResponse
	GetStatusCode() *int32
	SetBody(v *GetCorpTasksResponseBody) *GetCorpTasksResponse
	GetBody() *GetCorpTasksResponseBody
}

type GetCorpTasksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCorpTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCorpTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCorpTasksResponse) GoString() string {
	return s.String()
}

func (s *GetCorpTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCorpTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCorpTasksResponse) GetBody() *GetCorpTasksResponseBody {
	return s.Body
}

func (s *GetCorpTasksResponse) SetHeaders(v map[string]*string) *GetCorpTasksResponse {
	s.Headers = v
	return s
}

func (s *GetCorpTasksResponse) SetStatusCode(v int32) *GetCorpTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCorpTasksResponse) SetBody(v *GetCorpTasksResponseBody) *GetCorpTasksResponse {
	s.Body = v
	return s
}

func (s *GetCorpTasksResponse) Validate() error {
	return dara.Validate(s)
}
