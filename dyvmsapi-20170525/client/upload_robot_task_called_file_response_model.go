// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRobotTaskCalledFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadRobotTaskCalledFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadRobotTaskCalledFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadRobotTaskCalledFileResponseBody) *UploadRobotTaskCalledFileResponse
	GetBody() *UploadRobotTaskCalledFileResponseBody
}

type UploadRobotTaskCalledFileResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadRobotTaskCalledFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadRobotTaskCalledFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadRobotTaskCalledFileResponse) GoString() string {
	return s.String()
}

func (s *UploadRobotTaskCalledFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadRobotTaskCalledFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadRobotTaskCalledFileResponse) GetBody() *UploadRobotTaskCalledFileResponseBody {
	return s.Body
}

func (s *UploadRobotTaskCalledFileResponse) SetHeaders(v map[string]*string) *UploadRobotTaskCalledFileResponse {
	s.Headers = v
	return s
}

func (s *UploadRobotTaskCalledFileResponse) SetStatusCode(v int32) *UploadRobotTaskCalledFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadRobotTaskCalledFileResponse) SetBody(v *UploadRobotTaskCalledFileResponseBody) *UploadRobotTaskCalledFileResponse {
	s.Body = v
	return s
}

func (s *UploadRobotTaskCalledFileResponse) Validate() error {
	return dara.Validate(s)
}
