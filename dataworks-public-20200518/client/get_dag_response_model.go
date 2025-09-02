// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDagResponse
	GetStatusCode() *int32
	SetBody(v *GetDagResponseBody) *GetDagResponse
	GetBody() *GetDagResponseBody
}

type GetDagResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDagResponse) GoString() string {
	return s.String()
}

func (s *GetDagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDagResponse) GetBody() *GetDagResponseBody {
	return s.Body
}

func (s *GetDagResponse) SetHeaders(v map[string]*string) *GetDagResponse {
	s.Headers = v
	return s
}

func (s *GetDagResponse) SetStatusCode(v int32) *GetDagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDagResponse) SetBody(v *GetDagResponseBody) *GetDagResponse {
	s.Body = v
	return s
}

func (s *GetDagResponse) Validate() error {
	return dara.Validate(s)
}
