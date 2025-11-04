// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivePackageChannelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLivePackageChannelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLivePackageChannelGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetLivePackageChannelGroupResponseBody) *GetLivePackageChannelGroupResponse
	GetBody() *GetLivePackageChannelGroupResponseBody
}

type GetLivePackageChannelGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLivePackageChannelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLivePackageChannelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageChannelGroupResponse) GoString() string {
	return s.String()
}

func (s *GetLivePackageChannelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLivePackageChannelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLivePackageChannelGroupResponse) GetBody() *GetLivePackageChannelGroupResponseBody {
	return s.Body
}

func (s *GetLivePackageChannelGroupResponse) SetHeaders(v map[string]*string) *GetLivePackageChannelGroupResponse {
	s.Headers = v
	return s
}

func (s *GetLivePackageChannelGroupResponse) SetStatusCode(v int32) *GetLivePackageChannelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLivePackageChannelGroupResponse) SetBody(v *GetLivePackageChannelGroupResponseBody) *GetLivePackageChannelGroupResponse {
	s.Body = v
	return s
}

func (s *GetLivePackageChannelGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
