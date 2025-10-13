// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUmodelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUmodelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUmodelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUmodelResponseBody) *UpdateUmodelResponse
	GetBody() *UpdateUmodelResponseBody
}

type UpdateUmodelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUmodelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUmodelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUmodelResponse) GoString() string {
	return s.String()
}

func (s *UpdateUmodelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUmodelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUmodelResponse) GetBody() *UpdateUmodelResponseBody {
	return s.Body
}

func (s *UpdateUmodelResponse) SetHeaders(v map[string]*string) *UpdateUmodelResponse {
	s.Headers = v
	return s
}

func (s *UpdateUmodelResponse) SetStatusCode(v int32) *UpdateUmodelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUmodelResponse) SetBody(v *UpdateUmodelResponseBody) *UpdateUmodelResponse {
	s.Body = v
	return s
}

func (s *UpdateUmodelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
