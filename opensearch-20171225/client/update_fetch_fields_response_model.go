// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFetchFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFetchFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFetchFieldsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFetchFieldsResponseBody) *UpdateFetchFieldsResponse
	GetBody() *UpdateFetchFieldsResponseBody
}

type UpdateFetchFieldsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFetchFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFetchFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFetchFieldsResponse) GoString() string {
	return s.String()
}

func (s *UpdateFetchFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFetchFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFetchFieldsResponse) GetBody() *UpdateFetchFieldsResponseBody {
	return s.Body
}

func (s *UpdateFetchFieldsResponse) SetHeaders(v map[string]*string) *UpdateFetchFieldsResponse {
	s.Headers = v
	return s
}

func (s *UpdateFetchFieldsResponse) SetStatusCode(v int32) *UpdateFetchFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFetchFieldsResponse) SetBody(v *UpdateFetchFieldsResponseBody) *UpdateFetchFieldsResponse {
	s.Body = v
	return s
}

func (s *UpdateFetchFieldsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
