// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceUserGroupAndDeviceGroupRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceUserGroupAndDeviceGroupRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceUserGroupAndDeviceGroupRelationResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceUserGroupAndDeviceGroupRelationResponseBody) *AddFaceUserGroupAndDeviceGroupRelationResponse
	GetBody() *AddFaceUserGroupAndDeviceGroupRelationResponseBody
}

type AddFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponse) GetBody() *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	return s.Body
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponse) SetHeaders(v map[string]*string) *AddFaceUserGroupAndDeviceGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponse) SetStatusCode(v int32) *AddFaceUserGroupAndDeviceGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *AddFaceUserGroupAndDeviceGroupRelationResponseBody) *AddFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
