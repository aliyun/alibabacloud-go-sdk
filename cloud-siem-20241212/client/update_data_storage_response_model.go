// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataStorageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataStorageResponseBody) *UpdateDataStorageResponse
	GetBody() *UpdateDataStorageResponseBody
}

type UpdateDataStorageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataStorageResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataStorageResponse) GetBody() *UpdateDataStorageResponseBody {
	return s.Body
}

func (s *UpdateDataStorageResponse) SetHeaders(v map[string]*string) *UpdateDataStorageResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataStorageResponse) SetStatusCode(v int32) *UpdateDataStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataStorageResponse) SetBody(v *UpdateDataStorageResponseBody) *UpdateDataStorageResponse {
	s.Body = v
	return s
}

func (s *UpdateDataStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
