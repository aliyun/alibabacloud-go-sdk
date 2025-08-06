// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarksConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWatermarksConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWatermarksConsoleResponse
	GetStatusCode() *int32
	SetBody(v *GetWatermarksConsoleResponseBody) *GetWatermarksConsoleResponse
	GetBody() *GetWatermarksConsoleResponseBody
}

type GetWatermarksConsoleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWatermarksConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWatermarksConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarksConsoleResponse) GoString() string {
	return s.String()
}

func (s *GetWatermarksConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWatermarksConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWatermarksConsoleResponse) GetBody() *GetWatermarksConsoleResponseBody {
	return s.Body
}

func (s *GetWatermarksConsoleResponse) SetHeaders(v map[string]*string) *GetWatermarksConsoleResponse {
	s.Headers = v
	return s
}

func (s *GetWatermarksConsoleResponse) SetStatusCode(v int32) *GetWatermarksConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWatermarksConsoleResponse) SetBody(v *GetWatermarksConsoleResponseBody) *GetWatermarksConsoleResponse {
	s.Body = v
	return s
}

func (s *GetWatermarksConsoleResponse) Validate() error {
	return dara.Validate(s)
}
