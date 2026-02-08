// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMaterialTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMaterialTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryMaterialTaskDetailResponseBody) *QueryMaterialTaskDetailResponse
	GetBody() *QueryMaterialTaskDetailResponseBody
}

type QueryMaterialTaskDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMaterialTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMaterialTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryMaterialTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMaterialTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMaterialTaskDetailResponse) GetBody() *QueryMaterialTaskDetailResponseBody {
	return s.Body
}

func (s *QueryMaterialTaskDetailResponse) SetHeaders(v map[string]*string) *QueryMaterialTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryMaterialTaskDetailResponse) SetStatusCode(v int32) *QueryMaterialTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMaterialTaskDetailResponse) SetBody(v *QueryMaterialTaskDetailResponseBody) *QueryMaterialTaskDetailResponse {
	s.Body = v
	return s
}

func (s *QueryMaterialTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
