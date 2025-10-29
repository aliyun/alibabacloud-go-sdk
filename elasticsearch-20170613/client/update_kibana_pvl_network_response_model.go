// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKibanaPvlNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKibanaPvlNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKibanaPvlNetworkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKibanaPvlNetworkResponseBody) *UpdateKibanaPvlNetworkResponse
	GetBody() *UpdateKibanaPvlNetworkResponseBody
}

type UpdateKibanaPvlNetworkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKibanaPvlNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKibanaPvlNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaPvlNetworkResponse) GoString() string {
	return s.String()
}

func (s *UpdateKibanaPvlNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKibanaPvlNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKibanaPvlNetworkResponse) GetBody() *UpdateKibanaPvlNetworkResponseBody {
	return s.Body
}

func (s *UpdateKibanaPvlNetworkResponse) SetHeaders(v map[string]*string) *UpdateKibanaPvlNetworkResponse {
	s.Headers = v
	return s
}

func (s *UpdateKibanaPvlNetworkResponse) SetStatusCode(v int32) *UpdateKibanaPvlNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKibanaPvlNetworkResponse) SetBody(v *UpdateKibanaPvlNetworkResponseBody) *UpdateKibanaPvlNetworkResponse {
	s.Body = v
	return s
}

func (s *UpdateKibanaPvlNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
