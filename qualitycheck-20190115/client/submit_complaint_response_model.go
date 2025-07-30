// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitComplaintResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitComplaintResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitComplaintResponse
	GetStatusCode() *int32
	SetBody(v *SubmitComplaintResponseBody) *SubmitComplaintResponse
	GetBody() *SubmitComplaintResponseBody
}

type SubmitComplaintResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitComplaintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitComplaintResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitComplaintResponse) GoString() string {
	return s.String()
}

func (s *SubmitComplaintResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitComplaintResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitComplaintResponse) GetBody() *SubmitComplaintResponseBody {
	return s.Body
}

func (s *SubmitComplaintResponse) SetHeaders(v map[string]*string) *SubmitComplaintResponse {
	s.Headers = v
	return s
}

func (s *SubmitComplaintResponse) SetStatusCode(v int32) *SubmitComplaintResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitComplaintResponse) SetBody(v *SubmitComplaintResponseBody) *SubmitComplaintResponse {
	s.Body = v
	return s
}

func (s *SubmitComplaintResponse) Validate() error {
	return dara.Validate(s)
}
