// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOfficeSiteAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOfficeSiteAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOfficeSiteAcceleratorResponseBody) *ModifyOfficeSiteAcceleratorResponse
	GetBody() *ModifyOfficeSiteAcceleratorResponseBody
}

type ModifyOfficeSiteAcceleratorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOfficeSiteAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOfficeSiteAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOfficeSiteAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOfficeSiteAcceleratorResponse) GetBody() *ModifyOfficeSiteAcceleratorResponseBody {
	return s.Body
}

func (s *ModifyOfficeSiteAcceleratorResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteAcceleratorResponse) SetStatusCode(v int32) *ModifyOfficeSiteAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOfficeSiteAcceleratorResponse) SetBody(v *ModifyOfficeSiteAcceleratorResponseBody) *ModifyOfficeSiteAcceleratorResponse {
	s.Body = v
	return s
}

func (s *ModifyOfficeSiteAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
