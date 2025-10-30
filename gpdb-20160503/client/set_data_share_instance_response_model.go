// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataShareInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDataShareInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDataShareInstanceResponse
	GetStatusCode() *int32
	SetBody(v *SetDataShareInstanceResponseBody) *SetDataShareInstanceResponse
	GetBody() *SetDataShareInstanceResponseBody
}

type SetDataShareInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDataShareInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDataShareInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDataShareInstanceResponse) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDataShareInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDataShareInstanceResponse) GetBody() *SetDataShareInstanceResponseBody {
	return s.Body
}

func (s *SetDataShareInstanceResponse) SetHeaders(v map[string]*string) *SetDataShareInstanceResponse {
	s.Headers = v
	return s
}

func (s *SetDataShareInstanceResponse) SetStatusCode(v int32) *SetDataShareInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataShareInstanceResponse) SetBody(v *SetDataShareInstanceResponseBody) *SetDataShareInstanceResponse {
	s.Body = v
	return s
}

func (s *SetDataShareInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
