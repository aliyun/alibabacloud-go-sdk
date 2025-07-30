// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeTwoWayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeTwoWayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeTwoWayResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeTwoWayResponseBody) *UpgradeTwoWayResponse
	GetBody() *UpgradeTwoWayResponseBody
}

type UpgradeTwoWayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeTwoWayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeTwoWayResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeTwoWayResponse) GoString() string {
	return s.String()
}

func (s *UpgradeTwoWayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeTwoWayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeTwoWayResponse) GetBody() *UpgradeTwoWayResponseBody {
	return s.Body
}

func (s *UpgradeTwoWayResponse) SetHeaders(v map[string]*string) *UpgradeTwoWayResponse {
	s.Headers = v
	return s
}

func (s *UpgradeTwoWayResponse) SetStatusCode(v int32) *UpgradeTwoWayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeTwoWayResponse) SetBody(v *UpgradeTwoWayResponseBody) *UpgradeTwoWayResponse {
	s.Body = v
	return s
}

func (s *UpgradeTwoWayResponse) Validate() error {
	return dara.Validate(s)
}
