// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfflineTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOfflineTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOfflineTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetOfflineTaskResponseBody) *GetOfflineTaskResponse
	GetBody() *GetOfflineTaskResponseBody
}

type GetOfflineTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOfflineTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOfflineTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskResponse) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOfflineTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOfflineTaskResponse) GetBody() *GetOfflineTaskResponseBody {
	return s.Body
}

func (s *GetOfflineTaskResponse) SetHeaders(v map[string]*string) *GetOfflineTaskResponse {
	s.Headers = v
	return s
}

func (s *GetOfflineTaskResponse) SetStatusCode(v int32) *GetOfflineTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOfflineTaskResponse) SetBody(v *GetOfflineTaskResponseBody) *GetOfflineTaskResponse {
	s.Body = v
	return s
}

func (s *GetOfflineTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
