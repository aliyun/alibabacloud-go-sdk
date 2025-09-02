// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaColumnLineageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaColumnLineageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaColumnLineageResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaColumnLineageResponseBody) *GetMetaColumnLineageResponse
	GetBody() *GetMetaColumnLineageResponseBody
}

type GetMetaColumnLineageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaColumnLineageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaColumnLineageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaColumnLineageResponse) GoString() string {
	return s.String()
}

func (s *GetMetaColumnLineageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaColumnLineageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaColumnLineageResponse) GetBody() *GetMetaColumnLineageResponseBody {
	return s.Body
}

func (s *GetMetaColumnLineageResponse) SetHeaders(v map[string]*string) *GetMetaColumnLineageResponse {
	s.Headers = v
	return s
}

func (s *GetMetaColumnLineageResponse) SetStatusCode(v int32) *GetMetaColumnLineageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaColumnLineageResponse) SetBody(v *GetMetaColumnLineageResponseBody) *GetMetaColumnLineageResponse {
	s.Body = v
	return s
}

func (s *GetMetaColumnLineageResponse) Validate() error {
	return dara.Validate(s)
}
