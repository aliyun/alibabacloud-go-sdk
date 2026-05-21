// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscardUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DiscardUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DiscardUpgradeResponse
	GetStatusCode() *int32
	SetBody(v *DiscardUpgradeResponseBody) *DiscardUpgradeResponse
	GetBody() *DiscardUpgradeResponseBody
}

type DiscardUpgradeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DiscardUpgradeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DiscardUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s DiscardUpgradeResponse) GoString() string {
	return s.String()
}

func (s *DiscardUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DiscardUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DiscardUpgradeResponse) GetBody() *DiscardUpgradeResponseBody {
	return s.Body
}

func (s *DiscardUpgradeResponse) SetHeaders(v map[string]*string) *DiscardUpgradeResponse {
	s.Headers = v
	return s
}

func (s *DiscardUpgradeResponse) SetStatusCode(v int32) *DiscardUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *DiscardUpgradeResponse) SetBody(v *DiscardUpgradeResponseBody) *DiscardUpgradeResponse {
	s.Body = v
	return s
}

func (s *DiscardUpgradeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
