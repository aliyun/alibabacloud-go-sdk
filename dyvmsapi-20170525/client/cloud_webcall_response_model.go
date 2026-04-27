// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudWebcallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudWebcallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudWebcallResponse
	GetStatusCode() *int32
	SetBody(v *CloudWebcallResponseBody) *CloudWebcallResponse
	GetBody() *CloudWebcallResponseBody
}

type CloudWebcallResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudWebcallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudWebcallResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudWebcallResponse) GoString() string {
	return s.String()
}

func (s *CloudWebcallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudWebcallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudWebcallResponse) GetBody() *CloudWebcallResponseBody {
	return s.Body
}

func (s *CloudWebcallResponse) SetHeaders(v map[string]*string) *CloudWebcallResponse {
	s.Headers = v
	return s
}

func (s *CloudWebcallResponse) SetStatusCode(v int32) *CloudWebcallResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudWebcallResponse) SetBody(v *CloudWebcallResponseBody) *CloudWebcallResponse {
	s.Body = v
	return s
}

func (s *CloudWebcallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
