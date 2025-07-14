// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataInterpretationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DataInterpretationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DataInterpretationResponse
	GetStatusCode() *int32
	SetBody(v *DataInterpretationResponseBody) *DataInterpretationResponse
	GetBody() *DataInterpretationResponseBody
}

type DataInterpretationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DataInterpretationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DataInterpretationResponse) String() string {
	return dara.Prettify(s)
}

func (s DataInterpretationResponse) GoString() string {
	return s.String()
}

func (s *DataInterpretationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DataInterpretationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DataInterpretationResponse) GetBody() *DataInterpretationResponseBody {
	return s.Body
}

func (s *DataInterpretationResponse) SetHeaders(v map[string]*string) *DataInterpretationResponse {
	s.Headers = v
	return s
}

func (s *DataInterpretationResponse) SetStatusCode(v int32) *DataInterpretationResponse {
	s.StatusCode = &v
	return s
}

func (s *DataInterpretationResponse) SetBody(v *DataInterpretationResponseBody) *DataInterpretationResponse {
	s.Body = v
	return s
}

func (s *DataInterpretationResponse) Validate() error {
	return dara.Validate(s)
}
