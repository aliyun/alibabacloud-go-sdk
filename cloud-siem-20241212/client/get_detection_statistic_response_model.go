// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectionStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDetectionStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDetectionStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetDetectionStatisticResponseBody) *GetDetectionStatisticResponse
	GetBody() *GetDetectionStatisticResponseBody
}

type GetDetectionStatisticResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetectionStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetectionStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDetectionStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetDetectionStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDetectionStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDetectionStatisticResponse) GetBody() *GetDetectionStatisticResponseBody {
	return s.Body
}

func (s *GetDetectionStatisticResponse) SetHeaders(v map[string]*string) *GetDetectionStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetDetectionStatisticResponse) SetStatusCode(v int32) *GetDetectionStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetectionStatisticResponse) SetBody(v *GetDetectionStatisticResponseBody) *GetDetectionStatisticResponse {
	s.Body = v
	return s
}

func (s *GetDetectionStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
