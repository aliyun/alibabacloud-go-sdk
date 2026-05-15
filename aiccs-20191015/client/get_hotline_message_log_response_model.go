// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineMessageLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotlineMessageLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotlineMessageLogResponse
	GetStatusCode() *int32
	SetBody(v *GetHotlineMessageLogResponseBody) *GetHotlineMessageLogResponse
	GetBody() *GetHotlineMessageLogResponseBody
}

type GetHotlineMessageLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotlineMessageLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotlineMessageLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineMessageLogResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineMessageLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotlineMessageLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotlineMessageLogResponse) GetBody() *GetHotlineMessageLogResponseBody {
	return s.Body
}

func (s *GetHotlineMessageLogResponse) SetHeaders(v map[string]*string) *GetHotlineMessageLogResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineMessageLogResponse) SetStatusCode(v int32) *GetHotlineMessageLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineMessageLogResponse) SetBody(v *GetHotlineMessageLogResponseBody) *GetHotlineMessageLogResponse {
	s.Body = v
	return s
}

func (s *GetHotlineMessageLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
