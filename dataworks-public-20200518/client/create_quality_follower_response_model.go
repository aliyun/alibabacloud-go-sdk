// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityFollowerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQualityFollowerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQualityFollowerResponse
	GetStatusCode() *int32
	SetBody(v *CreateQualityFollowerResponseBody) *CreateQualityFollowerResponse
	GetBody() *CreateQualityFollowerResponseBody
}

type CreateQualityFollowerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQualityFollowerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQualityFollowerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityFollowerResponse) GoString() string {
	return s.String()
}

func (s *CreateQualityFollowerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQualityFollowerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQualityFollowerResponse) GetBody() *CreateQualityFollowerResponseBody {
	return s.Body
}

func (s *CreateQualityFollowerResponse) SetHeaders(v map[string]*string) *CreateQualityFollowerResponse {
	s.Headers = v
	return s
}

func (s *CreateQualityFollowerResponse) SetStatusCode(v int32) *CreateQualityFollowerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQualityFollowerResponse) SetBody(v *CreateQualityFollowerResponseBody) *CreateQualityFollowerResponse {
	s.Body = v
	return s
}

func (s *CreateQualityFollowerResponse) Validate() error {
	return dara.Validate(s)
}
