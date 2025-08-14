// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDNAJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDNAJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDNAJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDNAJobResponseBody) *SubmitDNAJobResponse
	GetBody() *SubmitDNAJobResponseBody
}

type SubmitDNAJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDNAJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDNAJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDNAJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDNAJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDNAJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDNAJobResponse) GetBody() *SubmitDNAJobResponseBody {
	return s.Body
}

func (s *SubmitDNAJobResponse) SetHeaders(v map[string]*string) *SubmitDNAJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDNAJobResponse) SetStatusCode(v int32) *SubmitDNAJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDNAJobResponse) SetBody(v *SubmitDNAJobResponseBody) *SubmitDNAJobResponse {
	s.Body = v
	return s
}

func (s *SubmitDNAJobResponse) Validate() error {
	return dara.Validate(s)
}
