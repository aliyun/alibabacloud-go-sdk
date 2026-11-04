// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceUserGroupAndDeviceGroupRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceUserGroupAndDeviceGroupRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceUserGroupAndDeviceGroupRelationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) *DeleteFaceUserGroupAndDeviceGroupRelationResponse
	GetBody() *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody
}

type DeleteFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponse) GetBody() *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody {
	return s.Body
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponse) SetHeaders(v map[string]*string) *DeleteFaceUserGroupAndDeviceGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponse) SetStatusCode(v int32) *DeleteFaceUserGroupAndDeviceGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *DeleteFaceUserGroupAndDeviceGroupRelationResponseBody) *DeleteFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
