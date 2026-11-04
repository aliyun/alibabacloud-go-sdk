// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaceUserGroupAndDeviceGroupRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFaceUserGroupAndDeviceGroupRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFaceUserGroupAndDeviceGroupRelationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) *UpdateFaceUserGroupAndDeviceGroupRelationResponse
	GetBody() *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody
}

type UpdateFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponse) GetBody() *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody {
	return s.Body
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponse) SetHeaders(v map[string]*string) *UpdateFaceUserGroupAndDeviceGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponse) SetStatusCode(v int32) *UpdateFaceUserGroupAndDeviceGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *UpdateFaceUserGroupAndDeviceGroupRelationResponseBody) *UpdateFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
