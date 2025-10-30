// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountFactoryBaselineItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccountFactoryBaselineItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccountFactoryBaselineItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListAccountFactoryBaselineItemsResponseBody) *ListAccountFactoryBaselineItemsResponse
	GetBody() *ListAccountFactoryBaselineItemsResponseBody
}

type ListAccountFactoryBaselineItemsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountFactoryBaselineItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountFactoryBaselineItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccountFactoryBaselineItemsResponse) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselineItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccountFactoryBaselineItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccountFactoryBaselineItemsResponse) GetBody() *ListAccountFactoryBaselineItemsResponseBody {
	return s.Body
}

func (s *ListAccountFactoryBaselineItemsResponse) SetHeaders(v map[string]*string) *ListAccountFactoryBaselineItemsResponse {
	s.Headers = v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponse) SetStatusCode(v int32) *ListAccountFactoryBaselineItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponse) SetBody(v *ListAccountFactoryBaselineItemsResponseBody) *ListAccountFactoryBaselineItemsResponse {
	s.Body = v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
