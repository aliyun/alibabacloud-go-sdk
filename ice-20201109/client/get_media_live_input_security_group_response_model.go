// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLiveInputSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaLiveInputSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaLiveInputSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaLiveInputSecurityGroupResponseBody) *GetMediaLiveInputSecurityGroupResponse
	GetBody() *GetMediaLiveInputSecurityGroupResponseBody
}

type GetMediaLiveInputSecurityGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaLiveInputSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaLiveInputSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveInputSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *GetMediaLiveInputSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaLiveInputSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaLiveInputSecurityGroupResponse) GetBody() *GetMediaLiveInputSecurityGroupResponseBody {
	return s.Body
}

func (s *GetMediaLiveInputSecurityGroupResponse) SetHeaders(v map[string]*string) *GetMediaLiveInputSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *GetMediaLiveInputSecurityGroupResponse) SetStatusCode(v int32) *GetMediaLiveInputSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaLiveInputSecurityGroupResponse) SetBody(v *GetMediaLiveInputSecurityGroupResponseBody) *GetMediaLiveInputSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *GetMediaLiveInputSecurityGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
