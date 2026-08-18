// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateDomainItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateDomainItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateDomainItemsResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateDomainItemsResponseBody) *BatchCreateDomainItemsResponse
	GetBody() *BatchCreateDomainItemsResponseBody
}

type BatchCreateDomainItemsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateDomainItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateDomainItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateDomainItemsResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateDomainItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateDomainItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateDomainItemsResponse) GetBody() *BatchCreateDomainItemsResponseBody {
	return s.Body
}

func (s *BatchCreateDomainItemsResponse) SetHeaders(v map[string]*string) *BatchCreateDomainItemsResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateDomainItemsResponse) SetStatusCode(v int32) *BatchCreateDomainItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateDomainItemsResponse) SetBody(v *BatchCreateDomainItemsResponseBody) *BatchCreateDomainItemsResponse {
	s.Body = v
	return s
}

func (s *BatchCreateDomainItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
