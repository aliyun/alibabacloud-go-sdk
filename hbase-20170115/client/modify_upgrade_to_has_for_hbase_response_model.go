// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUpgradeToHasForHbaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUpgradeToHasForHbaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUpgradeToHasForHbaseResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUpgradeToHasForHbaseResponseBody) *ModifyUpgradeToHasForHbaseResponse
	GetBody() *ModifyUpgradeToHasForHbaseResponseBody
}

type ModifyUpgradeToHasForHbaseResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUpgradeToHasForHbaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUpgradeToHasForHbaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUpgradeToHasForHbaseResponse) GoString() string {
	return s.String()
}

func (s *ModifyUpgradeToHasForHbaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUpgradeToHasForHbaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUpgradeToHasForHbaseResponse) GetBody() *ModifyUpgradeToHasForHbaseResponseBody {
	return s.Body
}

func (s *ModifyUpgradeToHasForHbaseResponse) SetHeaders(v map[string]*string) *ModifyUpgradeToHasForHbaseResponse {
	s.Headers = v
	return s
}

func (s *ModifyUpgradeToHasForHbaseResponse) SetStatusCode(v int32) *ModifyUpgradeToHasForHbaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseResponse) SetBody(v *ModifyUpgradeToHasForHbaseResponseBody) *ModifyUpgradeToHasForHbaseResponse {
	s.Body = v
	return s
}

func (s *ModifyUpgradeToHasForHbaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
