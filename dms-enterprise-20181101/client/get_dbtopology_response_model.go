// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBTopologyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDBTopologyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDBTopologyResponse
	GetStatusCode() *int32
	SetBody(v *GetDBTopologyResponseBody) *GetDBTopologyResponse
	GetBody() *GetDBTopologyResponseBody
}

type GetDBTopologyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDBTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDBTopologyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDBTopologyResponse) GoString() string {
	return s.String()
}

func (s *GetDBTopologyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDBTopologyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDBTopologyResponse) GetBody() *GetDBTopologyResponseBody {
	return s.Body
}

func (s *GetDBTopologyResponse) SetHeaders(v map[string]*string) *GetDBTopologyResponse {
	s.Headers = v
	return s
}

func (s *GetDBTopologyResponse) SetStatusCode(v int32) *GetDBTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDBTopologyResponse) SetBody(v *GetDBTopologyResponseBody) *GetDBTopologyResponse {
	s.Body = v
	return s
}

func (s *GetDBTopologyResponse) Validate() error {
	return dara.Validate(s)
}
