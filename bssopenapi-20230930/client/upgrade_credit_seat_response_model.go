// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeCreditSeatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeCreditSeatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeCreditSeatResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeCreditSeatResponseBody) *UpgradeCreditSeatResponse
	GetBody() *UpgradeCreditSeatResponseBody
}

type UpgradeCreditSeatResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeCreditSeatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeCreditSeatResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeCreditSeatResponse) GoString() string {
	return s.String()
}

func (s *UpgradeCreditSeatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeCreditSeatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeCreditSeatResponse) GetBody() *UpgradeCreditSeatResponseBody {
	return s.Body
}

func (s *UpgradeCreditSeatResponse) SetHeaders(v map[string]*string) *UpgradeCreditSeatResponse {
	s.Headers = v
	return s
}

func (s *UpgradeCreditSeatResponse) SetStatusCode(v int32) *UpgradeCreditSeatResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeCreditSeatResponse) SetBody(v *UpgradeCreditSeatResponseBody) *UpgradeCreditSeatResponse {
	s.Body = v
	return s
}

func (s *UpgradeCreditSeatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
