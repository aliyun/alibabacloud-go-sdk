// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStopApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStopApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *BatchStopApplicationsResponseBody) *BatchStopApplicationsResponse
	GetBody() *BatchStopApplicationsResponseBody
}

type BatchStopApplicationsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStopApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStopApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStopApplicationsResponse) GoString() string {
	return s.String()
}

func (s *BatchStopApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStopApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStopApplicationsResponse) GetBody() *BatchStopApplicationsResponseBody {
	return s.Body
}

func (s *BatchStopApplicationsResponse) SetHeaders(v map[string]*string) *BatchStopApplicationsResponse {
	s.Headers = v
	return s
}

func (s *BatchStopApplicationsResponse) SetStatusCode(v int32) *BatchStopApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStopApplicationsResponse) SetBody(v *BatchStopApplicationsResponseBody) *BatchStopApplicationsResponse {
	s.Body = v
	return s
}

func (s *BatchStopApplicationsResponse) Validate() error {
	return dara.Validate(s)
}
