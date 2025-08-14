// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCopyrightExtractJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitCopyrightExtractJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitCopyrightExtractJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitCopyrightExtractJobResponseBody) *SubmitCopyrightExtractJobResponse
	GetBody() *SubmitCopyrightExtractJobResponseBody
}

type SubmitCopyrightExtractJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCopyrightExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCopyrightExtractJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightExtractJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitCopyrightExtractJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitCopyrightExtractJobResponse) GetBody() *SubmitCopyrightExtractJobResponseBody {
	return s.Body
}

func (s *SubmitCopyrightExtractJobResponse) SetHeaders(v map[string]*string) *SubmitCopyrightExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitCopyrightExtractJobResponse) SetStatusCode(v int32) *SubmitCopyrightExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCopyrightExtractJobResponse) SetBody(v *SubmitCopyrightExtractJobResponseBody) *SubmitCopyrightExtractJobResponse {
	s.Body = v
	return s
}

func (s *SubmitCopyrightExtractJobResponse) Validate() error {
	return dara.Validate(s)
}
