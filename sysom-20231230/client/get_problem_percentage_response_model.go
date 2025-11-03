// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProblemPercentageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProblemPercentageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProblemPercentageResponse
	GetStatusCode() *int32
	SetBody(v *GetProblemPercentageResponseBody) *GetProblemPercentageResponse
	GetBody() *GetProblemPercentageResponseBody
}

type GetProblemPercentageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProblemPercentageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProblemPercentageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProblemPercentageResponse) GoString() string {
	return s.String()
}

func (s *GetProblemPercentageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProblemPercentageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProblemPercentageResponse) GetBody() *GetProblemPercentageResponseBody {
	return s.Body
}

func (s *GetProblemPercentageResponse) SetHeaders(v map[string]*string) *GetProblemPercentageResponse {
	s.Headers = v
	return s
}

func (s *GetProblemPercentageResponse) SetStatusCode(v int32) *GetProblemPercentageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProblemPercentageResponse) SetBody(v *GetProblemPercentageResponseBody) *GetProblemPercentageResponse {
	s.Body = v
	return s
}

func (s *GetProblemPercentageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
