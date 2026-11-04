// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFaceDeviceFromDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveFaceDeviceFromDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveFaceDeviceFromDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveFaceDeviceFromDeviceGroupResponseBody) *RemoveFaceDeviceFromDeviceGroupResponse
	GetBody() *RemoveFaceDeviceFromDeviceGroupResponseBody
}

type RemoveFaceDeviceFromDeviceGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveFaceDeviceFromDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveFaceDeviceFromDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveFaceDeviceFromDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveFaceDeviceFromDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveFaceDeviceFromDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveFaceDeviceFromDeviceGroupResponse) GetBody() *RemoveFaceDeviceFromDeviceGroupResponseBody {
	return s.Body
}

func (s *RemoveFaceDeviceFromDeviceGroupResponse) SetHeaders(v map[string]*string) *RemoveFaceDeviceFromDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponse) SetStatusCode(v int32) *RemoveFaceDeviceFromDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponse) SetBody(v *RemoveFaceDeviceFromDeviceGroupResponseBody) *RemoveFaceDeviceFromDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
