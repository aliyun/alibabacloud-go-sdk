// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDNAInitializationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDNAInitializationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDNAInitializationJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDNAInitializationJobResponseBody) *SubmitDNAInitializationJobResponse
	GetBody() *SubmitDNAInitializationJobResponseBody
}

type SubmitDNAInitializationJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDNAInitializationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDNAInitializationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDNAInitializationJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDNAInitializationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDNAInitializationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDNAInitializationJobResponse) GetBody() *SubmitDNAInitializationJobResponseBody {
	return s.Body
}

func (s *SubmitDNAInitializationJobResponse) SetHeaders(v map[string]*string) *SubmitDNAInitializationJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDNAInitializationJobResponse) SetStatusCode(v int32) *SubmitDNAInitializationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDNAInitializationJobResponse) SetBody(v *SubmitDNAInitializationJobResponseBody) *SubmitDNAInitializationJobResponse {
	s.Body = v
	return s
}

func (s *SubmitDNAInitializationJobResponse) Validate() error {
	return dara.Validate(s)
}
