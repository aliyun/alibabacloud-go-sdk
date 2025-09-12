// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDataSetTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDataSetTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDataSetTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDataSetTaskResponseBody) *SubmitDataSetTaskResponse
	GetBody() *SubmitDataSetTaskResponseBody
}

type SubmitDataSetTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDataSetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDataSetTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDataSetTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitDataSetTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDataSetTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDataSetTaskResponse) GetBody() *SubmitDataSetTaskResponseBody {
	return s.Body
}

func (s *SubmitDataSetTaskResponse) SetHeaders(v map[string]*string) *SubmitDataSetTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitDataSetTaskResponse) SetStatusCode(v int32) *SubmitDataSetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDataSetTaskResponse) SetBody(v *SubmitDataSetTaskResponseBody) *SubmitDataSetTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitDataSetTaskResponse) Validate() error {
	return dara.Validate(s)
}
