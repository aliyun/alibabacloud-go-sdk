// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSampleDataByTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSampleDataByTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSampleDataByTextResponse
	GetStatusCode() *int32
	SetBody(v *AddSampleDataByTextResponseBody) *AddSampleDataByTextResponse
	GetBody() *AddSampleDataByTextResponseBody
}

type AddSampleDataByTextResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSampleDataByTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSampleDataByTextResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSampleDataByTextResponse) GoString() string {
	return s.String()
}

func (s *AddSampleDataByTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSampleDataByTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSampleDataByTextResponse) GetBody() *AddSampleDataByTextResponseBody {
	return s.Body
}

func (s *AddSampleDataByTextResponse) SetHeaders(v map[string]*string) *AddSampleDataByTextResponse {
	s.Headers = v
	return s
}

func (s *AddSampleDataByTextResponse) SetStatusCode(v int32) *AddSampleDataByTextResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSampleDataByTextResponse) SetBody(v *AddSampleDataByTextResponseBody) *AddSampleDataByTextResponse {
	s.Body = v
	return s
}

func (s *AddSampleDataByTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
