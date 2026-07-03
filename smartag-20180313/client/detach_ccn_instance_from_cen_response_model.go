// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachCcnInstanceFromCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachCcnInstanceFromCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachCcnInstanceFromCenResponse
	GetStatusCode() *int32
	SetBody(v *DetachCcnInstanceFromCenResponseBody) *DetachCcnInstanceFromCenResponse
	GetBody() *DetachCcnInstanceFromCenResponseBody
}

type DetachCcnInstanceFromCenResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachCcnInstanceFromCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachCcnInstanceFromCenResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachCcnInstanceFromCenResponse) GoString() string {
	return s.String()
}

func (s *DetachCcnInstanceFromCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachCcnInstanceFromCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachCcnInstanceFromCenResponse) GetBody() *DetachCcnInstanceFromCenResponseBody {
	return s.Body
}

func (s *DetachCcnInstanceFromCenResponse) SetHeaders(v map[string]*string) *DetachCcnInstanceFromCenResponse {
	s.Headers = v
	return s
}

func (s *DetachCcnInstanceFromCenResponse) SetStatusCode(v int32) *DetachCcnInstanceFromCenResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachCcnInstanceFromCenResponse) SetBody(v *DetachCcnInstanceFromCenResponseBody) *DetachCcnInstanceFromCenResponse {
	s.Body = v
	return s
}

func (s *DetachCcnInstanceFromCenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
