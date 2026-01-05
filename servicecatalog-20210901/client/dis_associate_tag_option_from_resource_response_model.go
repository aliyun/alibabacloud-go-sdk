// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisAssociateTagOptionFromResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisAssociateTagOptionFromResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisAssociateTagOptionFromResourceResponse
	GetStatusCode() *int32
	SetBody(v *DisAssociateTagOptionFromResourceResponseBody) *DisAssociateTagOptionFromResourceResponse
	GetBody() *DisAssociateTagOptionFromResourceResponseBody
}

type DisAssociateTagOptionFromResourceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisAssociateTagOptionFromResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisAssociateTagOptionFromResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DisAssociateTagOptionFromResourceResponse) GoString() string {
	return s.String()
}

func (s *DisAssociateTagOptionFromResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisAssociateTagOptionFromResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisAssociateTagOptionFromResourceResponse) GetBody() *DisAssociateTagOptionFromResourceResponseBody {
	return s.Body
}

func (s *DisAssociateTagOptionFromResourceResponse) SetHeaders(v map[string]*string) *DisAssociateTagOptionFromResourceResponse {
	s.Headers = v
	return s
}

func (s *DisAssociateTagOptionFromResourceResponse) SetStatusCode(v int32) *DisAssociateTagOptionFromResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisAssociateTagOptionFromResourceResponse) SetBody(v *DisAssociateTagOptionFromResourceResponseBody) *DisAssociateTagOptionFromResourceResponse {
	s.Body = v
	return s
}

func (s *DisAssociateTagOptionFromResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
