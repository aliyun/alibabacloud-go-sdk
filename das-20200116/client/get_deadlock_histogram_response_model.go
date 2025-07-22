// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadlockHistogramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeadlockHistogramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeadlockHistogramResponse
	GetStatusCode() *int32
	SetBody(v *GetDeadlockHistogramResponseBody) *GetDeadlockHistogramResponse
	GetBody() *GetDeadlockHistogramResponseBody
}

type GetDeadlockHistogramResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeadlockHistogramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeadlockHistogramResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeadlockHistogramResponse) GoString() string {
	return s.String()
}

func (s *GetDeadlockHistogramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeadlockHistogramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeadlockHistogramResponse) GetBody() *GetDeadlockHistogramResponseBody {
	return s.Body
}

func (s *GetDeadlockHistogramResponse) SetHeaders(v map[string]*string) *GetDeadlockHistogramResponse {
	s.Headers = v
	return s
}

func (s *GetDeadlockHistogramResponse) SetStatusCode(v int32) *GetDeadlockHistogramResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeadlockHistogramResponse) SetBody(v *GetDeadlockHistogramResponseBody) *GetDeadlockHistogramResponse {
	s.Body = v
	return s
}

func (s *GetDeadlockHistogramResponse) Validate() error {
	return dara.Validate(s)
}
