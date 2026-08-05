// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfflineTaskLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOfflineTaskLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOfflineTaskLogResponse
	GetStatusCode() *int32
	SetBody(v *GetOfflineTaskLogResponseBody) *GetOfflineTaskLogResponse
	GetBody() *GetOfflineTaskLogResponseBody
}

type GetOfflineTaskLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOfflineTaskLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOfflineTaskLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskLogResponse) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOfflineTaskLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOfflineTaskLogResponse) GetBody() *GetOfflineTaskLogResponseBody {
	return s.Body
}

func (s *GetOfflineTaskLogResponse) SetHeaders(v map[string]*string) *GetOfflineTaskLogResponse {
	s.Headers = v
	return s
}

func (s *GetOfflineTaskLogResponse) SetStatusCode(v int32) *GetOfflineTaskLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOfflineTaskLogResponse) SetBody(v *GetOfflineTaskLogResponseBody) *GetOfflineTaskLogResponse {
	s.Body = v
	return s
}

func (s *GetOfflineTaskLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
