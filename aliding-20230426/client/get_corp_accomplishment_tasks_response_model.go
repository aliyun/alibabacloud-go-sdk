// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCorpAccomplishmentTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCorpAccomplishmentTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCorpAccomplishmentTasksResponse
	GetStatusCode() *int32
	SetBody(v *GetCorpAccomplishmentTasksResponseBody) *GetCorpAccomplishmentTasksResponse
	GetBody() *GetCorpAccomplishmentTasksResponseBody
}

type GetCorpAccomplishmentTasksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCorpAccomplishmentTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCorpAccomplishmentTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCorpAccomplishmentTasksResponse) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCorpAccomplishmentTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCorpAccomplishmentTasksResponse) GetBody() *GetCorpAccomplishmentTasksResponseBody {
	return s.Body
}

func (s *GetCorpAccomplishmentTasksResponse) SetHeaders(v map[string]*string) *GetCorpAccomplishmentTasksResponse {
	s.Headers = v
	return s
}

func (s *GetCorpAccomplishmentTasksResponse) SetStatusCode(v int32) *GetCorpAccomplishmentTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponse) SetBody(v *GetCorpAccomplishmentTasksResponseBody) *GetCorpAccomplishmentTasksResponse {
	s.Body = v
	return s
}

func (s *GetCorpAccomplishmentTasksResponse) Validate() error {
	return dara.Validate(s)
}
