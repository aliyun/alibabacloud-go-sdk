// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelProxyResponse
	GetStatusCode() *int32
	SetBody(v *ModelProxyResult) *CreateModelProxyResponse
	GetBody() *ModelProxyResult
}

type CreateModelProxyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelProxyResult  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProxyResponse) GoString() string {
	return s.String()
}

func (s *CreateModelProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelProxyResponse) GetBody() *ModelProxyResult {
	return s.Body
}

func (s *CreateModelProxyResponse) SetHeaders(v map[string]*string) *CreateModelProxyResponse {
	s.Headers = v
	return s
}

func (s *CreateModelProxyResponse) SetStatusCode(v int32) *CreateModelProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelProxyResponse) SetBody(v *ModelProxyResult) *CreateModelProxyResponse {
	s.Body = v
	return s
}

func (s *CreateModelProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
