// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepareUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PrepareUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PrepareUpgradeResponse
	GetStatusCode() *int32
	SetBody(v *PrepareUpgradeResponseBody) *PrepareUpgradeResponse
	GetBody() *PrepareUpgradeResponseBody
}

type PrepareUpgradeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrepareUpgradeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrepareUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s PrepareUpgradeResponse) GoString() string {
	return s.String()
}

func (s *PrepareUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PrepareUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PrepareUpgradeResponse) GetBody() *PrepareUpgradeResponseBody {
	return s.Body
}

func (s *PrepareUpgradeResponse) SetHeaders(v map[string]*string) *PrepareUpgradeResponse {
	s.Headers = v
	return s
}

func (s *PrepareUpgradeResponse) SetStatusCode(v int32) *PrepareUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *PrepareUpgradeResponse) SetBody(v *PrepareUpgradeResponseBody) *PrepareUpgradeResponse {
	s.Body = v
	return s
}

func (s *PrepareUpgradeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
