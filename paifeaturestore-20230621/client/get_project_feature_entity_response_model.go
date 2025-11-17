// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectFeatureEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectFeatureEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectFeatureEntityResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectFeatureEntityResponseBody) *GetProjectFeatureEntityResponse
	GetBody() *GetProjectFeatureEntityResponseBody
}

type GetProjectFeatureEntityResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectFeatureEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectFeatureEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectFeatureEntityResponse) GoString() string {
	return s.String()
}

func (s *GetProjectFeatureEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectFeatureEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectFeatureEntityResponse) GetBody() *GetProjectFeatureEntityResponseBody {
	return s.Body
}

func (s *GetProjectFeatureEntityResponse) SetHeaders(v map[string]*string) *GetProjectFeatureEntityResponse {
	s.Headers = v
	return s
}

func (s *GetProjectFeatureEntityResponse) SetStatusCode(v int32) *GetProjectFeatureEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectFeatureEntityResponse) SetBody(v *GetProjectFeatureEntityResponseBody) *GetProjectFeatureEntityResponse {
	s.Body = v
	return s
}

func (s *GetProjectFeatureEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
