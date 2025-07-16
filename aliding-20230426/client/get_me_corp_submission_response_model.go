// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeCorpSubmissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMeCorpSubmissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMeCorpSubmissionResponse
	GetStatusCode() *int32
	SetBody(v *GetMeCorpSubmissionResponseBody) *GetMeCorpSubmissionResponse
	GetBody() *GetMeCorpSubmissionResponseBody
}

type GetMeCorpSubmissionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMeCorpSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMeCorpSubmissionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMeCorpSubmissionResponse) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMeCorpSubmissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMeCorpSubmissionResponse) GetBody() *GetMeCorpSubmissionResponseBody {
	return s.Body
}

func (s *GetMeCorpSubmissionResponse) SetHeaders(v map[string]*string) *GetMeCorpSubmissionResponse {
	s.Headers = v
	return s
}

func (s *GetMeCorpSubmissionResponse) SetStatusCode(v int32) *GetMeCorpSubmissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMeCorpSubmissionResponse) SetBody(v *GetMeCorpSubmissionResponseBody) *GetMeCorpSubmissionResponse {
	s.Body = v
	return s
}

func (s *GetMeCorpSubmissionResponse) Validate() error {
	return dara.Validate(s)
}
