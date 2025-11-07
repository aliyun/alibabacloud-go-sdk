// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpsItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOpsItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOpsItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListOpsItemsResponseBody) *ListOpsItemsResponse
	GetBody() *ListOpsItemsResponseBody
}

type ListOpsItemsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOpsItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOpsItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOpsItemsResponse) GoString() string {
	return s.String()
}

func (s *ListOpsItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOpsItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOpsItemsResponse) GetBody() *ListOpsItemsResponseBody {
	return s.Body
}

func (s *ListOpsItemsResponse) SetHeaders(v map[string]*string) *ListOpsItemsResponse {
	s.Headers = v
	return s
}

func (s *ListOpsItemsResponse) SetStatusCode(v int32) *ListOpsItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpsItemsResponse) SetBody(v *ListOpsItemsResponseBody) *ListOpsItemsResponse {
	s.Body = v
	return s
}

func (s *ListOpsItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
