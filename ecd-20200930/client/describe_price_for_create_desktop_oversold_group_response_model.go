// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceForCreateDesktopOversoldGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePriceForCreateDesktopOversoldGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePriceForCreateDesktopOversoldGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribePriceForCreateDesktopOversoldGroupResponseBody) *DescribePriceForCreateDesktopOversoldGroupResponse
	GetBody() *DescribePriceForCreateDesktopOversoldGroupResponseBody
}

type DescribePriceForCreateDesktopOversoldGroupResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePriceForCreateDesktopOversoldGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePriceForCreateDesktopOversoldGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForCreateDesktopOversoldGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponse) GetBody() *DescribePriceForCreateDesktopOversoldGroupResponseBody {
	return s.Body
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponse) SetHeaders(v map[string]*string) *DescribePriceForCreateDesktopOversoldGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponse) SetStatusCode(v int32) *DescribePriceForCreateDesktopOversoldGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponse) SetBody(v *DescribePriceForCreateDesktopOversoldGroupResponseBody) *DescribePriceForCreateDesktopOversoldGroupResponse {
	s.Body = v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponse) Validate() error {
	return dara.Validate(s)
}
