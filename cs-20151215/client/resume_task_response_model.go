// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeTaskResponse
	GetStatusCode() *int32
}

type ResumeTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s ResumeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeTaskResponse) GoString() string {
	return s.String()
}

func (s *ResumeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeTaskResponse) SetHeaders(v map[string]*string) *ResumeTaskResponse {
	s.Headers = v
	return s
}

func (s *ResumeTaskResponse) SetStatusCode(v int32) *ResumeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeTaskResponse) Validate() error {
	return dara.Validate(s)
}
