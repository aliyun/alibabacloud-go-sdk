// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocParsingTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDocParsingTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDocParsingTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDocParsingTaskResponseBody) *SubmitDocParsingTaskResponse
	GetBody() *SubmitDocParsingTaskResponseBody
}

type SubmitDocParsingTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocParsingTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocParsingTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParsingTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocParsingTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDocParsingTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDocParsingTaskResponse) GetBody() *SubmitDocParsingTaskResponseBody {
	return s.Body
}

func (s *SubmitDocParsingTaskResponse) SetHeaders(v map[string]*string) *SubmitDocParsingTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocParsingTaskResponse) SetStatusCode(v int32) *SubmitDocParsingTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocParsingTaskResponse) SetBody(v *SubmitDocParsingTaskResponseBody) *SubmitDocParsingTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitDocParsingTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
