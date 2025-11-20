// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelProxyResponse
	GetStatusCode() *int32
	SetBody(v *ModelProxyResult) *GetModelProxyResponse
	GetBody() *ModelProxyResult
}

type GetModelProxyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelProxyResult  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelProxyResponse) GoString() string {
	return s.String()
}

func (s *GetModelProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelProxyResponse) GetBody() *ModelProxyResult {
	return s.Body
}

func (s *GetModelProxyResponse) SetHeaders(v map[string]*string) *GetModelProxyResponse {
	s.Headers = v
	return s
}

func (s *GetModelProxyResponse) SetStatusCode(v int32) *GetModelProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelProxyResponse) SetBody(v *ModelProxyResult) *GetModelProxyResponse {
	s.Body = v
	return s
}

func (s *GetModelProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
