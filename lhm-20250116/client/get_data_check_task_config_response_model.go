// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTaskConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCheckTaskConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCheckTaskConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCheckTaskConfigResponseBody) *GetDataCheckTaskConfigResponse
	GetBody() *GetDataCheckTaskConfigResponseBody
}

type GetDataCheckTaskConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCheckTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCheckTaskConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDataCheckTaskConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCheckTaskConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCheckTaskConfigResponse) GetBody() *GetDataCheckTaskConfigResponseBody {
	return s.Body
}

func (s *GetDataCheckTaskConfigResponse) SetHeaders(v map[string]*string) *GetDataCheckTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDataCheckTaskConfigResponse) SetStatusCode(v int32) *GetDataCheckTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCheckTaskConfigResponse) SetBody(v *GetDataCheckTaskConfigResponseBody) *GetDataCheckTaskConfigResponse {
	s.Body = v
	return s
}

func (s *GetDataCheckTaskConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
