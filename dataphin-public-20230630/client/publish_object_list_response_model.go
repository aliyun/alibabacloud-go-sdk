// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishObjectListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishObjectListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishObjectListResponse
	GetStatusCode() *int32
	SetBody(v *PublishObjectListResponseBody) *PublishObjectListResponse
	GetBody() *PublishObjectListResponseBody
}

type PublishObjectListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishObjectListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishObjectListResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishObjectListResponse) GoString() string {
	return s.String()
}

func (s *PublishObjectListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishObjectListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishObjectListResponse) GetBody() *PublishObjectListResponseBody {
	return s.Body
}

func (s *PublishObjectListResponse) SetHeaders(v map[string]*string) *PublishObjectListResponse {
	s.Headers = v
	return s
}

func (s *PublishObjectListResponse) SetStatusCode(v int32) *PublishObjectListResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishObjectListResponse) SetBody(v *PublishObjectListResponseBody) *PublishObjectListResponse {
	s.Body = v
	return s
}

func (s *PublishObjectListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
