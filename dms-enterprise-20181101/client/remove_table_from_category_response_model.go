// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTableFromCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveTableFromCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveTableFromCategoryResponse
	GetStatusCode() *int32
	SetBody(v *RemoveTableFromCategoryResponseBody) *RemoveTableFromCategoryResponse
	GetBody() *RemoveTableFromCategoryResponseBody
}

type RemoveTableFromCategoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTableFromCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveTableFromCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveTableFromCategoryResponse) GoString() string {
	return s.String()
}

func (s *RemoveTableFromCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveTableFromCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveTableFromCategoryResponse) GetBody() *RemoveTableFromCategoryResponseBody {
	return s.Body
}

func (s *RemoveTableFromCategoryResponse) SetHeaders(v map[string]*string) *RemoveTableFromCategoryResponse {
	s.Headers = v
	return s
}

func (s *RemoveTableFromCategoryResponse) SetStatusCode(v int32) *RemoveTableFromCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTableFromCategoryResponse) SetBody(v *RemoveTableFromCategoryResponseBody) *RemoveTableFromCategoryResponse {
	s.Body = v
	return s
}

func (s *RemoveTableFromCategoryResponse) Validate() error {
	return dara.Validate(s)
}
