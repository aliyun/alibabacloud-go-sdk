// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDomainItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteDomainItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteDomainItemsResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteDomainItemsResponseBody) *BatchDeleteDomainItemsResponse
	GetBody() *BatchDeleteDomainItemsResponseBody
}

type BatchDeleteDomainItemsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteDomainItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteDomainItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDomainItemsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteDomainItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteDomainItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteDomainItemsResponse) GetBody() *BatchDeleteDomainItemsResponseBody {
	return s.Body
}

func (s *BatchDeleteDomainItemsResponse) SetHeaders(v map[string]*string) *BatchDeleteDomainItemsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteDomainItemsResponse) SetStatusCode(v int32) *BatchDeleteDomainItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteDomainItemsResponse) SetBody(v *BatchDeleteDomainItemsResponseBody) *BatchDeleteDomainItemsResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteDomainItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
