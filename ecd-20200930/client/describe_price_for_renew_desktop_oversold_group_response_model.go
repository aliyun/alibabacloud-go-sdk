// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceForRenewDesktopOversoldGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePriceForRenewDesktopOversoldGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePriceForRenewDesktopOversoldGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribePriceForRenewDesktopOversoldGroupResponseBody) *DescribePriceForRenewDesktopOversoldGroupResponse
	GetBody() *DescribePriceForRenewDesktopOversoldGroupResponseBody
}

type DescribePriceForRenewDesktopOversoldGroupResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePriceForRenewDesktopOversoldGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePriceForRenewDesktopOversoldGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForRenewDesktopOversoldGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponse) GetBody() *DescribePriceForRenewDesktopOversoldGroupResponseBody {
	return s.Body
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponse) SetHeaders(v map[string]*string) *DescribePriceForRenewDesktopOversoldGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponse) SetStatusCode(v int32) *DescribePriceForRenewDesktopOversoldGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponse) SetBody(v *DescribePriceForRenewDesktopOversoldGroupResponseBody) *DescribePriceForRenewDesktopOversoldGroupResponse {
	s.Body = v
	return s
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
