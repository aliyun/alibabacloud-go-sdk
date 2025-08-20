// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStackResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStackResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetStackResourceResponseBody) *GetStackResourceResponse
	GetBody() *GetStackResourceResponseBody
}

type GetStackResourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStackResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStackResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStackResourceResponse) GoString() string {
	return s.String()
}

func (s *GetStackResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStackResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStackResourceResponse) GetBody() *GetStackResourceResponseBody {
	return s.Body
}

func (s *GetStackResourceResponse) SetHeaders(v map[string]*string) *GetStackResourceResponse {
	s.Headers = v
	return s
}

func (s *GetStackResourceResponse) SetStatusCode(v int32) *GetStackResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackResourceResponse) SetBody(v *GetStackResourceResponseBody) *GetStackResourceResponse {
	s.Body = v
	return s
}

func (s *GetStackResourceResponse) Validate() error {
	return dara.Validate(s)
}
