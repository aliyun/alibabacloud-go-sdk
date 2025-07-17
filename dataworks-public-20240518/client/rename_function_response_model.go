// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenameFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenameFunctionResponse
	GetStatusCode() *int32
	SetBody(v *RenameFunctionResponseBody) *RenameFunctionResponse
	GetBody() *RenameFunctionResponseBody
}

type RenameFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s RenameFunctionResponse) GoString() string {
	return s.String()
}

func (s *RenameFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenameFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenameFunctionResponse) GetBody() *RenameFunctionResponseBody {
	return s.Body
}

func (s *RenameFunctionResponse) SetHeaders(v map[string]*string) *RenameFunctionResponse {
	s.Headers = v
	return s
}

func (s *RenameFunctionResponse) SetStatusCode(v int32) *RenameFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameFunctionResponse) SetBody(v *RenameFunctionResponseBody) *RenameFunctionResponse {
	s.Body = v
	return s
}

func (s *RenameFunctionResponse) Validate() error {
	return dara.Validate(s)
}
