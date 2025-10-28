// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceMissingIndexListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceMissingIndexListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceMissingIndexListResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceMissingIndexListResponseBody) *GetInstanceMissingIndexListResponse
	GetBody() *GetInstanceMissingIndexListResponseBody
}

type GetInstanceMissingIndexListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceMissingIndexListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceMissingIndexListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMissingIndexListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceMissingIndexListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceMissingIndexListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceMissingIndexListResponse) GetBody() *GetInstanceMissingIndexListResponseBody {
	return s.Body
}

func (s *GetInstanceMissingIndexListResponse) SetHeaders(v map[string]*string) *GetInstanceMissingIndexListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceMissingIndexListResponse) SetStatusCode(v int32) *GetInstanceMissingIndexListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceMissingIndexListResponse) SetBody(v *GetInstanceMissingIndexListResponseBody) *GetInstanceMissingIndexListResponse {
	s.Body = v
	return s
}

func (s *GetInstanceMissingIndexListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
