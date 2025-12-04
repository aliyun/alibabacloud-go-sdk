// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnAssociateVpdCidrBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnAssociateVpdCidrBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnAssociateVpdCidrBlockResponse
	GetStatusCode() *int32
	SetBody(v *UnAssociateVpdCidrBlockResponseBody) *UnAssociateVpdCidrBlockResponse
	GetBody() *UnAssociateVpdCidrBlockResponseBody
}

type UnAssociateVpdCidrBlockResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnAssociateVpdCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnAssociateVpdCidrBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s UnAssociateVpdCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *UnAssociateVpdCidrBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnAssociateVpdCidrBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnAssociateVpdCidrBlockResponse) GetBody() *UnAssociateVpdCidrBlockResponseBody {
	return s.Body
}

func (s *UnAssociateVpdCidrBlockResponse) SetHeaders(v map[string]*string) *UnAssociateVpdCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *UnAssociateVpdCidrBlockResponse) SetStatusCode(v int32) *UnAssociateVpdCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *UnAssociateVpdCidrBlockResponse) SetBody(v *UnAssociateVpdCidrBlockResponseBody) *UnAssociateVpdCidrBlockResponse {
	s.Body = v
	return s
}

func (s *UnAssociateVpdCidrBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
