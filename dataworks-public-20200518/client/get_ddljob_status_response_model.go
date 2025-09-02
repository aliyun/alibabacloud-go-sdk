// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDDLJobStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDDLJobStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDDLJobStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetDDLJobStatusResponseBody) *GetDDLJobStatusResponse
	GetBody() *GetDDLJobStatusResponseBody
}

type GetDDLJobStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDDLJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDDLJobStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDDLJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDDLJobStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDDLJobStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDDLJobStatusResponse) GetBody() *GetDDLJobStatusResponseBody {
	return s.Body
}

func (s *GetDDLJobStatusResponse) SetHeaders(v map[string]*string) *GetDDLJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDDLJobStatusResponse) SetStatusCode(v int32) *GetDDLJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDDLJobStatusResponse) SetBody(v *GetDDLJobStatusResponseBody) *GetDDLJobStatusResponse {
	s.Body = v
	return s
}

func (s *GetDDLJobStatusResponse) Validate() error {
	return dara.Validate(s)
}
