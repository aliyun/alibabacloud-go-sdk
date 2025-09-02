// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableLineageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTableLineageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTableLineageResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTableLineageResponseBody) *GetMetaTableLineageResponse
	GetBody() *GetMetaTableLineageResponseBody
}

type GetMetaTableLineageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTableLineageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTableLineageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableLineageResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableLineageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTableLineageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTableLineageResponse) GetBody() *GetMetaTableLineageResponseBody {
	return s.Body
}

func (s *GetMetaTableLineageResponse) SetHeaders(v map[string]*string) *GetMetaTableLineageResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableLineageResponse) SetStatusCode(v int32) *GetMetaTableLineageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTableLineageResponse) SetBody(v *GetMetaTableLineageResponseBody) *GetMetaTableLineageResponse {
	s.Body = v
	return s
}

func (s *GetMetaTableLineageResponse) Validate() error {
	return dara.Validate(s)
}
