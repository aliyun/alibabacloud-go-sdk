// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobNameListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobNameListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobNameListResponse
	GetStatusCode() *int32
	SetBody(v *GetJobNameListResponseBody) *GetJobNameListResponse
	GetBody() *GetJobNameListResponseBody
}

type GetJobNameListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobNameListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobNameListResponse) GoString() string {
	return s.String()
}

func (s *GetJobNameListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobNameListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobNameListResponse) GetBody() *GetJobNameListResponseBody {
	return s.Body
}

func (s *GetJobNameListResponse) SetHeaders(v map[string]*string) *GetJobNameListResponse {
	s.Headers = v
	return s
}

func (s *GetJobNameListResponse) SetStatusCode(v int32) *GetJobNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobNameListResponse) SetBody(v *GetJobNameListResponseBody) *GetJobNameListResponse {
	s.Body = v
	return s
}

func (s *GetJobNameListResponse) Validate() error {
	return dara.Validate(s)
}
