// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTableMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTableMetaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTableMetaResponseBody) *UpdateTableMetaResponse
	GetBody() *UpdateTableMetaResponseBody
}

type UpdateTableMetaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTableMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTableMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTableMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTableMetaResponse) GetBody() *UpdateTableMetaResponseBody {
	return s.Body
}

func (s *UpdateTableMetaResponse) SetHeaders(v map[string]*string) *UpdateTableMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableMetaResponse) SetStatusCode(v int32) *UpdateTableMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableMetaResponse) SetBody(v *UpdateTableMetaResponseBody) *UpdateTableMetaResponse {
	s.Body = v
	return s
}

func (s *UpdateTableMetaResponse) Validate() error {
	return dara.Validate(s)
}
