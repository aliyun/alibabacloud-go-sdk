// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKibanaPvlNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKibanaPvlNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKibanaPvlNetworkResponse
	GetStatusCode() *int32
	SetBody(v *ListKibanaPvlNetworkResponseBody) *ListKibanaPvlNetworkResponse
	GetBody() *ListKibanaPvlNetworkResponseBody
}

type ListKibanaPvlNetworkResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKibanaPvlNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKibanaPvlNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKibanaPvlNetworkResponse) GoString() string {
	return s.String()
}

func (s *ListKibanaPvlNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKibanaPvlNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKibanaPvlNetworkResponse) GetBody() *ListKibanaPvlNetworkResponseBody {
	return s.Body
}

func (s *ListKibanaPvlNetworkResponse) SetHeaders(v map[string]*string) *ListKibanaPvlNetworkResponse {
	s.Headers = v
	return s
}

func (s *ListKibanaPvlNetworkResponse) SetStatusCode(v int32) *ListKibanaPvlNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKibanaPvlNetworkResponse) SetBody(v *ListKibanaPvlNetworkResponseBody) *ListKibanaPvlNetworkResponse {
	s.Body = v
	return s
}

func (s *ListKibanaPvlNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
