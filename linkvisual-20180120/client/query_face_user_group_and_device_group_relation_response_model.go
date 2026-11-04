// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserGroupAndDeviceGroupRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceUserGroupAndDeviceGroupRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceUserGroupAndDeviceGroupRelationResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) *QueryFaceUserGroupAndDeviceGroupRelationResponse
	GetBody() *QueryFaceUserGroupAndDeviceGroupRelationResponseBody
}

type QueryFaceUserGroupAndDeviceGroupRelationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserGroupAndDeviceGroupRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponse) GetBody() *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	return s.Body
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponse) SetHeaders(v map[string]*string) *QueryFaceUserGroupAndDeviceGroupRelationResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponse) SetStatusCode(v int32) *QueryFaceUserGroupAndDeviceGroupRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponse) SetBody(v *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) *QueryFaceUserGroupAndDeviceGroupRelationResponse {
	s.Body = v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
