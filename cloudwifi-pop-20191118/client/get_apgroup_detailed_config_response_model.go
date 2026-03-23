// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupDetailedConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApgroupDetailedConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApgroupDetailedConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetApgroupDetailedConfigResponseBody) *GetApgroupDetailedConfigResponse
	GetBody() *GetApgroupDetailedConfigResponseBody
}

type GetApgroupDetailedConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApgroupDetailedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApgroupDetailedConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupDetailedConfigResponse) GoString() string {
	return s.String()
}

func (s *GetApgroupDetailedConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApgroupDetailedConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApgroupDetailedConfigResponse) GetBody() *GetApgroupDetailedConfigResponseBody {
	return s.Body
}

func (s *GetApgroupDetailedConfigResponse) SetHeaders(v map[string]*string) *GetApgroupDetailedConfigResponse {
	s.Headers = v
	return s
}

func (s *GetApgroupDetailedConfigResponse) SetStatusCode(v int32) *GetApgroupDetailedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApgroupDetailedConfigResponse) SetBody(v *GetApgroupDetailedConfigResponseBody) *GetApgroupDetailedConfigResponse {
	s.Body = v
	return s
}

func (s *GetApgroupDetailedConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
