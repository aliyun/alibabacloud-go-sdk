// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiDimTableRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiDimTableRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiDimTableRecordResponseBody) *GetMultiDimTableRecordResponse
	GetBody() *GetMultiDimTableRecordResponseBody
}

type GetMultiDimTableRecordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiDimTableRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiDimTableRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableRecordResponse) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiDimTableRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiDimTableRecordResponse) GetBody() *GetMultiDimTableRecordResponseBody {
	return s.Body
}

func (s *GetMultiDimTableRecordResponse) SetHeaders(v map[string]*string) *GetMultiDimTableRecordResponse {
	s.Headers = v
	return s
}

func (s *GetMultiDimTableRecordResponse) SetStatusCode(v int32) *GetMultiDimTableRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiDimTableRecordResponse) SetBody(v *GetMultiDimTableRecordResponseBody) *GetMultiDimTableRecordResponse {
	s.Body = v
	return s
}

func (s *GetMultiDimTableRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
