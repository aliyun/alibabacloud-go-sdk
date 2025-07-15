// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchExportWordTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchExportWordTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchExportWordTaskResponse
	GetStatusCode() *int32
	SetBody(v *FetchExportWordTaskResponseBody) *FetchExportWordTaskResponse
	GetBody() *FetchExportWordTaskResponseBody
}

type FetchExportWordTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchExportWordTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchExportWordTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchExportWordTaskResponse) GoString() string {
	return s.String()
}

func (s *FetchExportWordTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchExportWordTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchExportWordTaskResponse) GetBody() *FetchExportWordTaskResponseBody {
	return s.Body
}

func (s *FetchExportWordTaskResponse) SetHeaders(v map[string]*string) *FetchExportWordTaskResponse {
	s.Headers = v
	return s
}

func (s *FetchExportWordTaskResponse) SetStatusCode(v int32) *FetchExportWordTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchExportWordTaskResponse) SetBody(v *FetchExportWordTaskResponseBody) *FetchExportWordTaskResponse {
	s.Body = v
	return s
}

func (s *FetchExportWordTaskResponse) Validate() error {
	return dara.Validate(s)
}
