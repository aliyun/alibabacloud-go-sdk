// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudResourceResponseBody) *GetCloudResourceResponse
	GetBody() *GetCloudResourceResponseBody
}

type GetCloudResourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudResourceResponse) GoString() string {
	return s.String()
}

func (s *GetCloudResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudResourceResponse) GetBody() *GetCloudResourceResponseBody {
	return s.Body
}

func (s *GetCloudResourceResponse) SetHeaders(v map[string]*string) *GetCloudResourceResponse {
	s.Headers = v
	return s
}

func (s *GetCloudResourceResponse) SetStatusCode(v int32) *GetCloudResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudResourceResponse) SetBody(v *GetCloudResourceResponseBody) *GetCloudResourceResponse {
	s.Body = v
	return s
}

func (s *GetCloudResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
