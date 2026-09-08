// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadCategoryGroupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadCategoryGroupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadCategoryGroupListResponse
	GetStatusCode() *int32
	SetBody(v *ReadCategoryGroupListResponseBody) *ReadCategoryGroupListResponse
	GetBody() *ReadCategoryGroupListResponseBody
}

type ReadCategoryGroupListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadCategoryGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadCategoryGroupListResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadCategoryGroupListResponse) GoString() string {
	return s.String()
}

func (s *ReadCategoryGroupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadCategoryGroupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadCategoryGroupListResponse) GetBody() *ReadCategoryGroupListResponseBody {
	return s.Body
}

func (s *ReadCategoryGroupListResponse) SetHeaders(v map[string]*string) *ReadCategoryGroupListResponse {
	s.Headers = v
	return s
}

func (s *ReadCategoryGroupListResponse) SetStatusCode(v int32) *ReadCategoryGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadCategoryGroupListResponse) SetBody(v *ReadCategoryGroupListResponseBody) *ReadCategoryGroupListResponse {
	s.Body = v
	return s
}

func (s *ReadCategoryGroupListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
