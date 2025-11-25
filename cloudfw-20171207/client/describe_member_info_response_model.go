// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMemberInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMemberInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMemberInfoResponseBody) *DescribeMemberInfoResponse
	GetBody() *DescribeMemberInfoResponseBody
}

type DescribeMemberInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMemberInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMemberInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeMemberInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMemberInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMemberInfoResponse) GetBody() *DescribeMemberInfoResponseBody {
	return s.Body
}

func (s *DescribeMemberInfoResponse) SetHeaders(v map[string]*string) *DescribeMemberInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeMemberInfoResponse) SetStatusCode(v int32) *DescribeMemberInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMemberInfoResponse) SetBody(v *DescribeMemberInfoResponseBody) *DescribeMemberInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeMemberInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
