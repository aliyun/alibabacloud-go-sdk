// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAINodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAINodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAINodeResponse
	GetStatusCode() *int32
	SetBody(v *AddAINodeResponseBody) *AddAINodeResponse
	GetBody() *AddAINodeResponseBody
}

type AddAINodeResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAINodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAINodeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAINodeResponse) GoString() string {
	return s.String()
}

func (s *AddAINodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAINodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAINodeResponse) GetBody() *AddAINodeResponseBody {
	return s.Body
}

func (s *AddAINodeResponse) SetHeaders(v map[string]*string) *AddAINodeResponse {
	s.Headers = v
	return s
}

func (s *AddAINodeResponse) SetStatusCode(v int32) *AddAINodeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAINodeResponse) SetBody(v *AddAINodeResponseBody) *AddAINodeResponse {
	s.Body = v
	return s
}

func (s *AddAINodeResponse) Validate() error {
	return dara.Validate(s)
}
