// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTableLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTableLevelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTableLevelResponseBody) *UpdateTableLevelResponse
	GetBody() *UpdateTableLevelResponseBody
}

type UpdateTableLevelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTableLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTableLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableLevelResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTableLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTableLevelResponse) GetBody() *UpdateTableLevelResponseBody {
	return s.Body
}

func (s *UpdateTableLevelResponse) SetHeaders(v map[string]*string) *UpdateTableLevelResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableLevelResponse) SetStatusCode(v int32) *UpdateTableLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableLevelResponse) SetBody(v *UpdateTableLevelResponseBody) *UpdateTableLevelResponse {
	s.Body = v
	return s
}

func (s *UpdateTableLevelResponse) Validate() error {
	return dara.Validate(s)
}
