// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeInfoResponseBody) *UpgradeInfoResponse
	GetBody() *UpgradeInfoResponseBody
}

type UpgradeInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInfoResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeInfoResponse) GetBody() *UpgradeInfoResponseBody {
	return s.Body
}

func (s *UpgradeInfoResponse) SetHeaders(v map[string]*string) *UpgradeInfoResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInfoResponse) SetStatusCode(v int32) *UpgradeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeInfoResponse) SetBody(v *UpgradeInfoResponseBody) *UpgradeInfoResponse {
	s.Body = v
	return s
}

func (s *UpgradeInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
