// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberTotalStatInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMemberTotalStatInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMemberTotalStatInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMemberTotalStatInfoResponseBody) *DescribeMemberTotalStatInfoResponse
	GetBody() *DescribeMemberTotalStatInfoResponseBody
}

type DescribeMemberTotalStatInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMemberTotalStatInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMemberTotalStatInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberTotalStatInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeMemberTotalStatInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMemberTotalStatInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMemberTotalStatInfoResponse) GetBody() *DescribeMemberTotalStatInfoResponseBody {
	return s.Body
}

func (s *DescribeMemberTotalStatInfoResponse) SetHeaders(v map[string]*string) *DescribeMemberTotalStatInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeMemberTotalStatInfoResponse) SetStatusCode(v int32) *DescribeMemberTotalStatInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMemberTotalStatInfoResponse) SetBody(v *DescribeMemberTotalStatInfoResponseBody) *DescribeMemberTotalStatInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeMemberTotalStatInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
