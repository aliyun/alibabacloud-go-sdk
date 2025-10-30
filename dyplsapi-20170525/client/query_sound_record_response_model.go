// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySoundRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySoundRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySoundRecordResponse
	GetStatusCode() *int32
	SetBody(v *QuerySoundRecordResponseBody) *QuerySoundRecordResponse
	GetBody() *QuerySoundRecordResponseBody
}

type QuerySoundRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySoundRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySoundRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySoundRecordResponse) GoString() string {
	return s.String()
}

func (s *QuerySoundRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySoundRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySoundRecordResponse) GetBody() *QuerySoundRecordResponseBody {
	return s.Body
}

func (s *QuerySoundRecordResponse) SetHeaders(v map[string]*string) *QuerySoundRecordResponse {
	s.Headers = v
	return s
}

func (s *QuerySoundRecordResponse) SetStatusCode(v int32) *QuerySoundRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySoundRecordResponse) SetBody(v *QuerySoundRecordResponseBody) *QuerySoundRecordResponse {
	s.Body = v
	return s
}

func (s *QuerySoundRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
