// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAndroidInstancesInGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAndroidInstancesInGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAndroidInstancesInGroupResponse
	GetStatusCode() *int32
	SetBody(v *ResetAndroidInstancesInGroupResponseBody) *ResetAndroidInstancesInGroupResponse
	GetBody() *ResetAndroidInstancesInGroupResponseBody
}

type ResetAndroidInstancesInGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAndroidInstancesInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAndroidInstancesInGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAndroidInstancesInGroupResponse) GoString() string {
	return s.String()
}

func (s *ResetAndroidInstancesInGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAndroidInstancesInGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAndroidInstancesInGroupResponse) GetBody() *ResetAndroidInstancesInGroupResponseBody {
	return s.Body
}

func (s *ResetAndroidInstancesInGroupResponse) SetHeaders(v map[string]*string) *ResetAndroidInstancesInGroupResponse {
	s.Headers = v
	return s
}

func (s *ResetAndroidInstancesInGroupResponse) SetStatusCode(v int32) *ResetAndroidInstancesInGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAndroidInstancesInGroupResponse) SetBody(v *ResetAndroidInstancesInGroupResponseBody) *ResetAndroidInstancesInGroupResponse {
	s.Body = v
	return s
}

func (s *ResetAndroidInstancesInGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
