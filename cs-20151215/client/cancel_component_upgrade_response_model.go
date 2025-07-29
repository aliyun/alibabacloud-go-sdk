// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelComponentUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelComponentUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelComponentUpgradeResponse
	GetStatusCode() *int32
}

type CancelComponentUpgradeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CancelComponentUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *CancelComponentUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelComponentUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelComponentUpgradeResponse) SetHeaders(v map[string]*string) *CancelComponentUpgradeResponse {
	s.Headers = v
	return s
}

func (s *CancelComponentUpgradeResponse) SetStatusCode(v int32) *CancelComponentUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelComponentUpgradeResponse) Validate() error {
	return dara.Validate(s)
}
