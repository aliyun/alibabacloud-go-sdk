// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarFsObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarFsObjectsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarFsObjectsResponseBody) *DescribePolarFsObjectsResponse
	GetBody() *DescribePolarFsObjectsResponseBody
}

type DescribePolarFsObjectsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarFsObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarFsObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsObjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarFsObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarFsObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarFsObjectsResponse) GetBody() *DescribePolarFsObjectsResponseBody {
	return s.Body
}

func (s *DescribePolarFsObjectsResponse) SetHeaders(v map[string]*string) *DescribePolarFsObjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarFsObjectsResponse) SetStatusCode(v int32) *DescribePolarFsObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarFsObjectsResponse) SetBody(v *DescribePolarFsObjectsResponseBody) *DescribePolarFsObjectsResponse {
	s.Body = v
	return s
}

func (s *DescribePolarFsObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
