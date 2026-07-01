// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOperationRecordByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOperationRecordByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetOperationRecordByIdResponseBody) *GetOperationRecordByIdResponse
	GetBody() *GetOperationRecordByIdResponseBody
}

type GetOperationRecordByIdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOperationRecordByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOperationRecordByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordByIdResponse) GoString() string {
	return s.String()
}

func (s *GetOperationRecordByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOperationRecordByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOperationRecordByIdResponse) GetBody() *GetOperationRecordByIdResponseBody {
	return s.Body
}

func (s *GetOperationRecordByIdResponse) SetHeaders(v map[string]*string) *GetOperationRecordByIdResponse {
	s.Headers = v
	return s
}

func (s *GetOperationRecordByIdResponse) SetStatusCode(v int32) *GetOperationRecordByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationRecordByIdResponse) SetBody(v *GetOperationRecordByIdResponseBody) *GetOperationRecordByIdResponse {
	s.Body = v
	return s
}

func (s *GetOperationRecordByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
