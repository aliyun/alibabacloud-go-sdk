// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRecordContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRecordContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRecordContentResponse
	GetStatusCode() *int32
	SetBody(v *StartRecordContentResponseBody) *StartRecordContentResponse
	GetBody() *StartRecordContentResponseBody
}

type StartRecordContentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRecordContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRecordContentResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRecordContentResponse) GoString() string {
	return s.String()
}

func (s *StartRecordContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRecordContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRecordContentResponse) GetBody() *StartRecordContentResponseBody {
	return s.Body
}

func (s *StartRecordContentResponse) SetHeaders(v map[string]*string) *StartRecordContentResponse {
	s.Headers = v
	return s
}

func (s *StartRecordContentResponse) SetStatusCode(v int32) *StartRecordContentResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRecordContentResponse) SetBody(v *StartRecordContentResponseBody) *StartRecordContentResponse {
	s.Body = v
	return s
}

func (s *StartRecordContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
