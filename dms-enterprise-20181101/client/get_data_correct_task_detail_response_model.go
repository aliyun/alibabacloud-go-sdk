// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCorrectTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCorrectTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCorrectTaskDetailResponseBody) *GetDataCorrectTaskDetailResponse
	GetBody() *GetDataCorrectTaskDetailResponseBody
}

type GetDataCorrectTaskDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCorrectTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCorrectTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDataCorrectTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCorrectTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCorrectTaskDetailResponse) GetBody() *GetDataCorrectTaskDetailResponseBody {
	return s.Body
}

func (s *GetDataCorrectTaskDetailResponse) SetHeaders(v map[string]*string) *GetDataCorrectTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDataCorrectTaskDetailResponse) SetStatusCode(v int32) *GetDataCorrectTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponse) SetBody(v *GetDataCorrectTaskDetailResponseBody) *GetDataCorrectTaskDetailResponse {
	s.Body = v
	return s
}

func (s *GetDataCorrectTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
