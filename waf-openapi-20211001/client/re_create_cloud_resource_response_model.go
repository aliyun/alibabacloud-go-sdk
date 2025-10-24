// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReCreateCloudResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReCreateCloudResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReCreateCloudResourceResponse
	GetStatusCode() *int32
	SetBody(v *ReCreateCloudResourceResponseBody) *ReCreateCloudResourceResponse
	GetBody() *ReCreateCloudResourceResponseBody
}

type ReCreateCloudResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReCreateCloudResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReCreateCloudResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReCreateCloudResourceResponse) GoString() string {
	return s.String()
}

func (s *ReCreateCloudResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReCreateCloudResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReCreateCloudResourceResponse) GetBody() *ReCreateCloudResourceResponseBody {
	return s.Body
}

func (s *ReCreateCloudResourceResponse) SetHeaders(v map[string]*string) *ReCreateCloudResourceResponse {
	s.Headers = v
	return s
}

func (s *ReCreateCloudResourceResponse) SetStatusCode(v int32) *ReCreateCloudResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReCreateCloudResourceResponse) SetBody(v *ReCreateCloudResourceResponseBody) *ReCreateCloudResourceResponse {
	s.Body = v
	return s
}

func (s *ReCreateCloudResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
