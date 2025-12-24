// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarTaskAndActionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarTaskAndActionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarTaskAndActionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarTaskAndActionsResponseBody) *DescribeSoarTaskAndActionsResponse
	GetBody() *DescribeSoarTaskAndActionsResponseBody
}

type DescribeSoarTaskAndActionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarTaskAndActionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarTaskAndActionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarTaskAndActionsResponse) GetBody() *DescribeSoarTaskAndActionsResponseBody {
	return s.Body
}

func (s *DescribeSoarTaskAndActionsResponse) SetHeaders(v map[string]*string) *DescribeSoarTaskAndActionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarTaskAndActionsResponse) SetStatusCode(v int32) *DescribeSoarTaskAndActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponse) SetBody(v *DescribeSoarTaskAndActionsResponseBody) *DescribeSoarTaskAndActionsResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarTaskAndActionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
