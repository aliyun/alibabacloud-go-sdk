// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfuseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePxfuseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePxfuseInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribePxfuseInfoResponseBody) *DescribePxfuseInfoResponse
	GetBody() *DescribePxfuseInfoResponseBody
}

type DescribePxfuseInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePxfuseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePxfuseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribePxfuseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePxfuseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePxfuseInfoResponse) GetBody() *DescribePxfuseInfoResponseBody {
	return s.Body
}

func (s *DescribePxfuseInfoResponse) SetHeaders(v map[string]*string) *DescribePxfuseInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribePxfuseInfoResponse) SetStatusCode(v int32) *DescribePxfuseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePxfuseInfoResponse) SetBody(v *DescribePxfuseInfoResponseBody) *DescribePxfuseInfoResponse {
	s.Body = v
	return s
}

func (s *DescribePxfuseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
