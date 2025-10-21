// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelInputContentAsyncDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelInputContentAsyncDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelInputContentAsyncDetectResponse
	GetStatusCode() *int32
	SetBody(v *ModelInputContentAsyncDetectResponseBody) *ModelInputContentAsyncDetectResponse
	GetBody() *ModelInputContentAsyncDetectResponseBody
}

type ModelInputContentAsyncDetectResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelInputContentAsyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelInputContentAsyncDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentAsyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ModelInputContentAsyncDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelInputContentAsyncDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelInputContentAsyncDetectResponse) GetBody() *ModelInputContentAsyncDetectResponseBody {
	return s.Body
}

func (s *ModelInputContentAsyncDetectResponse) SetHeaders(v map[string]*string) *ModelInputContentAsyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ModelInputContentAsyncDetectResponse) SetStatusCode(v int32) *ModelInputContentAsyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponse) SetBody(v *ModelInputContentAsyncDetectResponseBody) *ModelInputContentAsyncDetectResponse {
	s.Body = v
	return s
}

func (s *ModelInputContentAsyncDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
