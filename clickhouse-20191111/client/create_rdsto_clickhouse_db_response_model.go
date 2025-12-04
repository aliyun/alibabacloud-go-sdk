// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRDSToClickhouseDbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRDSToClickhouseDbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRDSToClickhouseDbResponse
	GetStatusCode() *int32
	SetBody(v *CreateRDSToClickhouseDbResponseBody) *CreateRDSToClickhouseDbResponse
	GetBody() *CreateRDSToClickhouseDbResponseBody
}

type CreateRDSToClickhouseDbResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRDSToClickhouseDbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRDSToClickhouseDbResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRDSToClickhouseDbResponse) GoString() string {
	return s.String()
}

func (s *CreateRDSToClickhouseDbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRDSToClickhouseDbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRDSToClickhouseDbResponse) GetBody() *CreateRDSToClickhouseDbResponseBody {
	return s.Body
}

func (s *CreateRDSToClickhouseDbResponse) SetHeaders(v map[string]*string) *CreateRDSToClickhouseDbResponse {
	s.Headers = v
	return s
}

func (s *CreateRDSToClickhouseDbResponse) SetStatusCode(v int32) *CreateRDSToClickhouseDbResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRDSToClickhouseDbResponse) SetBody(v *CreateRDSToClickhouseDbResponseBody) *CreateRDSToClickhouseDbResponse {
	s.Body = v
	return s
}

func (s *CreateRDSToClickhouseDbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
