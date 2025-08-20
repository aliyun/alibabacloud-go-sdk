// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsDatasoureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApsDatasoureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApsDatasoureResponse
	GetStatusCode() *int32
	SetBody(v *CreateApsDatasoureResponseBody) *CreateApsDatasoureResponse
	GetBody() *CreateApsDatasoureResponseBody
}

type CreateApsDatasoureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApsDatasoureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApsDatasoureResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApsDatasoureResponse) GoString() string {
	return s.String()
}

func (s *CreateApsDatasoureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApsDatasoureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApsDatasoureResponse) GetBody() *CreateApsDatasoureResponseBody {
	return s.Body
}

func (s *CreateApsDatasoureResponse) SetHeaders(v map[string]*string) *CreateApsDatasoureResponse {
	s.Headers = v
	return s
}

func (s *CreateApsDatasoureResponse) SetStatusCode(v int32) *CreateApsDatasoureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApsDatasoureResponse) SetBody(v *CreateApsDatasoureResponseBody) *CreateApsDatasoureResponse {
	s.Body = v
	return s
}

func (s *CreateApsDatasoureResponse) Validate() error {
	return dara.Validate(s)
}
