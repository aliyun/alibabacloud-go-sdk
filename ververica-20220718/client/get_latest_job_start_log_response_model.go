// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLatestJobStartLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLatestJobStartLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLatestJobStartLogResponse
	GetStatusCode() *int32
	SetBody(v *GetLatestJobStartLogResponseBody) *GetLatestJobStartLogResponse
	GetBody() *GetLatestJobStartLogResponseBody
}

type GetLatestJobStartLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLatestJobStartLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLatestJobStartLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLatestJobStartLogResponse) GoString() string {
	return s.String()
}

func (s *GetLatestJobStartLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLatestJobStartLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLatestJobStartLogResponse) GetBody() *GetLatestJobStartLogResponseBody {
	return s.Body
}

func (s *GetLatestJobStartLogResponse) SetHeaders(v map[string]*string) *GetLatestJobStartLogResponse {
	s.Headers = v
	return s
}

func (s *GetLatestJobStartLogResponse) SetStatusCode(v int32) *GetLatestJobStartLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLatestJobStartLogResponse) SetBody(v *GetLatestJobStartLogResponseBody) *GetLatestJobStartLogResponse {
	s.Body = v
	return s
}

func (s *GetLatestJobStartLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
