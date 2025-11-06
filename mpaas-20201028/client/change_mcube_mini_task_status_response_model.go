// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMcubeMiniTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeMcubeMiniTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeMcubeMiniTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *ChangeMcubeMiniTaskStatusResponseBody) *ChangeMcubeMiniTaskStatusResponse
	GetBody() *ChangeMcubeMiniTaskStatusResponseBody
}

type ChangeMcubeMiniTaskStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeMcubeMiniTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeMcubeMiniTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubeMiniTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeMcubeMiniTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeMcubeMiniTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeMcubeMiniTaskStatusResponse) GetBody() *ChangeMcubeMiniTaskStatusResponseBody {
	return s.Body
}

func (s *ChangeMcubeMiniTaskStatusResponse) SetHeaders(v map[string]*string) *ChangeMcubeMiniTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeMcubeMiniTaskStatusResponse) SetStatusCode(v int32) *ChangeMcubeMiniTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusResponse) SetBody(v *ChangeMcubeMiniTaskStatusResponseBody) *ChangeMcubeMiniTaskStatusResponse {
	s.Body = v
	return s
}

func (s *ChangeMcubeMiniTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
