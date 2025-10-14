// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUpgradeItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckUpgradeItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckUpgradeItemResponse
	GetStatusCode() *int32
	SetBody(v *CheckUpgradeItemResponseBody) *CheckUpgradeItemResponse
	GetBody() *CheckUpgradeItemResponseBody
}

type CheckUpgradeItemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUpgradeItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUpgradeItemResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUpgradeItemResponse) GoString() string {
	return s.String()
}

func (s *CheckUpgradeItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckUpgradeItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckUpgradeItemResponse) GetBody() *CheckUpgradeItemResponseBody {
	return s.Body
}

func (s *CheckUpgradeItemResponse) SetHeaders(v map[string]*string) *CheckUpgradeItemResponse {
	s.Headers = v
	return s
}

func (s *CheckUpgradeItemResponse) SetStatusCode(v int32) *CheckUpgradeItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUpgradeItemResponse) SetBody(v *CheckUpgradeItemResponseBody) *CheckUpgradeItemResponse {
	s.Body = v
	return s
}

func (s *CheckUpgradeItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
