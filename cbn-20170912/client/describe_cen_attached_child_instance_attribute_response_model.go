// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenAttachedChildInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenAttachedChildInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenAttachedChildInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenAttachedChildInstanceAttributeResponseBody) *DescribeCenAttachedChildInstanceAttributeResponse
	GetBody() *DescribeCenAttachedChildInstanceAttributeResponseBody
}

type DescribeCenAttachedChildInstanceAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenAttachedChildInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenAttachedChildInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenAttachedChildInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenAttachedChildInstanceAttributeResponse) GetBody() *DescribeCenAttachedChildInstanceAttributeResponseBody {
	return s.Body
}

func (s *DescribeCenAttachedChildInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeCenAttachedChildInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponse) SetStatusCode(v int32) *DescribeCenAttachedChildInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponse) SetBody(v *DescribeCenAttachedChildInstanceAttributeResponseBody) *DescribeCenAttachedChildInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeCenAttachedChildInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
