// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootAndroidInstancesInGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootAndroidInstancesInGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootAndroidInstancesInGroupResponse
	GetStatusCode() *int32
	SetBody(v *RebootAndroidInstancesInGroupResponseBody) *RebootAndroidInstancesInGroupResponse
	GetBody() *RebootAndroidInstancesInGroupResponseBody
}

type RebootAndroidInstancesInGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootAndroidInstancesInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootAndroidInstancesInGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootAndroidInstancesInGroupResponse) GoString() string {
	return s.String()
}

func (s *RebootAndroidInstancesInGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootAndroidInstancesInGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootAndroidInstancesInGroupResponse) GetBody() *RebootAndroidInstancesInGroupResponseBody {
	return s.Body
}

func (s *RebootAndroidInstancesInGroupResponse) SetHeaders(v map[string]*string) *RebootAndroidInstancesInGroupResponse {
	s.Headers = v
	return s
}

func (s *RebootAndroidInstancesInGroupResponse) SetStatusCode(v int32) *RebootAndroidInstancesInGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootAndroidInstancesInGroupResponse) SetBody(v *RebootAndroidInstancesInGroupResponseBody) *RebootAndroidInstancesInGroupResponse {
	s.Body = v
	return s
}

func (s *RebootAndroidInstancesInGroupResponse) Validate() error {
	return dara.Validate(s)
}
