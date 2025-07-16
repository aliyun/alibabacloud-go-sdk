// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFeatureViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFeatureViewResponse
	GetStatusCode() *int32
	SetBody(v *CreateFeatureViewResponseBody) *CreateFeatureViewResponse
	GetBody() *CreateFeatureViewResponseBody
}

type CreateFeatureViewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFeatureViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFeatureViewResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureViewResponse) GoString() string {
	return s.String()
}

func (s *CreateFeatureViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFeatureViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFeatureViewResponse) GetBody() *CreateFeatureViewResponseBody {
	return s.Body
}

func (s *CreateFeatureViewResponse) SetHeaders(v map[string]*string) *CreateFeatureViewResponse {
	s.Headers = v
	return s
}

func (s *CreateFeatureViewResponse) SetStatusCode(v int32) *CreateFeatureViewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFeatureViewResponse) SetBody(v *CreateFeatureViewResponseBody) *CreateFeatureViewResponse {
	s.Body = v
	return s
}

func (s *CreateFeatureViewResponse) Validate() error {
	return dara.Validate(s)
}
