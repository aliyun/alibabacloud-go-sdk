// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiStageVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApiStageVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApiStageVariableResponse
	GetStatusCode() *int32
	SetBody(v *CreateApiStageVariableResponseBody) *CreateApiStageVariableResponse
	GetBody() *CreateApiStageVariableResponseBody
}

type CreateApiStageVariableResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiStageVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiStageVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApiStageVariableResponse) GoString() string {
	return s.String()
}

func (s *CreateApiStageVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApiStageVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApiStageVariableResponse) GetBody() *CreateApiStageVariableResponseBody {
	return s.Body
}

func (s *CreateApiStageVariableResponse) SetHeaders(v map[string]*string) *CreateApiStageVariableResponse {
	s.Headers = v
	return s
}

func (s *CreateApiStageVariableResponse) SetStatusCode(v int32) *CreateApiStageVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiStageVariableResponse) SetBody(v *CreateApiStageVariableResponseBody) *CreateApiStageVariableResponse {
	s.Body = v
	return s
}

func (s *CreateApiStageVariableResponse) Validate() error {
	return dara.Validate(s)
}
