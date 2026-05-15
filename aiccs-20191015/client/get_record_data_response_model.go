// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecordDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecordDataResponse
	GetStatusCode() *int32
	SetBody(v *GetRecordDataResponseBody) *GetRecordDataResponse
	GetBody() *GetRecordDataResponseBody
}

type GetRecordDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecordDataResponse) GoString() string {
	return s.String()
}

func (s *GetRecordDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecordDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecordDataResponse) GetBody() *GetRecordDataResponseBody {
	return s.Body
}

func (s *GetRecordDataResponse) SetHeaders(v map[string]*string) *GetRecordDataResponse {
	s.Headers = v
	return s
}

func (s *GetRecordDataResponse) SetStatusCode(v int32) *GetRecordDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordDataResponse) SetBody(v *GetRecordDataResponseBody) *GetRecordDataResponse {
	s.Body = v
	return s
}

func (s *GetRecordDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
