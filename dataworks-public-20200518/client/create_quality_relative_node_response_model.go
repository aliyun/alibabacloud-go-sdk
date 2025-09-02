// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityRelativeNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQualityRelativeNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQualityRelativeNodeResponse
	GetStatusCode() *int32
	SetBody(v *CreateQualityRelativeNodeResponseBody) *CreateQualityRelativeNodeResponse
	GetBody() *CreateQualityRelativeNodeResponseBody
}

type CreateQualityRelativeNodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQualityRelativeNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQualityRelativeNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityRelativeNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateQualityRelativeNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQualityRelativeNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQualityRelativeNodeResponse) GetBody() *CreateQualityRelativeNodeResponseBody {
	return s.Body
}

func (s *CreateQualityRelativeNodeResponse) SetHeaders(v map[string]*string) *CreateQualityRelativeNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateQualityRelativeNodeResponse) SetStatusCode(v int32) *CreateQualityRelativeNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQualityRelativeNodeResponse) SetBody(v *CreateQualityRelativeNodeResponseBody) *CreateQualityRelativeNodeResponse {
	s.Body = v
	return s
}

func (s *CreateQualityRelativeNodeResponse) Validate() error {
	return dara.Validate(s)
}
