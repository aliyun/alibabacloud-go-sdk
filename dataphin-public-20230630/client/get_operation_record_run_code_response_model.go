// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordRunCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOperationRecordRunCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOperationRecordRunCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetOperationRecordRunCodeResponseBody) *GetOperationRecordRunCodeResponse
	GetBody() *GetOperationRecordRunCodeResponseBody
}

type GetOperationRecordRunCodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOperationRecordRunCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOperationRecordRunCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordRunCodeResponse) GoString() string {
	return s.String()
}

func (s *GetOperationRecordRunCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOperationRecordRunCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOperationRecordRunCodeResponse) GetBody() *GetOperationRecordRunCodeResponseBody {
	return s.Body
}

func (s *GetOperationRecordRunCodeResponse) SetHeaders(v map[string]*string) *GetOperationRecordRunCodeResponse {
	s.Headers = v
	return s
}

func (s *GetOperationRecordRunCodeResponse) SetStatusCode(v int32) *GetOperationRecordRunCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationRecordRunCodeResponse) SetBody(v *GetOperationRecordRunCodeResponseBody) *GetOperationRecordRunCodeResponse {
	s.Body = v
	return s
}

func (s *GetOperationRecordRunCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
