// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeColumnarVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeColumnarVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeColumnarVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeColumnarVersionResponseBody) *UpgradeColumnarVersionResponse
	GetBody() *UpgradeColumnarVersionResponseBody
}

type UpgradeColumnarVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeColumnarVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeColumnarVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeColumnarVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeColumnarVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeColumnarVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeColumnarVersionResponse) GetBody() *UpgradeColumnarVersionResponseBody {
	return s.Body
}

func (s *UpgradeColumnarVersionResponse) SetHeaders(v map[string]*string) *UpgradeColumnarVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeColumnarVersionResponse) SetStatusCode(v int32) *UpgradeColumnarVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeColumnarVersionResponse) SetBody(v *UpgradeColumnarVersionResponseBody) *UpgradeColumnarVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeColumnarVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
