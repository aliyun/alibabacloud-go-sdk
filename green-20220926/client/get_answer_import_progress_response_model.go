// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnswerImportProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAnswerImportProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAnswerImportProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetAnswerImportProgressResponseBody) *GetAnswerImportProgressResponse
	GetBody() *GetAnswerImportProgressResponseBody
}

type GetAnswerImportProgressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAnswerImportProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAnswerImportProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAnswerImportProgressResponse) GoString() string {
	return s.String()
}

func (s *GetAnswerImportProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAnswerImportProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAnswerImportProgressResponse) GetBody() *GetAnswerImportProgressResponseBody {
	return s.Body
}

func (s *GetAnswerImportProgressResponse) SetHeaders(v map[string]*string) *GetAnswerImportProgressResponse {
	s.Headers = v
	return s
}

func (s *GetAnswerImportProgressResponse) SetStatusCode(v int32) *GetAnswerImportProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAnswerImportProgressResponse) SetBody(v *GetAnswerImportProgressResponseBody) *GetAnswerImportProgressResponse {
	s.Body = v
	return s
}

func (s *GetAnswerImportProgressResponse) Validate() error {
	return dara.Validate(s)
}
