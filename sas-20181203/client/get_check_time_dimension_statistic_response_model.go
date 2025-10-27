// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckTimeDimensionStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckTimeDimensionStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckTimeDimensionStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckTimeDimensionStatisticResponseBody) *GetCheckTimeDimensionStatisticResponse
	GetBody() *GetCheckTimeDimensionStatisticResponseBody
}

type GetCheckTimeDimensionStatisticResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckTimeDimensionStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckTimeDimensionStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckTimeDimensionStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetCheckTimeDimensionStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckTimeDimensionStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckTimeDimensionStatisticResponse) GetBody() *GetCheckTimeDimensionStatisticResponseBody {
	return s.Body
}

func (s *GetCheckTimeDimensionStatisticResponse) SetHeaders(v map[string]*string) *GetCheckTimeDimensionStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetCheckTimeDimensionStatisticResponse) SetStatusCode(v int32) *GetCheckTimeDimensionStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckTimeDimensionStatisticResponse) SetBody(v *GetCheckTimeDimensionStatisticResponseBody) *GetCheckTimeDimensionStatisticResponse {
	s.Body = v
	return s
}

func (s *GetCheckTimeDimensionStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
