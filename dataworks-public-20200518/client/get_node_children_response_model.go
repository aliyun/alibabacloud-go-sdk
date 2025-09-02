// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeChildrenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeChildrenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeChildrenResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeChildrenResponseBody) *GetNodeChildrenResponse
	GetBody() *GetNodeChildrenResponseBody
}

type GetNodeChildrenResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeChildrenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeChildrenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeChildrenResponse) GoString() string {
	return s.String()
}

func (s *GetNodeChildrenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeChildrenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeChildrenResponse) GetBody() *GetNodeChildrenResponseBody {
	return s.Body
}

func (s *GetNodeChildrenResponse) SetHeaders(v map[string]*string) *GetNodeChildrenResponse {
	s.Headers = v
	return s
}

func (s *GetNodeChildrenResponse) SetStatusCode(v int32) *GetNodeChildrenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeChildrenResponse) SetBody(v *GetNodeChildrenResponseBody) *GetNodeChildrenResponse {
	s.Body = v
	return s
}

func (s *GetNodeChildrenResponse) Validate() error {
	return dara.Validate(s)
}
