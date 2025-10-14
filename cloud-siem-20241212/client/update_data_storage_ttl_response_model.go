// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataStorageTtlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataStorageTtlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataStorageTtlResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataStorageTtlResponseBody) *UpdateDataStorageTtlResponse
	GetBody() *UpdateDataStorageTtlResponseBody
}

type UpdateDataStorageTtlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataStorageTtlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataStorageTtlResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataStorageTtlResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataStorageTtlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataStorageTtlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataStorageTtlResponse) GetBody() *UpdateDataStorageTtlResponseBody {
	return s.Body
}

func (s *UpdateDataStorageTtlResponse) SetHeaders(v map[string]*string) *UpdateDataStorageTtlResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataStorageTtlResponse) SetStatusCode(v int32) *UpdateDataStorageTtlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataStorageTtlResponse) SetBody(v *UpdateDataStorageTtlResponseBody) *UpdateDataStorageTtlResponse {
	s.Body = v
	return s
}

func (s *UpdateDataStorageTtlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
