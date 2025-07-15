// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocClusterTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDocClusterTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDocClusterTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDocClusterTaskResponseBody) *SubmitDocClusterTaskResponse
	GetBody() *SubmitDocClusterTaskResponseBody
}

type SubmitDocClusterTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocClusterTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocClusterTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocClusterTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDocClusterTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDocClusterTaskResponse) GetBody() *SubmitDocClusterTaskResponseBody {
	return s.Body
}

func (s *SubmitDocClusterTaskResponse) SetHeaders(v map[string]*string) *SubmitDocClusterTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocClusterTaskResponse) SetStatusCode(v int32) *SubmitDocClusterTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocClusterTaskResponse) SetBody(v *SubmitDocClusterTaskResponseBody) *SubmitDocClusterTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitDocClusterTaskResponse) Validate() error {
	return dara.Validate(s)
}
