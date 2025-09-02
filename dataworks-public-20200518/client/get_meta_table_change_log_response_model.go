// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableChangeLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTableChangeLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTableChangeLogResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTableChangeLogResponseBody) *GetMetaTableChangeLogResponse
	GetBody() *GetMetaTableChangeLogResponseBody
}

type GetMetaTableChangeLogResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTableChangeLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTableChangeLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableChangeLogResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableChangeLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTableChangeLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTableChangeLogResponse) GetBody() *GetMetaTableChangeLogResponseBody {
	return s.Body
}

func (s *GetMetaTableChangeLogResponse) SetHeaders(v map[string]*string) *GetMetaTableChangeLogResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableChangeLogResponse) SetStatusCode(v int32) *GetMetaTableChangeLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTableChangeLogResponse) SetBody(v *GetMetaTableChangeLogResponseBody) *GetMetaTableChangeLogResponse {
	s.Body = v
	return s
}

func (s *GetMetaTableChangeLogResponse) Validate() error {
	return dara.Validate(s)
}
