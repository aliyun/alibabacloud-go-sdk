// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeApplicationVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeApplicationVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeApplicationVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeApplicationVersionResponseBody) *UpgradeApplicationVersionResponse
	GetBody() *UpgradeApplicationVersionResponseBody
}

type UpgradeApplicationVersionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeApplicationVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeApplicationVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeApplicationVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeApplicationVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeApplicationVersionResponse) GetBody() *UpgradeApplicationVersionResponseBody {
	return s.Body
}

func (s *UpgradeApplicationVersionResponse) SetHeaders(v map[string]*string) *UpgradeApplicationVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeApplicationVersionResponse) SetStatusCode(v int32) *UpgradeApplicationVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeApplicationVersionResponse) SetBody(v *UpgradeApplicationVersionResponseBody) *UpgradeApplicationVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeApplicationVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
