// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTagOptionWithResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateTagOptionWithResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateTagOptionWithResourceResponse
	GetStatusCode() *int32
	SetBody(v *AssociateTagOptionWithResourceResponseBody) *AssociateTagOptionWithResourceResponse
	GetBody() *AssociateTagOptionWithResourceResponseBody
}

type AssociateTagOptionWithResourceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateTagOptionWithResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateTagOptionWithResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateTagOptionWithResourceResponse) GoString() string {
	return s.String()
}

func (s *AssociateTagOptionWithResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateTagOptionWithResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateTagOptionWithResourceResponse) GetBody() *AssociateTagOptionWithResourceResponseBody {
	return s.Body
}

func (s *AssociateTagOptionWithResourceResponse) SetHeaders(v map[string]*string) *AssociateTagOptionWithResourceResponse {
	s.Headers = v
	return s
}

func (s *AssociateTagOptionWithResourceResponse) SetStatusCode(v int32) *AssociateTagOptionWithResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateTagOptionWithResourceResponse) SetBody(v *AssociateTagOptionWithResourceResponseBody) *AssociateTagOptionWithResourceResponse {
	s.Body = v
	return s
}

func (s *AssociateTagOptionWithResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
