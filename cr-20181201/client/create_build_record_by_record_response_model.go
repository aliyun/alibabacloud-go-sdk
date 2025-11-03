// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBuildRecordByRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBuildRecordByRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBuildRecordByRecordResponse
	GetStatusCode() *int32
	SetBody(v *CreateBuildRecordByRecordResponseBody) *CreateBuildRecordByRecordResponse
	GetBody() *CreateBuildRecordByRecordResponseBody
}

type CreateBuildRecordByRecordResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBuildRecordByRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBuildRecordByRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBuildRecordByRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBuildRecordByRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBuildRecordByRecordResponse) GetBody() *CreateBuildRecordByRecordResponseBody {
	return s.Body
}

func (s *CreateBuildRecordByRecordResponse) SetHeaders(v map[string]*string) *CreateBuildRecordByRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateBuildRecordByRecordResponse) SetStatusCode(v int32) *CreateBuildRecordByRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBuildRecordByRecordResponse) SetBody(v *CreateBuildRecordByRecordResponseBody) *CreateBuildRecordByRecordResponse {
	s.Body = v
	return s
}

func (s *CreateBuildRecordByRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
