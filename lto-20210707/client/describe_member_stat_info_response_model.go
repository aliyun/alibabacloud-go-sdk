// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberStatInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMemberStatInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMemberStatInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMemberStatInfoResponseBody) *DescribeMemberStatInfoResponse
	GetBody() *DescribeMemberStatInfoResponseBody
}

type DescribeMemberStatInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMemberStatInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMemberStatInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberStatInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeMemberStatInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMemberStatInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMemberStatInfoResponse) GetBody() *DescribeMemberStatInfoResponseBody {
	return s.Body
}

func (s *DescribeMemberStatInfoResponse) SetHeaders(v map[string]*string) *DescribeMemberStatInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeMemberStatInfoResponse) SetStatusCode(v int32) *DescribeMemberStatInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMemberStatInfoResponse) SetBody(v *DescribeMemberStatInfoResponseBody) *DescribeMemberStatInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeMemberStatInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
