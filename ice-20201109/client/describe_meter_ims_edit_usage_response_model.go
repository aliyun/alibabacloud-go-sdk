// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsEditUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMeterImsEditUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMeterImsEditUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMeterImsEditUsageResponseBody) *DescribeMeterImsEditUsageResponse
	GetBody() *DescribeMeterImsEditUsageResponseBody
}

type DescribeMeterImsEditUsageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterImsEditUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterImsEditUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsEditUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsEditUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMeterImsEditUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMeterImsEditUsageResponse) GetBody() *DescribeMeterImsEditUsageResponseBody {
	return s.Body
}

func (s *DescribeMeterImsEditUsageResponse) SetHeaders(v map[string]*string) *DescribeMeterImsEditUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterImsEditUsageResponse) SetStatusCode(v int32) *DescribeMeterImsEditUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterImsEditUsageResponse) SetBody(v *DescribeMeterImsEditUsageResponseBody) *DescribeMeterImsEditUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeMeterImsEditUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
