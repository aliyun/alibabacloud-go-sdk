// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuousAdjustResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContinuousAdjustResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContinuousAdjustResponse
	GetStatusCode() *int32
	SetBody(v *ContinuousAdjustResponseBody) *ContinuousAdjustResponse
	GetBody() *ContinuousAdjustResponseBody
}

type ContinuousAdjustResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinuousAdjustResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinuousAdjustResponse) String() string {
	return dara.Prettify(s)
}

func (s ContinuousAdjustResponse) GoString() string {
	return s.String()
}

func (s *ContinuousAdjustResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContinuousAdjustResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContinuousAdjustResponse) GetBody() *ContinuousAdjustResponseBody {
	return s.Body
}

func (s *ContinuousAdjustResponse) SetHeaders(v map[string]*string) *ContinuousAdjustResponse {
	s.Headers = v
	return s
}

func (s *ContinuousAdjustResponse) SetStatusCode(v int32) *ContinuousAdjustResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinuousAdjustResponse) SetBody(v *ContinuousAdjustResponseBody) *ContinuousAdjustResponse {
	s.Body = v
	return s
}

func (s *ContinuousAdjustResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
