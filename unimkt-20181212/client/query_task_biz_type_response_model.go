// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskBizTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTaskBizTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTaskBizTypeResponse
	GetStatusCode() *int32
	SetBody(v *QueryTaskBizTypeResponseBody) *QueryTaskBizTypeResponse
	GetBody() *QueryTaskBizTypeResponseBody
}

type QueryTaskBizTypeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskBizTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskBizTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskBizTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTaskBizTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTaskBizTypeResponse) GetBody() *QueryTaskBizTypeResponseBody {
	return s.Body
}

func (s *QueryTaskBizTypeResponse) SetHeaders(v map[string]*string) *QueryTaskBizTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskBizTypeResponse) SetStatusCode(v int32) *QueryTaskBizTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskBizTypeResponse) SetBody(v *QueryTaskBizTypeResponseBody) *QueryTaskBizTypeResponse {
	s.Body = v
	return s
}

func (s *QueryTaskBizTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
