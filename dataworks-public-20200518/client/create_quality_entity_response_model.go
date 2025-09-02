// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQualityEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQualityEntityResponse
	GetStatusCode() *int32
	SetBody(v *CreateQualityEntityResponseBody) *CreateQualityEntityResponse
	GetBody() *CreateQualityEntityResponseBody
}

type CreateQualityEntityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQualityEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQualityEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateQualityEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQualityEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQualityEntityResponse) GetBody() *CreateQualityEntityResponseBody {
	return s.Body
}

func (s *CreateQualityEntityResponse) SetHeaders(v map[string]*string) *CreateQualityEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateQualityEntityResponse) SetStatusCode(v int32) *CreateQualityEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQualityEntityResponse) SetBody(v *CreateQualityEntityResponseBody) *CreateQualityEntityResponse {
	s.Body = v
	return s
}

func (s *CreateQualityEntityResponse) Validate() error {
	return dara.Validate(s)
}
