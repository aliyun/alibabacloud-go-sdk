// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerAccessOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsumerAccessOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsumerAccessOverviewResponse
	GetStatusCode() *int32
	SetBody(v *GetConsumerAccessOverviewResponseBody) *GetConsumerAccessOverviewResponse
	GetBody() *GetConsumerAccessOverviewResponseBody
}

type GetConsumerAccessOverviewResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerAccessOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerAccessOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerAccessOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerAccessOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsumerAccessOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsumerAccessOverviewResponse) GetBody() *GetConsumerAccessOverviewResponseBody {
	return s.Body
}

func (s *GetConsumerAccessOverviewResponse) SetHeaders(v map[string]*string) *GetConsumerAccessOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerAccessOverviewResponse) SetStatusCode(v int32) *GetConsumerAccessOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerAccessOverviewResponse) SetBody(v *GetConsumerAccessOverviewResponseBody) *GetConsumerAccessOverviewResponse {
	s.Body = v
	return s
}

func (s *GetConsumerAccessOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
