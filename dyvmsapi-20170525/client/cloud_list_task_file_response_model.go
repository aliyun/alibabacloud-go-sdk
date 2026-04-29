// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListTaskFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListTaskFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListTaskFileResponse
	GetStatusCode() *int32
	SetBody(v *CloudListTaskFileResponseBody) *CloudListTaskFileResponse
	GetBody() *CloudListTaskFileResponseBody
}

type CloudListTaskFileResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListTaskFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListTaskFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListTaskFileResponse) GoString() string {
	return s.String()
}

func (s *CloudListTaskFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListTaskFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListTaskFileResponse) GetBody() *CloudListTaskFileResponseBody {
	return s.Body
}

func (s *CloudListTaskFileResponse) SetHeaders(v map[string]*string) *CloudListTaskFileResponse {
	s.Headers = v
	return s
}

func (s *CloudListTaskFileResponse) SetStatusCode(v int32) *CloudListTaskFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListTaskFileResponse) SetBody(v *CloudListTaskFileResponseBody) *CloudListTaskFileResponse {
	s.Body = v
	return s
}

func (s *CloudListTaskFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
