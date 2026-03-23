// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupSsidConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApgroupSsidConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApgroupSsidConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetApgroupSsidConfigResponseBody) *GetApgroupSsidConfigResponse
	GetBody() *GetApgroupSsidConfigResponseBody
}

type GetApgroupSsidConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApgroupSsidConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApgroupSsidConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupSsidConfigResponse) GoString() string {
	return s.String()
}

func (s *GetApgroupSsidConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApgroupSsidConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApgroupSsidConfigResponse) GetBody() *GetApgroupSsidConfigResponseBody {
	return s.Body
}

func (s *GetApgroupSsidConfigResponse) SetHeaders(v map[string]*string) *GetApgroupSsidConfigResponse {
	s.Headers = v
	return s
}

func (s *GetApgroupSsidConfigResponse) SetStatusCode(v int32) *GetApgroupSsidConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApgroupSsidConfigResponse) SetBody(v *GetApgroupSsidConfigResponseBody) *GetApgroupSsidConfigResponse {
	s.Body = v
	return s
}

func (s *GetApgroupSsidConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
