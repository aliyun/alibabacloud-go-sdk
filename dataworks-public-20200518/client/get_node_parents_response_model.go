// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeParentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeParentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeParentsResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeParentsResponseBody) *GetNodeParentsResponse
	GetBody() *GetNodeParentsResponseBody
}

type GetNodeParentsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeParentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeParentsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeParentsResponse) GoString() string {
	return s.String()
}

func (s *GetNodeParentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeParentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeParentsResponse) GetBody() *GetNodeParentsResponseBody {
	return s.Body
}

func (s *GetNodeParentsResponse) SetHeaders(v map[string]*string) *GetNodeParentsResponse {
	s.Headers = v
	return s
}

func (s *GetNodeParentsResponse) SetStatusCode(v int32) *GetNodeParentsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeParentsResponse) SetBody(v *GetNodeParentsResponseBody) *GetNodeParentsResponse {
	s.Body = v
	return s
}

func (s *GetNodeParentsResponse) Validate() error {
	return dara.Validate(s)
}
