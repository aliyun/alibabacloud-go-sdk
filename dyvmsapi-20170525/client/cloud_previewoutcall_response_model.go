// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudPreviewoutcallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudPreviewoutcallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudPreviewoutcallResponse
	GetStatusCode() *int32
	SetBody(v *CloudPreviewoutcallResponseBody) *CloudPreviewoutcallResponse
	GetBody() *CloudPreviewoutcallResponseBody
}

type CloudPreviewoutcallResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudPreviewoutcallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudPreviewoutcallResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudPreviewoutcallResponse) GoString() string {
	return s.String()
}

func (s *CloudPreviewoutcallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudPreviewoutcallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudPreviewoutcallResponse) GetBody() *CloudPreviewoutcallResponseBody {
	return s.Body
}

func (s *CloudPreviewoutcallResponse) SetHeaders(v map[string]*string) *CloudPreviewoutcallResponse {
	s.Headers = v
	return s
}

func (s *CloudPreviewoutcallResponse) SetStatusCode(v int32) *CloudPreviewoutcallResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudPreviewoutcallResponse) SetBody(v *CloudPreviewoutcallResponseBody) *CloudPreviewoutcallResponse {
	s.Body = v
	return s
}

func (s *CloudPreviewoutcallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
