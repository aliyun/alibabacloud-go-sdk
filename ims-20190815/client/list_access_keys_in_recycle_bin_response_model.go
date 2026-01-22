// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessKeysInRecycleBinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccessKeysInRecycleBinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccessKeysInRecycleBinResponse
	GetStatusCode() *int32
	SetBody(v *ListAccessKeysInRecycleBinResponseBody) *ListAccessKeysInRecycleBinResponse
	GetBody() *ListAccessKeysInRecycleBinResponseBody
}

type ListAccessKeysInRecycleBinResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessKeysInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessKeysInRecycleBinResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccessKeysInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *ListAccessKeysInRecycleBinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccessKeysInRecycleBinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccessKeysInRecycleBinResponse) GetBody() *ListAccessKeysInRecycleBinResponseBody {
	return s.Body
}

func (s *ListAccessKeysInRecycleBinResponse) SetHeaders(v map[string]*string) *ListAccessKeysInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *ListAccessKeysInRecycleBinResponse) SetStatusCode(v int32) *ListAccessKeysInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessKeysInRecycleBinResponse) SetBody(v *ListAccessKeysInRecycleBinResponseBody) *ListAccessKeysInRecycleBinResponse {
	s.Body = v
	return s
}

func (s *ListAccessKeysInRecycleBinResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
