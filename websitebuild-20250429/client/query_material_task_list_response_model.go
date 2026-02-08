// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMaterialTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMaterialTaskListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMaterialTaskListResponseBody) *QueryMaterialTaskListResponse
	GetBody() *QueryMaterialTaskListResponseBody
}

type QueryMaterialTaskListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMaterialTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMaterialTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialTaskListResponse) GoString() string {
	return s.String()
}

func (s *QueryMaterialTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMaterialTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMaterialTaskListResponse) GetBody() *QueryMaterialTaskListResponseBody {
	return s.Body
}

func (s *QueryMaterialTaskListResponse) SetHeaders(v map[string]*string) *QueryMaterialTaskListResponse {
	s.Headers = v
	return s
}

func (s *QueryMaterialTaskListResponse) SetStatusCode(v int32) *QueryMaterialTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMaterialTaskListResponse) SetBody(v *QueryMaterialTaskListResponseBody) *QueryMaterialTaskListResponse {
	s.Body = v
	return s
}

func (s *QueryMaterialTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
