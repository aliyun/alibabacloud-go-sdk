// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsDataSourceConfigItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMmsDataSourceConfigItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMmsDataSourceConfigItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListMmsDataSourceConfigItemsResponseBody) *ListMmsDataSourceConfigItemsResponse
	GetBody() *ListMmsDataSourceConfigItemsResponseBody
}

type ListMmsDataSourceConfigItemsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsDataSourceConfigItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsDataSourceConfigItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDataSourceConfigItemsResponse) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourceConfigItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMmsDataSourceConfigItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMmsDataSourceConfigItemsResponse) GetBody() *ListMmsDataSourceConfigItemsResponseBody {
	return s.Body
}

func (s *ListMmsDataSourceConfigItemsResponse) SetHeaders(v map[string]*string) *ListMmsDataSourceConfigItemsResponse {
	s.Headers = v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponse) SetStatusCode(v int32) *ListMmsDataSourceConfigItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponse) SetBody(v *ListMmsDataSourceConfigItemsResponseBody) *ListMmsDataSourceConfigItemsResponse {
	s.Body = v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
