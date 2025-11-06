// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAckInfoByIntervalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAckInfoByIntervalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAckInfoByIntervalResponse
	GetStatusCode() *int32
	SetBody(v *GetAckInfoByIntervalResponseBody) *GetAckInfoByIntervalResponse
	GetBody() *GetAckInfoByIntervalResponseBody
}

type GetAckInfoByIntervalResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAckInfoByIntervalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAckInfoByIntervalResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAckInfoByIntervalResponse) GoString() string {
	return s.String()
}

func (s *GetAckInfoByIntervalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAckInfoByIntervalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAckInfoByIntervalResponse) GetBody() *GetAckInfoByIntervalResponseBody {
	return s.Body
}

func (s *GetAckInfoByIntervalResponse) SetHeaders(v map[string]*string) *GetAckInfoByIntervalResponse {
	s.Headers = v
	return s
}

func (s *GetAckInfoByIntervalResponse) SetStatusCode(v int32) *GetAckInfoByIntervalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAckInfoByIntervalResponse) SetBody(v *GetAckInfoByIntervalResponseBody) *GetAckInfoByIntervalResponse {
	s.Body = v
	return s
}

func (s *GetAckInfoByIntervalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
