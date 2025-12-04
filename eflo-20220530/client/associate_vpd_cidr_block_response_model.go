// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateVpdCidrBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateVpdCidrBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateVpdCidrBlockResponse
	GetStatusCode() *int32
	SetBody(v *AssociateVpdCidrBlockResponseBody) *AssociateVpdCidrBlockResponse
	GetBody() *AssociateVpdCidrBlockResponseBody
}

type AssociateVpdCidrBlockResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateVpdCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateVpdCidrBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateVpdCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *AssociateVpdCidrBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateVpdCidrBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateVpdCidrBlockResponse) GetBody() *AssociateVpdCidrBlockResponseBody {
	return s.Body
}

func (s *AssociateVpdCidrBlockResponse) SetHeaders(v map[string]*string) *AssociateVpdCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *AssociateVpdCidrBlockResponse) SetStatusCode(v int32) *AssociateVpdCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateVpdCidrBlockResponse) SetBody(v *AssociateVpdCidrBlockResponseBody) *AssociateVpdCidrBlockResponse {
	s.Body = v
	return s
}

func (s *AssociateVpdCidrBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
