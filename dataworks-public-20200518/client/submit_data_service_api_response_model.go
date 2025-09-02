// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDataServiceApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDataServiceApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDataServiceApiResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDataServiceApiResponseBody) *SubmitDataServiceApiResponse
	GetBody() *SubmitDataServiceApiResponseBody
}

type SubmitDataServiceApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDataServiceApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDataServiceApiResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDataServiceApiResponse) GoString() string {
	return s.String()
}

func (s *SubmitDataServiceApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDataServiceApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDataServiceApiResponse) GetBody() *SubmitDataServiceApiResponseBody {
	return s.Body
}

func (s *SubmitDataServiceApiResponse) SetHeaders(v map[string]*string) *SubmitDataServiceApiResponse {
	s.Headers = v
	return s
}

func (s *SubmitDataServiceApiResponse) SetStatusCode(v int32) *SubmitDataServiceApiResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDataServiceApiResponse) SetBody(v *SubmitDataServiceApiResponseBody) *SubmitDataServiceApiResponse {
	s.Body = v
	return s
}

func (s *SubmitDataServiceApiResponse) Validate() error {
	return dara.Validate(s)
}
