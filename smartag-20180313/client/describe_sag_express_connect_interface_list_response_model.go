// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagExpressConnectInterfaceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagExpressConnectInterfaceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagExpressConnectInterfaceListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagExpressConnectInterfaceListResponseBody) *DescribeSagExpressConnectInterfaceListResponse
	GetBody() *DescribeSagExpressConnectInterfaceListResponseBody
}

type DescribeSagExpressConnectInterfaceListResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagExpressConnectInterfaceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagExpressConnectInterfaceListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagExpressConnectInterfaceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagExpressConnectInterfaceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagExpressConnectInterfaceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagExpressConnectInterfaceListResponse) GetBody() *DescribeSagExpressConnectInterfaceListResponseBody {
	return s.Body
}

func (s *DescribeSagExpressConnectInterfaceListResponse) SetHeaders(v map[string]*string) *DescribeSagExpressConnectInterfaceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponse) SetStatusCode(v int32) *DescribeSagExpressConnectInterfaceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponse) SetBody(v *DescribeSagExpressConnectInterfaceListResponseBody) *DescribeSagExpressConnectInterfaceListResponse {
	s.Body = v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
