// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineFlowVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnlineFlowVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnlineFlowVersionResponse
	GetStatusCode() *int32
	SetBody(v *OnlineFlowVersionResponseBody) *OnlineFlowVersionResponse
	GetBody() *OnlineFlowVersionResponseBody
}

type OnlineFlowVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineFlowVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineFlowVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s OnlineFlowVersionResponse) GoString() string {
	return s.String()
}

func (s *OnlineFlowVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnlineFlowVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnlineFlowVersionResponse) GetBody() *OnlineFlowVersionResponseBody {
	return s.Body
}

func (s *OnlineFlowVersionResponse) SetHeaders(v map[string]*string) *OnlineFlowVersionResponse {
	s.Headers = v
	return s
}

func (s *OnlineFlowVersionResponse) SetStatusCode(v int32) *OnlineFlowVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineFlowVersionResponse) SetBody(v *OnlineFlowVersionResponseBody) *OnlineFlowVersionResponse {
	s.Body = v
	return s
}

func (s *OnlineFlowVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
