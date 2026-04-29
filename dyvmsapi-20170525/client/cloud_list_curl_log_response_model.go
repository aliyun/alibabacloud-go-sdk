// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListCurlLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListCurlLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListCurlLogResponse
	GetStatusCode() *int32
	SetBody(v *CloudListCurlLogResponseBody) *CloudListCurlLogResponse
	GetBody() *CloudListCurlLogResponseBody
}

type CloudListCurlLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListCurlLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListCurlLogResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListCurlLogResponse) GoString() string {
	return s.String()
}

func (s *CloudListCurlLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListCurlLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListCurlLogResponse) GetBody() *CloudListCurlLogResponseBody {
	return s.Body
}

func (s *CloudListCurlLogResponse) SetHeaders(v map[string]*string) *CloudListCurlLogResponse {
	s.Headers = v
	return s
}

func (s *CloudListCurlLogResponse) SetStatusCode(v int32) *CloudListCurlLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListCurlLogResponse) SetBody(v *CloudListCurlLogResponseBody) *CloudListCurlLogResponse {
	s.Body = v
	return s
}

func (s *CloudListCurlLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
