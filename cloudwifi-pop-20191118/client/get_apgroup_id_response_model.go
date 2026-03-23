// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApgroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApgroupIdResponse
	GetStatusCode() *int32
	SetBody(v *GetApgroupIdResponseBody) *GetApgroupIdResponse
	GetBody() *GetApgroupIdResponseBody
}

type GetApgroupIdResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApgroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApgroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupIdResponse) GoString() string {
	return s.String()
}

func (s *GetApgroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApgroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApgroupIdResponse) GetBody() *GetApgroupIdResponseBody {
	return s.Body
}

func (s *GetApgroupIdResponse) SetHeaders(v map[string]*string) *GetApgroupIdResponse {
	s.Headers = v
	return s
}

func (s *GetApgroupIdResponse) SetStatusCode(v int32) *GetApgroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApgroupIdResponse) SetBody(v *GetApgroupIdResponseBody) *GetApgroupIdResponse {
	s.Body = v
	return s
}

func (s *GetApgroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
