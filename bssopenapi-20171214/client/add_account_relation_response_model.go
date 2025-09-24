// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAccountRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAccountRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAccountRelationResponse
	GetStatusCode() *int32
	SetBody(v *AddAccountRelationResponseBody) *AddAccountRelationResponse
	GetBody() *AddAccountRelationResponseBody
}

type AddAccountRelationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAccountRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAccountRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAccountRelationResponse) GoString() string {
	return s.String()
}

func (s *AddAccountRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAccountRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAccountRelationResponse) GetBody() *AddAccountRelationResponseBody {
	return s.Body
}

func (s *AddAccountRelationResponse) SetHeaders(v map[string]*string) *AddAccountRelationResponse {
	s.Headers = v
	return s
}

func (s *AddAccountRelationResponse) SetStatusCode(v int32) *AddAccountRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAccountRelationResponse) SetBody(v *AddAccountRelationResponseBody) *AddAccountRelationResponse {
	s.Body = v
	return s
}

func (s *AddAccountRelationResponse) Validate() error {
	return dara.Validate(s)
}
