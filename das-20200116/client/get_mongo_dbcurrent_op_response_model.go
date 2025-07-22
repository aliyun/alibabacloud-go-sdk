// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMongoDBCurrentOpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMongoDBCurrentOpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMongoDBCurrentOpResponse
	GetStatusCode() *int32
	SetBody(v *GetMongoDBCurrentOpResponseBody) *GetMongoDBCurrentOpResponse
	GetBody() *GetMongoDBCurrentOpResponseBody
}

type GetMongoDBCurrentOpResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMongoDBCurrentOpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMongoDBCurrentOpResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMongoDBCurrentOpResponse) GoString() string {
	return s.String()
}

func (s *GetMongoDBCurrentOpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMongoDBCurrentOpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMongoDBCurrentOpResponse) GetBody() *GetMongoDBCurrentOpResponseBody {
	return s.Body
}

func (s *GetMongoDBCurrentOpResponse) SetHeaders(v map[string]*string) *GetMongoDBCurrentOpResponse {
	s.Headers = v
	return s
}

func (s *GetMongoDBCurrentOpResponse) SetStatusCode(v int32) *GetMongoDBCurrentOpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMongoDBCurrentOpResponse) SetBody(v *GetMongoDBCurrentOpResponseBody) *GetMongoDBCurrentOpResponse {
	s.Body = v
	return s
}

func (s *GetMongoDBCurrentOpResponse) Validate() error {
	return dara.Validate(s)
}
