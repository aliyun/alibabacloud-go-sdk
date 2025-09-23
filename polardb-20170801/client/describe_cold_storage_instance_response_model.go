// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdStorageInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeColdStorageInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeColdStorageInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeColdStorageInstanceResponseBody) *DescribeColdStorageInstanceResponse
	GetBody() *DescribeColdStorageInstanceResponseBody
}

type DescribeColdStorageInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColdStorageInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColdStorageInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdStorageInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeColdStorageInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeColdStorageInstanceResponse) GetBody() *DescribeColdStorageInstanceResponseBody {
	return s.Body
}

func (s *DescribeColdStorageInstanceResponse) SetHeaders(v map[string]*string) *DescribeColdStorageInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeColdStorageInstanceResponse) SetStatusCode(v int32) *DescribeColdStorageInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColdStorageInstanceResponse) SetBody(v *DescribeColdStorageInstanceResponseBody) *DescribeColdStorageInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeColdStorageInstanceResponse) Validate() error {
	return dara.Validate(s)
}
