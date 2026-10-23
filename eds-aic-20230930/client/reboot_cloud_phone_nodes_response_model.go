// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootCloudPhoneNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootCloudPhoneNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootCloudPhoneNodesResponse
	GetStatusCode() *int32
	SetBody(v *RebootCloudPhoneNodesResponseBody) *RebootCloudPhoneNodesResponse
	GetBody() *RebootCloudPhoneNodesResponseBody
}

type RebootCloudPhoneNodesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootCloudPhoneNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootCloudPhoneNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootCloudPhoneNodesResponse) GoString() string {
	return s.String()
}

func (s *RebootCloudPhoneNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootCloudPhoneNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootCloudPhoneNodesResponse) GetBody() *RebootCloudPhoneNodesResponseBody {
	return s.Body
}

func (s *RebootCloudPhoneNodesResponse) SetHeaders(v map[string]*string) *RebootCloudPhoneNodesResponse {
	s.Headers = v
	return s
}

func (s *RebootCloudPhoneNodesResponse) SetStatusCode(v int32) *RebootCloudPhoneNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootCloudPhoneNodesResponse) SetBody(v *RebootCloudPhoneNodesResponseBody) *RebootCloudPhoneNodesResponse {
	s.Body = v
	return s
}

func (s *RebootCloudPhoneNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
