// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMetaLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMetaLogResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMetaLogResponseBody) *UpdateMetaLogResponse
	GetBody() *UpdateMetaLogResponseBody
}

type UpdateMetaLogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMetaLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMetaLogResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaLogResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetaLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMetaLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMetaLogResponse) GetBody() *UpdateMetaLogResponseBody {
	return s.Body
}

func (s *UpdateMetaLogResponse) SetHeaders(v map[string]*string) *UpdateMetaLogResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetaLogResponse) SetStatusCode(v int32) *UpdateMetaLogResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMetaLogResponse) SetBody(v *UpdateMetaLogResponseBody) *UpdateMetaLogResponse {
	s.Body = v
	return s
}

func (s *UpdateMetaLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
