// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRbacOrgUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRbacOrgUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRbacOrgUnitResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRbacOrgUnitResponseBody) *DeleteRbacOrgUnitResponse
	GetBody() *DeleteRbacOrgUnitResponseBody
}

type DeleteRbacOrgUnitResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRbacOrgUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRbacOrgUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRbacOrgUnitResponse) GoString() string {
	return s.String()
}

func (s *DeleteRbacOrgUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRbacOrgUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRbacOrgUnitResponse) GetBody() *DeleteRbacOrgUnitResponseBody {
	return s.Body
}

func (s *DeleteRbacOrgUnitResponse) SetHeaders(v map[string]*string) *DeleteRbacOrgUnitResponse {
	s.Headers = v
	return s
}

func (s *DeleteRbacOrgUnitResponse) SetStatusCode(v int32) *DeleteRbacOrgUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRbacOrgUnitResponse) SetBody(v *DeleteRbacOrgUnitResponseBody) *DeleteRbacOrgUnitResponse {
	s.Body = v
	return s
}

func (s *DeleteRbacOrgUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
