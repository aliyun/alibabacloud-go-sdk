// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceRecordResponseBody) *GetServiceRecordResponse
	GetBody() *GetServiceRecordResponseBody
}

type GetServiceRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRecordResponse) GoString() string {
	return s.String()
}

func (s *GetServiceRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceRecordResponse) GetBody() *GetServiceRecordResponseBody {
	return s.Body
}

func (s *GetServiceRecordResponse) SetHeaders(v map[string]*string) *GetServiceRecordResponse {
	s.Headers = v
	return s
}

func (s *GetServiceRecordResponse) SetStatusCode(v int32) *GetServiceRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceRecordResponse) SetBody(v *GetServiceRecordResponseBody) *GetServiceRecordResponse {
	s.Body = v
	return s
}

func (s *GetServiceRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
