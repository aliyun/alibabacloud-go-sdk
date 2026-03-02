// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOfficeSiteAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOfficeSiteAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOfficeSiteAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOfficeSiteAcceleratorResponseBody) *DeleteOfficeSiteAcceleratorResponse
	GetBody() *DeleteOfficeSiteAcceleratorResponseBody
}

type DeleteOfficeSiteAcceleratorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOfficeSiteAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOfficeSiteAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOfficeSiteAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSiteAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOfficeSiteAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOfficeSiteAcceleratorResponse) GetBody() *DeleteOfficeSiteAcceleratorResponseBody {
	return s.Body
}

func (s *DeleteOfficeSiteAcceleratorResponse) SetHeaders(v map[string]*string) *DeleteOfficeSiteAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *DeleteOfficeSiteAcceleratorResponse) SetStatusCode(v int32) *DeleteOfficeSiteAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOfficeSiteAcceleratorResponse) SetBody(v *DeleteOfficeSiteAcceleratorResponseBody) *DeleteOfficeSiteAcceleratorResponse {
	s.Body = v
	return s
}

func (s *DeleteOfficeSiteAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
