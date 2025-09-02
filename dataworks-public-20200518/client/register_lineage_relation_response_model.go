// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterLineageRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterLineageRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterLineageRelationResponse
	GetStatusCode() *int32
	SetBody(v *RegisterLineageRelationResponseBody) *RegisterLineageRelationResponse
	GetBody() *RegisterLineageRelationResponseBody
}

type RegisterLineageRelationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterLineageRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterLineageRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterLineageRelationResponse) GoString() string {
	return s.String()
}

func (s *RegisterLineageRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterLineageRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterLineageRelationResponse) GetBody() *RegisterLineageRelationResponseBody {
	return s.Body
}

func (s *RegisterLineageRelationResponse) SetHeaders(v map[string]*string) *RegisterLineageRelationResponse {
	s.Headers = v
	return s
}

func (s *RegisterLineageRelationResponse) SetStatusCode(v int32) *RegisterLineageRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterLineageRelationResponse) SetBody(v *RegisterLineageRelationResponseBody) *RegisterLineageRelationResponse {
	s.Body = v
	return s
}

func (s *RegisterLineageRelationResponse) Validate() error {
	return dara.Validate(s)
}
