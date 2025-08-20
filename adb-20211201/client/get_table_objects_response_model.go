// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableObjectsResponse
	GetStatusCode() *int32
	SetBody(v *GetTableObjectsResponseBody) *GetTableObjectsResponse
	GetBody() *GetTableObjectsResponseBody
}

type GetTableObjectsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableObjectsResponse) GoString() string {
	return s.String()
}

func (s *GetTableObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableObjectsResponse) GetBody() *GetTableObjectsResponseBody {
	return s.Body
}

func (s *GetTableObjectsResponse) SetHeaders(v map[string]*string) *GetTableObjectsResponse {
	s.Headers = v
	return s
}

func (s *GetTableObjectsResponse) SetStatusCode(v int32) *GetTableObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableObjectsResponse) SetBody(v *GetTableObjectsResponseBody) *GetTableObjectsResponse {
	s.Body = v
	return s
}

func (s *GetTableObjectsResponse) Validate() error {
	return dara.Validate(s)
}
