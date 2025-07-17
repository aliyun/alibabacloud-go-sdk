// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineDDLProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOnlineDDLProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOnlineDDLProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetOnlineDDLProgressResponseBody) *GetOnlineDDLProgressResponse
	GetBody() *GetOnlineDDLProgressResponseBody
}

type GetOnlineDDLProgressResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOnlineDDLProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOnlineDDLProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineDDLProgressResponse) GoString() string {
	return s.String()
}

func (s *GetOnlineDDLProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOnlineDDLProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOnlineDDLProgressResponse) GetBody() *GetOnlineDDLProgressResponseBody {
	return s.Body
}

func (s *GetOnlineDDLProgressResponse) SetHeaders(v map[string]*string) *GetOnlineDDLProgressResponse {
	s.Headers = v
	return s
}

func (s *GetOnlineDDLProgressResponse) SetStatusCode(v int32) *GetOnlineDDLProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOnlineDDLProgressResponse) SetBody(v *GetOnlineDDLProgressResponseBody) *GetOnlineDDLProgressResponse {
	s.Body = v
	return s
}

func (s *GetOnlineDDLProgressResponse) Validate() error {
	return dara.Validate(s)
}
