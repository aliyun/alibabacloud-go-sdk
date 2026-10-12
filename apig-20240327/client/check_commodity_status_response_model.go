// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCommodityStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCommodityStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCommodityStatusResponse
	GetStatusCode() *int32
	SetBody(v *CheckCommodityStatusResponseBody) *CheckCommodityStatusResponse
	GetBody() *CheckCommodityStatusResponseBody
}

type CheckCommodityStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCommodityStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCommodityStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCommodityStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckCommodityStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCommodityStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCommodityStatusResponse) GetBody() *CheckCommodityStatusResponseBody {
	return s.Body
}

func (s *CheckCommodityStatusResponse) SetHeaders(v map[string]*string) *CheckCommodityStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckCommodityStatusResponse) SetStatusCode(v int32) *CheckCommodityStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCommodityStatusResponse) SetBody(v *CheckCommodityStatusResponseBody) *CheckCommodityStatusResponse {
	s.Body = v
	return s
}

func (s *CheckCommodityStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
