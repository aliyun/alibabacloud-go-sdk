// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelProxyResponse
	GetStatusCode() *int32
	SetBody(v *ModelProxyResult) *UpdateModelProxyResponse
	GetBody() *ModelProxyResult
}

type UpdateModelProxyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelProxyResult  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelProxyResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelProxyResponse) GetBody() *ModelProxyResult {
	return s.Body
}

func (s *UpdateModelProxyResponse) SetHeaders(v map[string]*string) *UpdateModelProxyResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelProxyResponse) SetStatusCode(v int32) *UpdateModelProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelProxyResponse) SetBody(v *ModelProxyResult) *UpdateModelProxyResponse {
	s.Body = v
	return s
}

func (s *UpdateModelProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
