// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelOutputContentAsyncDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelOutputContentAsyncDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelOutputContentAsyncDetectResponse
	GetStatusCode() *int32
	SetBody(v *ModelOutputContentAsyncDetectResponseBody) *ModelOutputContentAsyncDetectResponse
	GetBody() *ModelOutputContentAsyncDetectResponseBody
}

type ModelOutputContentAsyncDetectResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelOutputContentAsyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelOutputContentAsyncDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentAsyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ModelOutputContentAsyncDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelOutputContentAsyncDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelOutputContentAsyncDetectResponse) GetBody() *ModelOutputContentAsyncDetectResponseBody {
	return s.Body
}

func (s *ModelOutputContentAsyncDetectResponse) SetHeaders(v map[string]*string) *ModelOutputContentAsyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ModelOutputContentAsyncDetectResponse) SetStatusCode(v int32) *ModelOutputContentAsyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponse) SetBody(v *ModelOutputContentAsyncDetectResponseBody) *ModelOutputContentAsyncDetectResponse {
	s.Body = v
	return s
}

func (s *ModelOutputContentAsyncDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
