// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmRelationResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmRelationResponseBody) *ConfirmRelationResponse
	GetBody() *ConfirmRelationResponseBody
}

type ConfirmRelationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmRelationResponse) GoString() string {
	return s.String()
}

func (s *ConfirmRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmRelationResponse) GetBody() *ConfirmRelationResponseBody {
	return s.Body
}

func (s *ConfirmRelationResponse) SetHeaders(v map[string]*string) *ConfirmRelationResponse {
	s.Headers = v
	return s
}

func (s *ConfirmRelationResponse) SetStatusCode(v int32) *ConfirmRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmRelationResponse) SetBody(v *ConfirmRelationResponseBody) *ConfirmRelationResponse {
	s.Body = v
	return s
}

func (s *ConfirmRelationResponse) Validate() error {
	return dara.Validate(s)
}
