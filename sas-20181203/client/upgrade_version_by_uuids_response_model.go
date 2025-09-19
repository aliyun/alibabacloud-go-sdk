// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeVersionByUuidsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeVersionByUuidsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeVersionByUuidsResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeVersionByUuidsResponseBody) *UpgradeVersionByUuidsResponse
	GetBody() *UpgradeVersionByUuidsResponseBody
}

type UpgradeVersionByUuidsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeVersionByUuidsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeVersionByUuidsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeVersionByUuidsResponse) GoString() string {
	return s.String()
}

func (s *UpgradeVersionByUuidsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeVersionByUuidsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeVersionByUuidsResponse) GetBody() *UpgradeVersionByUuidsResponseBody {
	return s.Body
}

func (s *UpgradeVersionByUuidsResponse) SetHeaders(v map[string]*string) *UpgradeVersionByUuidsResponse {
	s.Headers = v
	return s
}

func (s *UpgradeVersionByUuidsResponse) SetStatusCode(v int32) *UpgradeVersionByUuidsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeVersionByUuidsResponse) SetBody(v *UpgradeVersionByUuidsResponseBody) *UpgradeVersionByUuidsResponse {
	s.Body = v
	return s
}

func (s *UpgradeVersionByUuidsResponse) Validate() error {
	return dara.Validate(s)
}
