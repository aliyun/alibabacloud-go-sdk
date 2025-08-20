// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViewDDLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetViewDDLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetViewDDLResponse
	GetStatusCode() *int32
	SetBody(v *GetViewDDLResponseBody) *GetViewDDLResponse
	GetBody() *GetViewDDLResponseBody
}

type GetViewDDLResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetViewDDLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetViewDDLResponse) String() string {
	return dara.Prettify(s)
}

func (s GetViewDDLResponse) GoString() string {
	return s.String()
}

func (s *GetViewDDLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetViewDDLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetViewDDLResponse) GetBody() *GetViewDDLResponseBody {
	return s.Body
}

func (s *GetViewDDLResponse) SetHeaders(v map[string]*string) *GetViewDDLResponse {
	s.Headers = v
	return s
}

func (s *GetViewDDLResponse) SetStatusCode(v int32) *GetViewDDLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetViewDDLResponse) SetBody(v *GetViewDDLResponseBody) *GetViewDDLResponse {
	s.Body = v
	return s
}

func (s *GetViewDDLResponse) Validate() error {
	return dara.Validate(s)
}
