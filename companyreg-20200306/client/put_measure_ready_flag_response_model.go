// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMeasureReadyFlagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutMeasureReadyFlagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutMeasureReadyFlagResponse
	GetStatusCode() *int32
	SetBody(v *PutMeasureReadyFlagResponseBody) *PutMeasureReadyFlagResponse
	GetBody() *PutMeasureReadyFlagResponseBody
}

type PutMeasureReadyFlagResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutMeasureReadyFlagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutMeasureReadyFlagResponse) String() string {
	return dara.Prettify(s)
}

func (s PutMeasureReadyFlagResponse) GoString() string {
	return s.String()
}

func (s *PutMeasureReadyFlagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutMeasureReadyFlagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutMeasureReadyFlagResponse) GetBody() *PutMeasureReadyFlagResponseBody {
	return s.Body
}

func (s *PutMeasureReadyFlagResponse) SetHeaders(v map[string]*string) *PutMeasureReadyFlagResponse {
	s.Headers = v
	return s
}

func (s *PutMeasureReadyFlagResponse) SetStatusCode(v int32) *PutMeasureReadyFlagResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMeasureReadyFlagResponse) SetBody(v *PutMeasureReadyFlagResponseBody) *PutMeasureReadyFlagResponse {
	s.Body = v
	return s
}

func (s *PutMeasureReadyFlagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
