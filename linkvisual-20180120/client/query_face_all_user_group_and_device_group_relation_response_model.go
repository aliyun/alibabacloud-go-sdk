// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllUserGroupAndDeviceGroupRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) *QueryFaceAllUserGroupAndDeviceGroupRelationResponse
	GetBody() *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponse) GetBody() *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	return s.Body
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponse) SetHeaders(v map[string]*string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponse) SetStatusCode(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponse) SetBody(v *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) *QueryFaceAllUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
