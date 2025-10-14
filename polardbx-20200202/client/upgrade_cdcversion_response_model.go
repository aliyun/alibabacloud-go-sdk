// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeCDCVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeCDCVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeCDCVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeCDCVersionResponseBody) *UpgradeCDCVersionResponse
	GetBody() *UpgradeCDCVersionResponseBody
}

type UpgradeCDCVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeCDCVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeCDCVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeCDCVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeCDCVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeCDCVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeCDCVersionResponse) GetBody() *UpgradeCDCVersionResponseBody {
	return s.Body
}

func (s *UpgradeCDCVersionResponse) SetHeaders(v map[string]*string) *UpgradeCDCVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeCDCVersionResponse) SetStatusCode(v int32) *UpgradeCDCVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeCDCVersionResponse) SetBody(v *UpgradeCDCVersionResponseBody) *UpgradeCDCVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeCDCVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
