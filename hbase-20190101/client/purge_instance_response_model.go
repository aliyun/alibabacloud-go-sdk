// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PurgeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PurgeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *PurgeInstanceResponseBody) *PurgeInstanceResponse
	GetBody() *PurgeInstanceResponseBody
}

type PurgeInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurgeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurgeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s PurgeInstanceResponse) GoString() string {
	return s.String()
}

func (s *PurgeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurgeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PurgeInstanceResponse) GetBody() *PurgeInstanceResponseBody {
	return s.Body
}

func (s *PurgeInstanceResponse) SetHeaders(v map[string]*string) *PurgeInstanceResponse {
	s.Headers = v
	return s
}

func (s *PurgeInstanceResponse) SetStatusCode(v int32) *PurgeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PurgeInstanceResponse) SetBody(v *PurgeInstanceResponseBody) *PurgeInstanceResponse {
	s.Body = v
	return s
}

func (s *PurgeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
