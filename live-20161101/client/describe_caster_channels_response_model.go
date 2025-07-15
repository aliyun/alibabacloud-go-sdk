// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCasterChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCasterChannelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCasterChannelsResponseBody) *DescribeCasterChannelsResponse
	GetBody() *DescribeCasterChannelsResponseBody
}

type DescribeCasterChannelsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCasterChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCasterChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterChannelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCasterChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCasterChannelsResponse) GetBody() *DescribeCasterChannelsResponseBody {
	return s.Body
}

func (s *DescribeCasterChannelsResponse) SetHeaders(v map[string]*string) *DescribeCasterChannelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCasterChannelsResponse) SetStatusCode(v int32) *DescribeCasterChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCasterChannelsResponse) SetBody(v *DescribeCasterChannelsResponseBody) *DescribeCasterChannelsResponse {
	s.Body = v
	return s
}

func (s *DescribeCasterChannelsResponse) Validate() error {
	return dara.Validate(s)
}
