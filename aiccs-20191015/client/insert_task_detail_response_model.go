// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *InsertTaskDetailResponseBody) *InsertTaskDetailResponse
	GetBody() *InsertTaskDetailResponseBody
}

type InsertTaskDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *InsertTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertTaskDetailResponse) GetBody() *InsertTaskDetailResponseBody {
	return s.Body
}

func (s *InsertTaskDetailResponse) SetHeaders(v map[string]*string) *InsertTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *InsertTaskDetailResponse) SetStatusCode(v int32) *InsertTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertTaskDetailResponse) SetBody(v *InsertTaskDetailResponseBody) *InsertTaskDetailResponse {
	s.Body = v
	return s
}

func (s *InsertTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
