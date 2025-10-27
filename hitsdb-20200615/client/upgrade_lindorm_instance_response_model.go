// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeLindormInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeLindormInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeLindormInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeLindormInstanceResponseBody) *UpgradeLindormInstanceResponse
	GetBody() *UpgradeLindormInstanceResponseBody
}

type UpgradeLindormInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeLindormInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeLindormInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeLindormInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeLindormInstanceResponse) GetBody() *UpgradeLindormInstanceResponseBody {
	return s.Body
}

func (s *UpgradeLindormInstanceResponse) SetHeaders(v map[string]*string) *UpgradeLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeLindormInstanceResponse) SetStatusCode(v int32) *UpgradeLindormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeLindormInstanceResponse) SetBody(v *UpgradeLindormInstanceResponseBody) *UpgradeLindormInstanceResponse {
	s.Body = v
	return s
}

func (s *UpgradeLindormInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
