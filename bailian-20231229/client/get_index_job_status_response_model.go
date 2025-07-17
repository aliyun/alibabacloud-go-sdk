// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexJobStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIndexJobStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIndexJobStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetIndexJobStatusResponseBody) *GetIndexJobStatusResponse
	GetBody() *GetIndexJobStatusResponseBody
}

type GetIndexJobStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndexJobStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIndexJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIndexJobStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIndexJobStatusResponse) GetBody() *GetIndexJobStatusResponseBody {
	return s.Body
}

func (s *GetIndexJobStatusResponse) SetHeaders(v map[string]*string) *GetIndexJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetIndexJobStatusResponse) SetStatusCode(v int32) *GetIndexJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexJobStatusResponse) SetBody(v *GetIndexJobStatusResponseBody) *GetIndexJobStatusResponse {
	s.Body = v
	return s
}

func (s *GetIndexJobStatusResponse) Validate() error {
	return dara.Validate(s)
}
