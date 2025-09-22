// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFootprintListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFootprintListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFootprintListResponse
	GetStatusCode() *int32
	SetBody(v *GetFootprintListResponseBody) *GetFootprintListResponse
	GetBody() *GetFootprintListResponseBody
}

type GetFootprintListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFootprintListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFootprintListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFootprintListResponse) GoString() string {
	return s.String()
}

func (s *GetFootprintListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFootprintListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFootprintListResponse) GetBody() *GetFootprintListResponseBody {
	return s.Body
}

func (s *GetFootprintListResponse) SetHeaders(v map[string]*string) *GetFootprintListResponse {
	s.Headers = v
	return s
}

func (s *GetFootprintListResponse) SetStatusCode(v int32) *GetFootprintListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFootprintListResponse) SetBody(v *GetFootprintListResponseBody) *GetFootprintListResponse {
	s.Body = v
	return s
}

func (s *GetFootprintListResponse) Validate() error {
	return dara.Validate(s)
}
