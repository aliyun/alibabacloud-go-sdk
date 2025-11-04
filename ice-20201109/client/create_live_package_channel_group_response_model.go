// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePackageChannelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLivePackageChannelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLivePackageChannelGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateLivePackageChannelGroupResponseBody) *CreateLivePackageChannelGroupResponse
	GetBody() *CreateLivePackageChannelGroupResponseBody
}

type CreateLivePackageChannelGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLivePackageChannelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLivePackageChannelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageChannelGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateLivePackageChannelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLivePackageChannelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLivePackageChannelGroupResponse) GetBody() *CreateLivePackageChannelGroupResponseBody {
	return s.Body
}

func (s *CreateLivePackageChannelGroupResponse) SetHeaders(v map[string]*string) *CreateLivePackageChannelGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateLivePackageChannelGroupResponse) SetStatusCode(v int32) *CreateLivePackageChannelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLivePackageChannelGroupResponse) SetBody(v *CreateLivePackageChannelGroupResponseBody) *CreateLivePackageChannelGroupResponse {
	s.Body = v
	return s
}

func (s *CreateLivePackageChannelGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
