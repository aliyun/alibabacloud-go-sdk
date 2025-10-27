// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRdDefaultSyncListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRdDefaultSyncListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRdDefaultSyncListResponse
	GetStatusCode() *int32
	SetBody(v *CreateRdDefaultSyncListResponseBody) *CreateRdDefaultSyncListResponse
	GetBody() *CreateRdDefaultSyncListResponseBody
}

type CreateRdDefaultSyncListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRdDefaultSyncListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRdDefaultSyncListResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRdDefaultSyncListResponse) GoString() string {
	return s.String()
}

func (s *CreateRdDefaultSyncListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRdDefaultSyncListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRdDefaultSyncListResponse) GetBody() *CreateRdDefaultSyncListResponseBody {
	return s.Body
}

func (s *CreateRdDefaultSyncListResponse) SetHeaders(v map[string]*string) *CreateRdDefaultSyncListResponse {
	s.Headers = v
	return s
}

func (s *CreateRdDefaultSyncListResponse) SetStatusCode(v int32) *CreateRdDefaultSyncListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRdDefaultSyncListResponse) SetBody(v *CreateRdDefaultSyncListResponseBody) *CreateRdDefaultSyncListResponse {
	s.Body = v
	return s
}

func (s *CreateRdDefaultSyncListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
