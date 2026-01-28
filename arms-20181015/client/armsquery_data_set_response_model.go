// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iARMSQueryDataSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ARMSQueryDataSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ARMSQueryDataSetResponse
	GetStatusCode() *int32
	SetBody(v *ARMSQueryDataSetResponseBody) *ARMSQueryDataSetResponse
	GetBody() *ARMSQueryDataSetResponseBody
}

type ARMSQueryDataSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ARMSQueryDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ARMSQueryDataSetResponse) String() string {
	return dara.Prettify(s)
}

func (s ARMSQueryDataSetResponse) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ARMSQueryDataSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ARMSQueryDataSetResponse) GetBody() *ARMSQueryDataSetResponseBody {
	return s.Body
}

func (s *ARMSQueryDataSetResponse) SetHeaders(v map[string]*string) *ARMSQueryDataSetResponse {
	s.Headers = v
	return s
}

func (s *ARMSQueryDataSetResponse) SetStatusCode(v int32) *ARMSQueryDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ARMSQueryDataSetResponse) SetBody(v *ARMSQueryDataSetResponseBody) *ARMSQueryDataSetResponse {
	s.Body = v
	return s
}

func (s *ARMSQueryDataSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
