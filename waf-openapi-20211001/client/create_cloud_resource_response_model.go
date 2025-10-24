// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudResourceResponseBody) *CreateCloudResourceResponse
	GetBody() *CreateCloudResourceResponseBody
}

type CreateCloudResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudResourceResponse) GetBody() *CreateCloudResourceResponseBody {
	return s.Body
}

func (s *CreateCloudResourceResponse) SetHeaders(v map[string]*string) *CreateCloudResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudResourceResponse) SetStatusCode(v int32) *CreateCloudResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudResourceResponse) SetBody(v *CreateCloudResourceResponseBody) *CreateCloudResourceResponse {
	s.Body = v
	return s
}

func (s *CreateCloudResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
