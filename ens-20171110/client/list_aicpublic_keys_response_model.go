// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICPublicKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAICPublicKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAICPublicKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListAICPublicKeysResponseBody) *ListAICPublicKeysResponse
	GetBody() *ListAICPublicKeysResponseBody
}

type ListAICPublicKeysResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAICPublicKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAICPublicKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAICPublicKeysResponse) GoString() string {
	return s.String()
}

func (s *ListAICPublicKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAICPublicKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAICPublicKeysResponse) GetBody() *ListAICPublicKeysResponseBody {
	return s.Body
}

func (s *ListAICPublicKeysResponse) SetHeaders(v map[string]*string) *ListAICPublicKeysResponse {
	s.Headers = v
	return s
}

func (s *ListAICPublicKeysResponse) SetStatusCode(v int32) *ListAICPublicKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAICPublicKeysResponse) SetBody(v *ListAICPublicKeysResponseBody) *ListAICPublicKeysResponse {
	s.Body = v
	return s
}

func (s *ListAICPublicKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
