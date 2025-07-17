// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLhSpaceByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLhSpaceByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLhSpaceByNameResponse
	GetStatusCode() *int32
	SetBody(v *GetLhSpaceByNameResponseBody) *GetLhSpaceByNameResponse
	GetBody() *GetLhSpaceByNameResponseBody
}

type GetLhSpaceByNameResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLhSpaceByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLhSpaceByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLhSpaceByNameResponse) GoString() string {
	return s.String()
}

func (s *GetLhSpaceByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLhSpaceByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLhSpaceByNameResponse) GetBody() *GetLhSpaceByNameResponseBody {
	return s.Body
}

func (s *GetLhSpaceByNameResponse) SetHeaders(v map[string]*string) *GetLhSpaceByNameResponse {
	s.Headers = v
	return s
}

func (s *GetLhSpaceByNameResponse) SetStatusCode(v int32) *GetLhSpaceByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLhSpaceByNameResponse) SetBody(v *GetLhSpaceByNameResponseBody) *GetLhSpaceByNameResponse {
	s.Body = v
	return s
}

func (s *GetLhSpaceByNameResponse) Validate() error {
	return dara.Validate(s)
}
