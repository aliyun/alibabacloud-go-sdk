// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadSampleDataSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LoadSampleDataSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LoadSampleDataSetResponse
	GetStatusCode() *int32
	SetBody(v *LoadSampleDataSetResponseBody) *LoadSampleDataSetResponse
	GetBody() *LoadSampleDataSetResponseBody
}

type LoadSampleDataSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LoadSampleDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoadSampleDataSetResponse) String() string {
	return dara.Prettify(s)
}

func (s LoadSampleDataSetResponse) GoString() string {
	return s.String()
}

func (s *LoadSampleDataSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LoadSampleDataSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LoadSampleDataSetResponse) GetBody() *LoadSampleDataSetResponseBody {
	return s.Body
}

func (s *LoadSampleDataSetResponse) SetHeaders(v map[string]*string) *LoadSampleDataSetResponse {
	s.Headers = v
	return s
}

func (s *LoadSampleDataSetResponse) SetStatusCode(v int32) *LoadSampleDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *LoadSampleDataSetResponse) SetBody(v *LoadSampleDataSetResponseBody) *LoadSampleDataSetResponse {
	s.Body = v
	return s
}

func (s *LoadSampleDataSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
