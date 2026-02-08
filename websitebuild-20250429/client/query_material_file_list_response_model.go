// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialFileListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMaterialFileListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMaterialFileListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMaterialFileListResponseBody) *QueryMaterialFileListResponse
	GetBody() *QueryMaterialFileListResponseBody
}

type QueryMaterialFileListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMaterialFileListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMaterialFileListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileListResponse) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMaterialFileListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMaterialFileListResponse) GetBody() *QueryMaterialFileListResponseBody {
	return s.Body
}

func (s *QueryMaterialFileListResponse) SetHeaders(v map[string]*string) *QueryMaterialFileListResponse {
	s.Headers = v
	return s
}

func (s *QueryMaterialFileListResponse) SetStatusCode(v int32) *QueryMaterialFileListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMaterialFileListResponse) SetBody(v *QueryMaterialFileListResponseBody) *QueryMaterialFileListResponse {
	s.Body = v
	return s
}

func (s *QueryMaterialFileListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
