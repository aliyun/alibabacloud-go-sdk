// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudPhoneNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudPhoneNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudPhoneNodesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudPhoneNodesResponseBody) *DeleteCloudPhoneNodesResponse
	GetBody() *DeleteCloudPhoneNodesResponseBody
}

type DeleteCloudPhoneNodesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudPhoneNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudPhoneNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudPhoneNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudPhoneNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudPhoneNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudPhoneNodesResponse) GetBody() *DeleteCloudPhoneNodesResponseBody {
	return s.Body
}

func (s *DeleteCloudPhoneNodesResponse) SetHeaders(v map[string]*string) *DeleteCloudPhoneNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudPhoneNodesResponse) SetStatusCode(v int32) *DeleteCloudPhoneNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudPhoneNodesResponse) SetBody(v *DeleteCloudPhoneNodesResponseBody) *DeleteCloudPhoneNodesResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudPhoneNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
