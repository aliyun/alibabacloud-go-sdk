// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOperationRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOperationRecordsResponse
	GetStatusCode() *int32
	SetBody(v *GetOperationRecordsResponseBody) *GetOperationRecordsResponse
	GetBody() *GetOperationRecordsResponseBody
}

type GetOperationRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOperationRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOperationRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordsResponse) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOperationRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOperationRecordsResponse) GetBody() *GetOperationRecordsResponseBody {
	return s.Body
}

func (s *GetOperationRecordsResponse) SetHeaders(v map[string]*string) *GetOperationRecordsResponse {
	s.Headers = v
	return s
}

func (s *GetOperationRecordsResponse) SetStatusCode(v int32) *GetOperationRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationRecordsResponse) SetBody(v *GetOperationRecordsResponseBody) *GetOperationRecordsResponse {
	s.Body = v
	return s
}

func (s *GetOperationRecordsResponse) Validate() error {
	return dara.Validate(s)
}
