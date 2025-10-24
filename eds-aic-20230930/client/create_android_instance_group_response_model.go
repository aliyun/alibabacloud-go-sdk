// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndroidInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAndroidInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAndroidInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateAndroidInstanceGroupResponseBody) *CreateAndroidInstanceGroupResponse
	GetBody() *CreateAndroidInstanceGroupResponseBody
}

type CreateAndroidInstanceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndroidInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndroidInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAndroidInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAndroidInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAndroidInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAndroidInstanceGroupResponse) GetBody() *CreateAndroidInstanceGroupResponseBody {
	return s.Body
}

func (s *CreateAndroidInstanceGroupResponse) SetHeaders(v map[string]*string) *CreateAndroidInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAndroidInstanceGroupResponse) SetStatusCode(v int32) *CreateAndroidInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndroidInstanceGroupResponse) SetBody(v *CreateAndroidInstanceGroupResponseBody) *CreateAndroidInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *CreateAndroidInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
