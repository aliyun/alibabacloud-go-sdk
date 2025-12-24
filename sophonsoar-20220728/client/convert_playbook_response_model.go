// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *ConvertPlaybookResponseBody) *ConvertPlaybookResponse
	GetBody() *ConvertPlaybookResponseBody
}

type ConvertPlaybookResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertPlaybookResponse) GoString() string {
	return s.String()
}

func (s *ConvertPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertPlaybookResponse) GetBody() *ConvertPlaybookResponseBody {
	return s.Body
}

func (s *ConvertPlaybookResponse) SetHeaders(v map[string]*string) *ConvertPlaybookResponse {
	s.Headers = v
	return s
}

func (s *ConvertPlaybookResponse) SetStatusCode(v int32) *ConvertPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertPlaybookResponse) SetBody(v *ConvertPlaybookResponseBody) *ConvertPlaybookResponse {
	s.Body = v
	return s
}

func (s *ConvertPlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
