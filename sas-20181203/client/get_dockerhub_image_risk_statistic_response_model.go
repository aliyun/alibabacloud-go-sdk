// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerhubImageRiskStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDockerhubImageRiskStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDockerhubImageRiskStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetDockerhubImageRiskStatisticResponseBody) *GetDockerhubImageRiskStatisticResponse
	GetBody() *GetDockerhubImageRiskStatisticResponseBody
}

type GetDockerhubImageRiskStatisticResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDockerhubImageRiskStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDockerhubImageRiskStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDockerhubImageRiskStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetDockerhubImageRiskStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDockerhubImageRiskStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDockerhubImageRiskStatisticResponse) GetBody() *GetDockerhubImageRiskStatisticResponseBody {
	return s.Body
}

func (s *GetDockerhubImageRiskStatisticResponse) SetHeaders(v map[string]*string) *GetDockerhubImageRiskStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetDockerhubImageRiskStatisticResponse) SetStatusCode(v int32) *GetDockerhubImageRiskStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDockerhubImageRiskStatisticResponse) SetBody(v *GetDockerhubImageRiskStatisticResponseBody) *GetDockerhubImageRiskStatisticResponse {
	s.Body = v
	return s
}

func (s *GetDockerhubImageRiskStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
