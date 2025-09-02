// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalcEnginesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCalcEnginesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCalcEnginesResponse
	GetStatusCode() *int32
	SetBody(v *ListCalcEnginesResponseBody) *ListCalcEnginesResponse
	GetBody() *ListCalcEnginesResponseBody
}

type ListCalcEnginesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCalcEnginesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCalcEnginesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCalcEnginesResponse) GoString() string {
	return s.String()
}

func (s *ListCalcEnginesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCalcEnginesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCalcEnginesResponse) GetBody() *ListCalcEnginesResponseBody {
	return s.Body
}

func (s *ListCalcEnginesResponse) SetHeaders(v map[string]*string) *ListCalcEnginesResponse {
	s.Headers = v
	return s
}

func (s *ListCalcEnginesResponse) SetStatusCode(v int32) *ListCalcEnginesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCalcEnginesResponse) SetBody(v *ListCalcEnginesResponseBody) *ListCalcEnginesResponse {
	s.Body = v
	return s
}

func (s *ListCalcEnginesResponse) Validate() error {
	return dara.Validate(s)
}
