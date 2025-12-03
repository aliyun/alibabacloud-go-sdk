// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeColdStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeColdStorageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeColdStorageResponseBody) *DescribeColdStorageResponse
	GetBody() *DescribeColdStorageResponseBody
}

type DescribeColdStorageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColdStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColdStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeColdStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeColdStorageResponse) GetBody() *DescribeColdStorageResponseBody {
	return s.Body
}

func (s *DescribeColdStorageResponse) SetHeaders(v map[string]*string) *DescribeColdStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeColdStorageResponse) SetStatusCode(v int32) *DescribeColdStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColdStorageResponse) SetBody(v *DescribeColdStorageResponseBody) *DescribeColdStorageResponse {
	s.Body = v
	return s
}

func (s *DescribeColdStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
