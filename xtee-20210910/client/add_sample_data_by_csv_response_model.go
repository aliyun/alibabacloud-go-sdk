// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSampleDataByCsvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSampleDataByCsvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSampleDataByCsvResponse
	GetStatusCode() *int32
	SetBody(v *AddSampleDataByCsvResponseBody) *AddSampleDataByCsvResponse
	GetBody() *AddSampleDataByCsvResponseBody
}

type AddSampleDataByCsvResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSampleDataByCsvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSampleDataByCsvResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSampleDataByCsvResponse) GoString() string {
	return s.String()
}

func (s *AddSampleDataByCsvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSampleDataByCsvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSampleDataByCsvResponse) GetBody() *AddSampleDataByCsvResponseBody {
	return s.Body
}

func (s *AddSampleDataByCsvResponse) SetHeaders(v map[string]*string) *AddSampleDataByCsvResponse {
	s.Headers = v
	return s
}

func (s *AddSampleDataByCsvResponse) SetStatusCode(v int32) *AddSampleDataByCsvResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSampleDataByCsvResponse) SetBody(v *AddSampleDataByCsvResponseBody) *AddSampleDataByCsvResponse {
	s.Body = v
	return s
}

func (s *AddSampleDataByCsvResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
