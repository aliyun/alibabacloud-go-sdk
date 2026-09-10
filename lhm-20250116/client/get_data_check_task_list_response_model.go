// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCheckTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCheckTaskListResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCheckTaskListResponseBody) *GetDataCheckTaskListResponse
	GetBody() *GetDataCheckTaskListResponseBody
}

type GetDataCheckTaskListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCheckTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCheckTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTaskListResponse) GoString() string {
	return s.String()
}

func (s *GetDataCheckTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCheckTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCheckTaskListResponse) GetBody() *GetDataCheckTaskListResponseBody {
	return s.Body
}

func (s *GetDataCheckTaskListResponse) SetHeaders(v map[string]*string) *GetDataCheckTaskListResponse {
	s.Headers = v
	return s
}

func (s *GetDataCheckTaskListResponse) SetStatusCode(v int32) *GetDataCheckTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCheckTaskListResponse) SetBody(v *GetDataCheckTaskListResponseBody) *GetDataCheckTaskListResponse {
	s.Body = v
	return s
}

func (s *GetDataCheckTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
