// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaDataComponentPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetaDataComponentPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetaDataComponentPageResponse
	GetStatusCode() *int32
	SetBody(v *ListMetaDataComponentPageResponseBody) *ListMetaDataComponentPageResponse
	GetBody() *ListMetaDataComponentPageResponseBody
}

type ListMetaDataComponentPageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetaDataComponentPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetaDataComponentPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetaDataComponentPageResponse) GoString() string {
	return s.String()
}

func (s *ListMetaDataComponentPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetaDataComponentPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetaDataComponentPageResponse) GetBody() *ListMetaDataComponentPageResponseBody {
	return s.Body
}

func (s *ListMetaDataComponentPageResponse) SetHeaders(v map[string]*string) *ListMetaDataComponentPageResponse {
	s.Headers = v
	return s
}

func (s *ListMetaDataComponentPageResponse) SetStatusCode(v int32) *ListMetaDataComponentPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetaDataComponentPageResponse) SetBody(v *ListMetaDataComponentPageResponseBody) *ListMetaDataComponentPageResponse {
	s.Body = v
	return s
}

func (s *ListMetaDataComponentPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
