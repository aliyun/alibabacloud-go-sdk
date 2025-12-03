// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkitemRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkitemRecordResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkitemRecordResponseBody) *CreateWorkitemRecordResponse
	GetBody() *CreateWorkitemRecordResponseBody
}

type CreateWorkitemRecordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkitemRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkitemRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkitemRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkitemRecordResponse) GetBody() *CreateWorkitemRecordResponseBody {
	return s.Body
}

func (s *CreateWorkitemRecordResponse) SetHeaders(v map[string]*string) *CreateWorkitemRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkitemRecordResponse) SetStatusCode(v int32) *CreateWorkitemRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkitemRecordResponse) SetBody(v *CreateWorkitemRecordResponseBody) *CreateWorkitemRecordResponse {
	s.Body = v
	return s
}

func (s *CreateWorkitemRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
