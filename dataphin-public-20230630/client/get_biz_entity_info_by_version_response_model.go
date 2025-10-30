// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizEntityInfoByVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBizEntityInfoByVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBizEntityInfoByVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetBizEntityInfoByVersionResponseBody) *GetBizEntityInfoByVersionResponse
	GetBody() *GetBizEntityInfoByVersionResponseBody
}

type GetBizEntityInfoByVersionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBizEntityInfoByVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBizEntityInfoByVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoByVersionResponse) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoByVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBizEntityInfoByVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBizEntityInfoByVersionResponse) GetBody() *GetBizEntityInfoByVersionResponseBody {
	return s.Body
}

func (s *GetBizEntityInfoByVersionResponse) SetHeaders(v map[string]*string) *GetBizEntityInfoByVersionResponse {
	s.Headers = v
	return s
}

func (s *GetBizEntityInfoByVersionResponse) SetStatusCode(v int32) *GetBizEntityInfoByVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponse) SetBody(v *GetBizEntityInfoByVersionResponseBody) *GetBizEntityInfoByVersionResponse {
	s.Body = v
	return s
}

func (s *GetBizEntityInfoByVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
