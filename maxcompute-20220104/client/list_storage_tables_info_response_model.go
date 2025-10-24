// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStorageTablesInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStorageTablesInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStorageTablesInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListStorageTablesInfoResponseBody) *ListStorageTablesInfoResponse
	GetBody() *ListStorageTablesInfoResponseBody
}

type ListStorageTablesInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStorageTablesInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStorageTablesInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStorageTablesInfoResponse) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStorageTablesInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStorageTablesInfoResponse) GetBody() *ListStorageTablesInfoResponseBody {
	return s.Body
}

func (s *ListStorageTablesInfoResponse) SetHeaders(v map[string]*string) *ListStorageTablesInfoResponse {
	s.Headers = v
	return s
}

func (s *ListStorageTablesInfoResponse) SetStatusCode(v int32) *ListStorageTablesInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStorageTablesInfoResponse) SetBody(v *ListStorageTablesInfoResponseBody) *ListStorageTablesInfoResponse {
	s.Body = v
	return s
}

func (s *ListStorageTablesInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
