// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataSetResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataSetResponseBody) *CreateDataSetResponse
	GetBody() *CreateDataSetResponseBody
}

type CreateDataSetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataSetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSetResponse) GoString() string {
	return s.String()
}

func (s *CreateDataSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataSetResponse) GetBody() *CreateDataSetResponseBody {
	return s.Body
}

func (s *CreateDataSetResponse) SetHeaders(v map[string]*string) *CreateDataSetResponse {
	s.Headers = v
	return s
}

func (s *CreateDataSetResponse) SetStatusCode(v int32) *CreateDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataSetResponse) SetBody(v *CreateDataSetResponseBody) *CreateDataSetResponse {
	s.Body = v
	return s
}

func (s *CreateDataSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
