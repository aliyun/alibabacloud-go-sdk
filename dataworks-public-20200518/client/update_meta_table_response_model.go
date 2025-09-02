// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMetaTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMetaTableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMetaTableResponseBody) *UpdateMetaTableResponse
	GetBody() *UpdateMetaTableResponseBody
}

type UpdateMetaTableResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMetaTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMetaTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetaTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMetaTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMetaTableResponse) GetBody() *UpdateMetaTableResponseBody {
	return s.Body
}

func (s *UpdateMetaTableResponse) SetHeaders(v map[string]*string) *UpdateMetaTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetaTableResponse) SetStatusCode(v int32) *UpdateMetaTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMetaTableResponse) SetBody(v *UpdateMetaTableResponseBody) *UpdateMetaTableResponse {
	s.Body = v
	return s
}

func (s *UpdateMetaTableResponse) Validate() error {
	return dara.Validate(s)
}
