// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityCheckSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQualityCheckSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQualityCheckSchemeResponse
	GetStatusCode() *int32
	SetBody(v *CreateQualityCheckSchemeResponseBody) *CreateQualityCheckSchemeResponse
	GetBody() *CreateQualityCheckSchemeResponseBody
}

type CreateQualityCheckSchemeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQualityCheckSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQualityCheckSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityCheckSchemeResponse) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQualityCheckSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQualityCheckSchemeResponse) GetBody() *CreateQualityCheckSchemeResponseBody {
	return s.Body
}

func (s *CreateQualityCheckSchemeResponse) SetHeaders(v map[string]*string) *CreateQualityCheckSchemeResponse {
	s.Headers = v
	return s
}

func (s *CreateQualityCheckSchemeResponse) SetStatusCode(v int32) *CreateQualityCheckSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQualityCheckSchemeResponse) SetBody(v *CreateQualityCheckSchemeResponseBody) *CreateQualityCheckSchemeResponse {
	s.Body = v
	return s
}

func (s *CreateQualityCheckSchemeResponse) Validate() error {
	return dara.Validate(s)
}
