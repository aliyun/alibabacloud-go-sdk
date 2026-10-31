// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamOpsRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStreamOpsRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStreamOpsRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetStreamOpsRecordResponseBody) *GetStreamOpsRecordResponse
	GetBody() *GetStreamOpsRecordResponseBody
}

type GetStreamOpsRecordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStreamOpsRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStreamOpsRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStreamOpsRecordResponse) GoString() string {
	return s.String()
}

func (s *GetStreamOpsRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStreamOpsRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStreamOpsRecordResponse) GetBody() *GetStreamOpsRecordResponseBody {
	return s.Body
}

func (s *GetStreamOpsRecordResponse) SetHeaders(v map[string]*string) *GetStreamOpsRecordResponse {
	s.Headers = v
	return s
}

func (s *GetStreamOpsRecordResponse) SetStatusCode(v int32) *GetStreamOpsRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStreamOpsRecordResponse) SetBody(v *GetStreamOpsRecordResponseBody) *GetStreamOpsRecordResponse {
	s.Body = v
	return s
}

func (s *GetStreamOpsRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
