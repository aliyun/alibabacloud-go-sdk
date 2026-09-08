// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataMaskingInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataMaskingInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataMaskingInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataMaskingInstancesResponseBody) *ListDataMaskingInstancesResponse
	GetBody() *ListDataMaskingInstancesResponseBody
}

type ListDataMaskingInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataMaskingInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataMaskingInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListDataMaskingInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataMaskingInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataMaskingInstancesResponse) GetBody() *ListDataMaskingInstancesResponseBody {
	return s.Body
}

func (s *ListDataMaskingInstancesResponse) SetHeaders(v map[string]*string) *ListDataMaskingInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListDataMaskingInstancesResponse) SetStatusCode(v int32) *ListDataMaskingInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataMaskingInstancesResponse) SetBody(v *ListDataMaskingInstancesResponseBody) *ListDataMaskingInstancesResponse {
	s.Body = v
	return s
}

func (s *ListDataMaskingInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
