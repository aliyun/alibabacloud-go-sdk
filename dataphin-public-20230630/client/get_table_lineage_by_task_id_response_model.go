// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableLineageByTaskIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableLineageByTaskIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableLineageByTaskIdResponse
	GetStatusCode() *int32
	SetBody(v *GetTableLineageByTaskIdResponseBody) *GetTableLineageByTaskIdResponse
	GetBody() *GetTableLineageByTaskIdResponseBody
}

type GetTableLineageByTaskIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableLineageByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableLineageByTaskIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineageByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetTableLineageByTaskIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableLineageByTaskIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableLineageByTaskIdResponse) GetBody() *GetTableLineageByTaskIdResponseBody {
	return s.Body
}

func (s *GetTableLineageByTaskIdResponse) SetHeaders(v map[string]*string) *GetTableLineageByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetTableLineageByTaskIdResponse) SetStatusCode(v int32) *GetTableLineageByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableLineageByTaskIdResponse) SetBody(v *GetTableLineageByTaskIdResponseBody) *GetTableLineageByTaskIdResponse {
	s.Body = v
	return s
}

func (s *GetTableLineageByTaskIdResponse) Validate() error {
	return dara.Validate(s)
}
