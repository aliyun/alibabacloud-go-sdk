// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFeatureEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFeatureEntityResponse
	GetStatusCode() *int32
	SetBody(v *GetFeatureEntityResponseBody) *GetFeatureEntityResponse
	GetBody() *GetFeatureEntityResponseBody
}

type GetFeatureEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureEntityResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFeatureEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFeatureEntityResponse) GetBody() *GetFeatureEntityResponseBody {
	return s.Body
}

func (s *GetFeatureEntityResponse) SetHeaders(v map[string]*string) *GetFeatureEntityResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureEntityResponse) SetStatusCode(v int32) *GetFeatureEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureEntityResponse) SetBody(v *GetFeatureEntityResponseBody) *GetFeatureEntityResponse {
	s.Body = v
	return s
}

func (s *GetFeatureEntityResponse) Validate() error {
	return dara.Validate(s)
}
