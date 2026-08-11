// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseApiKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseApiKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseApiKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseApiKeysResponseBody) *ListContextDatabaseApiKeysResponse
	GetBody() *ListContextDatabaseApiKeysResponseBody
}

type ListContextDatabaseApiKeysResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseApiKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseApiKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseApiKeysResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseApiKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseApiKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseApiKeysResponse) GetBody() *ListContextDatabaseApiKeysResponseBody {
	return s.Body
}

func (s *ListContextDatabaseApiKeysResponse) SetHeaders(v map[string]*string) *ListContextDatabaseApiKeysResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseApiKeysResponse) SetStatusCode(v int32) *ListContextDatabaseApiKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponse) SetBody(v *ListContextDatabaseApiKeysResponseBody) *ListContextDatabaseApiKeysResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseApiKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
