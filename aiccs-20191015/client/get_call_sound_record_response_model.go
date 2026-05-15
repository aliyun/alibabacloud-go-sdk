// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallSoundRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCallSoundRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCallSoundRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetCallSoundRecordResponseBody) *GetCallSoundRecordResponse
	GetBody() *GetCallSoundRecordResponseBody
}

type GetCallSoundRecordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCallSoundRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCallSoundRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCallSoundRecordResponse) GoString() string {
	return s.String()
}

func (s *GetCallSoundRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCallSoundRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCallSoundRecordResponse) GetBody() *GetCallSoundRecordResponseBody {
	return s.Body
}

func (s *GetCallSoundRecordResponse) SetHeaders(v map[string]*string) *GetCallSoundRecordResponse {
	s.Headers = v
	return s
}

func (s *GetCallSoundRecordResponse) SetStatusCode(v int32) *GetCallSoundRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallSoundRecordResponse) SetBody(v *GetCallSoundRecordResponseBody) *GetCallSoundRecordResponse {
	s.Body = v
	return s
}

func (s *GetCallSoundRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
