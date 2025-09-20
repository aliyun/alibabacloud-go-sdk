// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageSplicingTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageSplicingTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageSplicingTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageSplicingTaskResponseBody) *CreateImageSplicingTaskResponse
	GetBody() *CreateImageSplicingTaskResponseBody
}

type CreateImageSplicingTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageSplicingTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageSplicingTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageSplicingTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageSplicingTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageSplicingTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageSplicingTaskResponse) GetBody() *CreateImageSplicingTaskResponseBody {
	return s.Body
}

func (s *CreateImageSplicingTaskResponse) SetHeaders(v map[string]*string) *CreateImageSplicingTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageSplicingTaskResponse) SetStatusCode(v int32) *CreateImageSplicingTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageSplicingTaskResponse) SetBody(v *CreateImageSplicingTaskResponseBody) *CreateImageSplicingTaskResponse {
	s.Body = v
	return s
}

func (s *CreateImageSplicingTaskResponse) Validate() error {
	return dara.Validate(s)
}
