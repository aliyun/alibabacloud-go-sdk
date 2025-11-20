// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMultiDimTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMultiDimTableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMultiDimTableResponseBody) *UpdateMultiDimTableResponse
	GetBody() *UpdateMultiDimTableResponseBody
}

type UpdateMultiDimTableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMultiDimTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMultiDimTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMultiDimTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMultiDimTableResponse) GetBody() *UpdateMultiDimTableResponseBody {
	return s.Body
}

func (s *UpdateMultiDimTableResponse) SetHeaders(v map[string]*string) *UpdateMultiDimTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateMultiDimTableResponse) SetStatusCode(v int32) *UpdateMultiDimTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMultiDimTableResponse) SetBody(v *UpdateMultiDimTableResponseBody) *UpdateMultiDimTableResponse {
	s.Body = v
	return s
}

func (s *UpdateMultiDimTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
