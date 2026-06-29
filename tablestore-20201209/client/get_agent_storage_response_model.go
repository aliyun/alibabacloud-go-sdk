// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentStorageResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentStorageResponseBody) *GetAgentStorageResponse
	GetBody() *GetAgentStorageResponseBody
}

type GetAgentStorageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStorageResponse) GoString() string {
	return s.String()
}

func (s *GetAgentStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentStorageResponse) GetBody() *GetAgentStorageResponseBody {
	return s.Body
}

func (s *GetAgentStorageResponse) SetHeaders(v map[string]*string) *GetAgentStorageResponse {
	s.Headers = v
	return s
}

func (s *GetAgentStorageResponse) SetStatusCode(v int32) *GetAgentStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentStorageResponse) SetBody(v *GetAgentStorageResponseBody) *GetAgentStorageResponse {
	s.Body = v
	return s
}

func (s *GetAgentStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
