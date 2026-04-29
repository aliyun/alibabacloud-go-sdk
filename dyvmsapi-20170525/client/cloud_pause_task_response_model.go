// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudPauseTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudPauseTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudPauseTaskResponse
	GetStatusCode() *int32
	SetBody(v *CloudPauseTaskResponseBody) *CloudPauseTaskResponse
	GetBody() *CloudPauseTaskResponseBody
}

type CloudPauseTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudPauseTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudPauseTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudPauseTaskResponse) GoString() string {
	return s.String()
}

func (s *CloudPauseTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudPauseTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudPauseTaskResponse) GetBody() *CloudPauseTaskResponseBody {
	return s.Body
}

func (s *CloudPauseTaskResponse) SetHeaders(v map[string]*string) *CloudPauseTaskResponse {
	s.Headers = v
	return s
}

func (s *CloudPauseTaskResponse) SetStatusCode(v int32) *CloudPauseTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudPauseTaskResponse) SetBody(v *CloudPauseTaskResponseBody) *CloudPauseTaskResponse {
	s.Body = v
	return s
}

func (s *CloudPauseTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
