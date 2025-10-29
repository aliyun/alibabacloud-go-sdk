// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarFsAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarFsAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarFsAttributeResponseBody) *DescribePolarFsAttributeResponse
	GetBody() *DescribePolarFsAttributeResponseBody
}

type DescribePolarFsAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarFsAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarFsAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarFsAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarFsAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarFsAttributeResponse) GetBody() *DescribePolarFsAttributeResponseBody {
	return s.Body
}

func (s *DescribePolarFsAttributeResponse) SetHeaders(v map[string]*string) *DescribePolarFsAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarFsAttributeResponse) SetStatusCode(v int32) *DescribePolarFsAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarFsAttributeResponse) SetBody(v *DescribePolarFsAttributeResponseBody) *DescribePolarFsAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribePolarFsAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
