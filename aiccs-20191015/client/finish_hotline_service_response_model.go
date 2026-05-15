// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishHotlineServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FinishHotlineServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FinishHotlineServiceResponse
	GetStatusCode() *int32
	SetBody(v *FinishHotlineServiceResponseBody) *FinishHotlineServiceResponse
	GetBody() *FinishHotlineServiceResponseBody
}

type FinishHotlineServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishHotlineServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s FinishHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FinishHotlineServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FinishHotlineServiceResponse) GetBody() *FinishHotlineServiceResponseBody {
	return s.Body
}

func (s *FinishHotlineServiceResponse) SetHeaders(v map[string]*string) *FinishHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *FinishHotlineServiceResponse) SetStatusCode(v int32) *FinishHotlineServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishHotlineServiceResponse) SetBody(v *FinishHotlineServiceResponseBody) *FinishHotlineServiceResponse {
	s.Body = v
	return s
}

func (s *FinishHotlineServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
