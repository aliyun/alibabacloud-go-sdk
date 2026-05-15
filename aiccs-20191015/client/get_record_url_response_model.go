// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecordUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecordUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetRecordUrlResponseBody) *GetRecordUrlResponse
	GetBody() *GetRecordUrlResponseBody
}

type GetRecordUrlResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecordUrlResponse) GoString() string {
	return s.String()
}

func (s *GetRecordUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecordUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecordUrlResponse) GetBody() *GetRecordUrlResponseBody {
	return s.Body
}

func (s *GetRecordUrlResponse) SetHeaders(v map[string]*string) *GetRecordUrlResponse {
	s.Headers = v
	return s
}

func (s *GetRecordUrlResponse) SetStatusCode(v int32) *GetRecordUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordUrlResponse) SetBody(v *GetRecordUrlResponseBody) *GetRecordUrlResponse {
	s.Body = v
	return s
}

func (s *GetRecordUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
