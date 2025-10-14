// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateLogStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateLogStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateLogStoreResponse
	GetStatusCode() *int32
	SetBody(v *ValidateLogStoreResponseBody) *ValidateLogStoreResponse
	GetBody() *ValidateLogStoreResponseBody
}

type ValidateLogStoreResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateLogStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateLogStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateLogStoreResponse) GoString() string {
	return s.String()
}

func (s *ValidateLogStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateLogStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateLogStoreResponse) GetBody() *ValidateLogStoreResponseBody {
	return s.Body
}

func (s *ValidateLogStoreResponse) SetHeaders(v map[string]*string) *ValidateLogStoreResponse {
	s.Headers = v
	return s
}

func (s *ValidateLogStoreResponse) SetStatusCode(v int32) *ValidateLogStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateLogStoreResponse) SetBody(v *ValidateLogStoreResponseBody) *ValidateLogStoreResponse {
	s.Body = v
	return s
}

func (s *ValidateLogStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
