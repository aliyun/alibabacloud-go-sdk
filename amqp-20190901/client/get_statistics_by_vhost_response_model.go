// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStatisticsByVhostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStatisticsByVhostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStatisticsByVhostResponse
	GetStatusCode() *int32
	SetBody(v *GetStatisticsByVhostResponseBody) *GetStatisticsByVhostResponse
	GetBody() *GetStatisticsByVhostResponseBody
}

type GetStatisticsByVhostResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStatisticsByVhostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStatisticsByVhostResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStatisticsByVhostResponse) GoString() string {
	return s.String()
}

func (s *GetStatisticsByVhostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStatisticsByVhostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStatisticsByVhostResponse) GetBody() *GetStatisticsByVhostResponseBody {
	return s.Body
}

func (s *GetStatisticsByVhostResponse) SetHeaders(v map[string]*string) *GetStatisticsByVhostResponse {
	s.Headers = v
	return s
}

func (s *GetStatisticsByVhostResponse) SetStatusCode(v int32) *GetStatisticsByVhostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStatisticsByVhostResponse) SetBody(v *GetStatisticsByVhostResponseBody) *GetStatisticsByVhostResponse {
	s.Body = v
	return s
}

func (s *GetStatisticsByVhostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
