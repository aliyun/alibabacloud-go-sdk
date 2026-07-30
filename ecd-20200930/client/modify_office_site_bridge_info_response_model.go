// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteBridgeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOfficeSiteBridgeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOfficeSiteBridgeInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOfficeSiteBridgeInfoResponseBody) *ModifyOfficeSiteBridgeInfoResponse
	GetBody() *ModifyOfficeSiteBridgeInfoResponseBody
}

type ModifyOfficeSiteBridgeInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOfficeSiteBridgeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOfficeSiteBridgeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteBridgeInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteBridgeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOfficeSiteBridgeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOfficeSiteBridgeInfoResponse) GetBody() *ModifyOfficeSiteBridgeInfoResponseBody {
	return s.Body
}

func (s *ModifyOfficeSiteBridgeInfoResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteBridgeInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteBridgeInfoResponse) SetStatusCode(v int32) *ModifyOfficeSiteBridgeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOfficeSiteBridgeInfoResponse) SetBody(v *ModifyOfficeSiteBridgeInfoResponseBody) *ModifyOfficeSiteBridgeInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyOfficeSiteBridgeInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
