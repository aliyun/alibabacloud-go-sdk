// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGeneratedContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGeneratedContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGeneratedContentResponse
	GetStatusCode() *int32
	SetBody(v *GetGeneratedContentResponseBody) *GetGeneratedContentResponse
	GetBody() *GetGeneratedContentResponseBody
}

type GetGeneratedContentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGeneratedContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGeneratedContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGeneratedContentResponse) GetBody() *GetGeneratedContentResponseBody {
	return s.Body
}

func (s *GetGeneratedContentResponse) SetHeaders(v map[string]*string) *GetGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *GetGeneratedContentResponse) SetStatusCode(v int32) *GetGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGeneratedContentResponse) SetBody(v *GetGeneratedContentResponseBody) *GetGeneratedContentResponse {
	s.Body = v
	return s
}

func (s *GetGeneratedContentResponse) Validate() error {
	return dara.Validate(s)
}
