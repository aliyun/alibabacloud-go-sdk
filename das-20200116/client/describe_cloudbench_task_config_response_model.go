// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudbenchTaskConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudbenchTaskConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudbenchTaskConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudbenchTaskConfigResponseBody) *DescribeCloudbenchTaskConfigResponse
	GetBody() *DescribeCloudbenchTaskConfigResponseBody
}

type DescribeCloudbenchTaskConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudbenchTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudbenchTaskConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudbenchTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudbenchTaskConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudbenchTaskConfigResponse) GetBody() *DescribeCloudbenchTaskConfigResponseBody {
	return s.Body
}

func (s *DescribeCloudbenchTaskConfigResponse) SetHeaders(v map[string]*string) *DescribeCloudbenchTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponse) SetStatusCode(v int32) *DescribeCloudbenchTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponse) SetBody(v *DescribeCloudbenchTaskConfigResponseBody) *DescribeCloudbenchTaskConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudbenchTaskConfigResponse) Validate() error {
	return dara.Validate(s)
}
