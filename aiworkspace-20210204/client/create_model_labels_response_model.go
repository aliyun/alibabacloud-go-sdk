// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelLabelsResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelLabelsResponseBody) *CreateModelLabelsResponse
	GetBody() *CreateModelLabelsResponseBody
}

type CreateModelLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelLabelsResponse) GoString() string {
	return s.String()
}

func (s *CreateModelLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelLabelsResponse) GetBody() *CreateModelLabelsResponseBody {
	return s.Body
}

func (s *CreateModelLabelsResponse) SetHeaders(v map[string]*string) *CreateModelLabelsResponse {
	s.Headers = v
	return s
}

func (s *CreateModelLabelsResponse) SetStatusCode(v int32) *CreateModelLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelLabelsResponse) SetBody(v *CreateModelLabelsResponseBody) *CreateModelLabelsResponse {
	s.Body = v
	return s
}

func (s *CreateModelLabelsResponse) Validate() error {
	return dara.Validate(s)
}
