// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetAppInstanceGroupResponseBody) *GetAppInstanceGroupResponse
	GetBody() *GetAppInstanceGroupResponseBody
}

type GetAppInstanceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppInstanceGroupResponse) GetBody() *GetAppInstanceGroupResponseBody {
	return s.Body
}

func (s *GetAppInstanceGroupResponse) SetHeaders(v map[string]*string) *GetAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetAppInstanceGroupResponse) SetStatusCode(v int32) *GetAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppInstanceGroupResponse) SetBody(v *GetAppInstanceGroupResponseBody) *GetAppInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *GetAppInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
