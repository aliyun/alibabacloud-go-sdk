// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataArchiveCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataArchiveCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataArchiveCountResponse
	GetStatusCode() *int32
	SetBody(v *GetDataArchiveCountResponseBody) *GetDataArchiveCountResponse
	GetBody() *GetDataArchiveCountResponseBody
}

type GetDataArchiveCountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataArchiveCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataArchiveCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveCountResponse) GoString() string {
	return s.String()
}

func (s *GetDataArchiveCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataArchiveCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataArchiveCountResponse) GetBody() *GetDataArchiveCountResponseBody {
	return s.Body
}

func (s *GetDataArchiveCountResponse) SetHeaders(v map[string]*string) *GetDataArchiveCountResponse {
	s.Headers = v
	return s
}

func (s *GetDataArchiveCountResponse) SetStatusCode(v int32) *GetDataArchiveCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataArchiveCountResponse) SetBody(v *GetDataArchiveCountResponseBody) *GetDataArchiveCountResponse {
	s.Body = v
	return s
}

func (s *GetDataArchiveCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
