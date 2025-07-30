// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReverseDtsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReverseDtsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReverseDtsJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateReverseDtsJobResponseBody) *CreateReverseDtsJobResponse
	GetBody() *CreateReverseDtsJobResponseBody
}

type CreateReverseDtsJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReverseDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReverseDtsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReverseDtsJobResponse) GoString() string {
	return s.String()
}

func (s *CreateReverseDtsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReverseDtsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReverseDtsJobResponse) GetBody() *CreateReverseDtsJobResponseBody {
	return s.Body
}

func (s *CreateReverseDtsJobResponse) SetHeaders(v map[string]*string) *CreateReverseDtsJobResponse {
	s.Headers = v
	return s
}

func (s *CreateReverseDtsJobResponse) SetStatusCode(v int32) *CreateReverseDtsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReverseDtsJobResponse) SetBody(v *CreateReverseDtsJobResponseBody) *CreateReverseDtsJobResponse {
	s.Body = v
	return s
}

func (s *CreateReverseDtsJobResponse) Validate() error {
	return dara.Validate(s)
}
