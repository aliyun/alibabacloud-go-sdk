// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiccsSmartCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AiccsSmartCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AiccsSmartCallResponse
	GetStatusCode() *int32
	SetBody(v *AiccsSmartCallResponseBody) *AiccsSmartCallResponse
	GetBody() *AiccsSmartCallResponseBody
}

type AiccsSmartCallResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AiccsSmartCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AiccsSmartCallResponse) String() string {
	return dara.Prettify(s)
}

func (s AiccsSmartCallResponse) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AiccsSmartCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AiccsSmartCallResponse) GetBody() *AiccsSmartCallResponseBody {
	return s.Body
}

func (s *AiccsSmartCallResponse) SetHeaders(v map[string]*string) *AiccsSmartCallResponse {
	s.Headers = v
	return s
}

func (s *AiccsSmartCallResponse) SetStatusCode(v int32) *AiccsSmartCallResponse {
	s.StatusCode = &v
	return s
}

func (s *AiccsSmartCallResponse) SetBody(v *AiccsSmartCallResponseBody) *AiccsSmartCallResponse {
	s.Body = v
	return s
}

func (s *AiccsSmartCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
