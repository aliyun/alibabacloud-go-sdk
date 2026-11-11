// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAccountResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudAccountResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudAccountResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudAccountResourceResponseBody) *GetCloudAccountResourceResponse
	GetBody() *GetCloudAccountResourceResponseBody
}

type GetCloudAccountResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudAccountResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudAccountResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountResourceResponse) GoString() string {
	return s.String()
}

func (s *GetCloudAccountResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudAccountResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudAccountResourceResponse) GetBody() *GetCloudAccountResourceResponseBody {
	return s.Body
}

func (s *GetCloudAccountResourceResponse) SetHeaders(v map[string]*string) *GetCloudAccountResourceResponse {
	s.Headers = v
	return s
}

func (s *GetCloudAccountResourceResponse) SetStatusCode(v int32) *GetCloudAccountResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudAccountResourceResponse) SetBody(v *GetCloudAccountResourceResponseBody) *GetCloudAccountResourceResponse {
	s.Body = v
	return s
}

func (s *GetCloudAccountResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
