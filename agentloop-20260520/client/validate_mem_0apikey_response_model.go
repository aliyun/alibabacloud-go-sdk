// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateMem0APIKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateMem0APIKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateMem0APIKeyResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *ValidateMem0APIKeyResponse
	GetBody() map[string]interface{}
}

type ValidateMem0APIKeyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateMem0APIKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateMem0APIKeyResponse) GoString() string {
	return s.String()
}

func (s *ValidateMem0APIKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateMem0APIKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateMem0APIKeyResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *ValidateMem0APIKeyResponse) SetHeaders(v map[string]*string) *ValidateMem0APIKeyResponse {
	s.Headers = v
	return s
}

func (s *ValidateMem0APIKeyResponse) SetStatusCode(v int32) *ValidateMem0APIKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateMem0APIKeyResponse) SetBody(v map[string]interface{}) *ValidateMem0APIKeyResponse {
	s.Body = v
	return s
}

func (s *ValidateMem0APIKeyResponse) Validate() error {
	return dara.Validate(s)
}
