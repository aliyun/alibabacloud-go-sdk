// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpgradeableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUpgradeableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUpgradeableResponse
	GetStatusCode() *int32
	SetBody(v *GetUpgradeableResponseBody) *GetUpgradeableResponse
	GetBody() *GetUpgradeableResponseBody
}

type GetUpgradeableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUpgradeableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUpgradeableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUpgradeableResponse) GoString() string {
	return s.String()
}

func (s *GetUpgradeableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUpgradeableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUpgradeableResponse) GetBody() *GetUpgradeableResponseBody {
	return s.Body
}

func (s *GetUpgradeableResponse) SetHeaders(v map[string]*string) *GetUpgradeableResponse {
	s.Headers = v
	return s
}

func (s *GetUpgradeableResponse) SetStatusCode(v int32) *GetUpgradeableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUpgradeableResponse) SetBody(v *GetUpgradeableResponseBody) *GetUpgradeableResponse {
	s.Body = v
	return s
}

func (s *GetUpgradeableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
