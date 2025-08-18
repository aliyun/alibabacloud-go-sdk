// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateListResponse
	GetStatusCode() *int32
	SetBody(v *CreateListResponseBody) *CreateListResponse
	GetBody() *CreateListResponseBody
}

type CreateListResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateListResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateListResponse) GoString() string {
	return s.String()
}

func (s *CreateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateListResponse) GetBody() *CreateListResponseBody {
	return s.Body
}

func (s *CreateListResponse) SetHeaders(v map[string]*string) *CreateListResponse {
	s.Headers = v
	return s
}

func (s *CreateListResponse) SetStatusCode(v int32) *CreateListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateListResponse) SetBody(v *CreateListResponseBody) *CreateListResponse {
	s.Body = v
	return s
}

func (s *CreateListResponse) Validate() error {
	return dara.Validate(s)
}
