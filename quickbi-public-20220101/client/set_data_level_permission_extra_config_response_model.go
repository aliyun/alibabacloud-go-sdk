// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataLevelPermissionExtraConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDataLevelPermissionExtraConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDataLevelPermissionExtraConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetDataLevelPermissionExtraConfigResponseBody) *SetDataLevelPermissionExtraConfigResponse
	GetBody() *SetDataLevelPermissionExtraConfigResponseBody
}

type SetDataLevelPermissionExtraConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDataLevelPermissionExtraConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDataLevelPermissionExtraConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDataLevelPermissionExtraConfigResponse) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionExtraConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDataLevelPermissionExtraConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDataLevelPermissionExtraConfigResponse) GetBody() *SetDataLevelPermissionExtraConfigResponseBody {
	return s.Body
}

func (s *SetDataLevelPermissionExtraConfigResponse) SetHeaders(v map[string]*string) *SetDataLevelPermissionExtraConfigResponse {
	s.Headers = v
	return s
}

func (s *SetDataLevelPermissionExtraConfigResponse) SetStatusCode(v int32) *SetDataLevelPermissionExtraConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigResponse) SetBody(v *SetDataLevelPermissionExtraConfigResponseBody) *SetDataLevelPermissionExtraConfigResponse {
	s.Body = v
	return s
}

func (s *SetDataLevelPermissionExtraConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
