// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectNotifyReceiversResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProjectNotifyReceiversResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProjectNotifyReceiversResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProjectNotifyReceiversResponseBody) *DeleteProjectNotifyReceiversResponse
	GetBody() *DeleteProjectNotifyReceiversResponseBody
}

type DeleteProjectNotifyReceiversResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectNotifyReceiversResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectNotifyReceiversResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectNotifyReceiversResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectNotifyReceiversResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProjectNotifyReceiversResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProjectNotifyReceiversResponse) GetBody() *DeleteProjectNotifyReceiversResponseBody {
	return s.Body
}

func (s *DeleteProjectNotifyReceiversResponse) SetHeaders(v map[string]*string) *DeleteProjectNotifyReceiversResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectNotifyReceiversResponse) SetStatusCode(v int32) *DeleteProjectNotifyReceiversResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectNotifyReceiversResponse) SetBody(v *DeleteProjectNotifyReceiversResponseBody) *DeleteProjectNotifyReceiversResponse {
	s.Body = v
	return s
}

func (s *DeleteProjectNotifyReceiversResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
