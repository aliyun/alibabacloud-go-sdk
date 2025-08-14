// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSearchIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSearchIndexResponse
	GetStatusCode() *int32
	SetBody(v *CreateSearchIndexResponseBody) *CreateSearchIndexResponse
	GetBody() *CreateSearchIndexResponseBody
}

type CreateSearchIndexResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSearchIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSearchIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateSearchIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSearchIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSearchIndexResponse) GetBody() *CreateSearchIndexResponseBody {
	return s.Body
}

func (s *CreateSearchIndexResponse) SetHeaders(v map[string]*string) *CreateSearchIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateSearchIndexResponse) SetStatusCode(v int32) *CreateSearchIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSearchIndexResponse) SetBody(v *CreateSearchIndexResponseBody) *CreateSearchIndexResponse {
	s.Body = v
	return s
}

func (s *CreateSearchIndexResponse) Validate() error {
	return dara.Validate(s)
}
