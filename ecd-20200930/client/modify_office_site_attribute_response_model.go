// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOfficeSiteAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOfficeSiteAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOfficeSiteAttributeResponseBody) *ModifyOfficeSiteAttributeResponse
	GetBody() *ModifyOfficeSiteAttributeResponseBody
}

type ModifyOfficeSiteAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOfficeSiteAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOfficeSiteAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOfficeSiteAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOfficeSiteAttributeResponse) GetBody() *ModifyOfficeSiteAttributeResponseBody {
	return s.Body
}

func (s *ModifyOfficeSiteAttributeResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteAttributeResponse) SetStatusCode(v int32) *ModifyOfficeSiteAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOfficeSiteAttributeResponse) SetBody(v *ModifyOfficeSiteAttributeResponseBody) *ModifyOfficeSiteAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyOfficeSiteAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
