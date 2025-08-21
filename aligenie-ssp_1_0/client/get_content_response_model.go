// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContentResponse
	GetStatusCode() *int32
	SetBody(v *GetContentResponseBody) *GetContentResponse
	GetBody() *GetContentResponseBody
}

type GetContentResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContentResponse) GoString() string {
	return s.String()
}

func (s *GetContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContentResponse) GetBody() *GetContentResponseBody {
	return s.Body
}

func (s *GetContentResponse) SetHeaders(v map[string]*string) *GetContentResponse {
	s.Headers = v
	return s
}

func (s *GetContentResponse) SetStatusCode(v int32) *GetContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContentResponse) SetBody(v *GetContentResponseBody) *GetContentResponse {
	s.Body = v
	return s
}

func (s *GetContentResponse) Validate() error {
	return dara.Validate(s)
}
