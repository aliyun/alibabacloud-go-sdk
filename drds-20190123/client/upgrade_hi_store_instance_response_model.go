// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeHiStoreInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeHiStoreInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeHiStoreInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeHiStoreInstanceResponseBody) *UpgradeHiStoreInstanceResponse
	GetBody() *UpgradeHiStoreInstanceResponseBody
}

type UpgradeHiStoreInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeHiStoreInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeHiStoreInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeHiStoreInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeHiStoreInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeHiStoreInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeHiStoreInstanceResponse) GetBody() *UpgradeHiStoreInstanceResponseBody {
	return s.Body
}

func (s *UpgradeHiStoreInstanceResponse) SetHeaders(v map[string]*string) *UpgradeHiStoreInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeHiStoreInstanceResponse) SetStatusCode(v int32) *UpgradeHiStoreInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeHiStoreInstanceResponse) SetBody(v *UpgradeHiStoreInstanceResponseBody) *UpgradeHiStoreInstanceResponse {
	s.Body = v
	return s
}

func (s *UpgradeHiStoreInstanceResponse) Validate() error {
	return dara.Validate(s)
}
