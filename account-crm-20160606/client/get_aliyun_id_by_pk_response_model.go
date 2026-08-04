// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliyunIdByPkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAliyunIdByPkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAliyunIdByPkResponse
	GetStatusCode() *int32
	SetBody(v *GetAliyunIdByPkResponseBody) *GetAliyunIdByPkResponse
	GetBody() *GetAliyunIdByPkResponseBody
}

type GetAliyunIdByPkResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAliyunIdByPkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAliyunIdByPkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAliyunIdByPkResponse) GoString() string {
	return s.String()
}

func (s *GetAliyunIdByPkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAliyunIdByPkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAliyunIdByPkResponse) GetBody() *GetAliyunIdByPkResponseBody {
	return s.Body
}

func (s *GetAliyunIdByPkResponse) SetHeaders(v map[string]*string) *GetAliyunIdByPkResponse {
	s.Headers = v
	return s
}

func (s *GetAliyunIdByPkResponse) SetStatusCode(v int32) *GetAliyunIdByPkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAliyunIdByPkResponse) SetBody(v *GetAliyunIdByPkResponseBody) *GetAliyunIdByPkResponse {
	s.Body = v
	return s
}

func (s *GetAliyunIdByPkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
