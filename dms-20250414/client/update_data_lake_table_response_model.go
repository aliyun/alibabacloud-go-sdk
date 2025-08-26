// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataLakeTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataLakeTableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataLakeTableResponseBody) *UpdateDataLakeTableResponse
	GetBody() *UpdateDataLakeTableResponseBody
}

type UpdateDataLakeTableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataLakeTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataLakeTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataLakeTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataLakeTableResponse) GetBody() *UpdateDataLakeTableResponseBody {
	return s.Body
}

func (s *UpdateDataLakeTableResponse) SetHeaders(v map[string]*string) *UpdateDataLakeTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataLakeTableResponse) SetStatusCode(v int32) *UpdateDataLakeTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataLakeTableResponse) SetBody(v *UpdateDataLakeTableResponseBody) *UpdateDataLakeTableResponse {
	s.Body = v
	return s
}

func (s *UpdateDataLakeTableResponse) Validate() error {
	return dara.Validate(s)
}
