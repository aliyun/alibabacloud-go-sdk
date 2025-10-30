// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpgradeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUpgradeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUpgradeStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetUpgradeStatusResponseBody) *GetUpgradeStatusResponse
	GetBody() *GetUpgradeStatusResponseBody
}

type GetUpgradeStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUpgradeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUpgradeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUpgradeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUpgradeStatusResponse) GetBody() *GetUpgradeStatusResponseBody {
	return s.Body
}

func (s *GetUpgradeStatusResponse) SetHeaders(v map[string]*string) *GetUpgradeStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUpgradeStatusResponse) SetStatusCode(v int32) *GetUpgradeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUpgradeStatusResponse) SetBody(v *GetUpgradeStatusResponseBody) *GetUpgradeStatusResponse {
	s.Body = v
	return s
}

func (s *GetUpgradeStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
