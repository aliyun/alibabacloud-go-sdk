// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApiDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApiDefinitionResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *GetApiDefinitionResponse
	GetBody() map[string]interface{}
}

type GetApiDefinitionResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApiDefinitionResponse) GoString() string {
	return s.String()
}

func (s *GetApiDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApiDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApiDefinitionResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *GetApiDefinitionResponse) SetHeaders(v map[string]*string) *GetApiDefinitionResponse {
	s.Headers = v
	return s
}

func (s *GetApiDefinitionResponse) SetStatusCode(v int32) *GetApiDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiDefinitionResponse) SetBody(v map[string]interface{}) *GetApiDefinitionResponse {
	s.Body = v
	return s
}

func (s *GetApiDefinitionResponse) Validate() error {
	return dara.Validate(s)
}
