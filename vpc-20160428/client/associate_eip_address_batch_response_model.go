// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEipAddressBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateEipAddressBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateEipAddressBatchResponse
	GetStatusCode() *int32
	SetBody(v *AssociateEipAddressBatchResponseBody) *AssociateEipAddressBatchResponse
	GetBody() *AssociateEipAddressBatchResponseBody
}

type AssociateEipAddressBatchResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateEipAddressBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateEipAddressBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateEipAddressBatchResponse) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateEipAddressBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateEipAddressBatchResponse) GetBody() *AssociateEipAddressBatchResponseBody {
	return s.Body
}

func (s *AssociateEipAddressBatchResponse) SetHeaders(v map[string]*string) *AssociateEipAddressBatchResponse {
	s.Headers = v
	return s
}

func (s *AssociateEipAddressBatchResponse) SetStatusCode(v int32) *AssociateEipAddressBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateEipAddressBatchResponse) SetBody(v *AssociateEipAddressBatchResponseBody) *AssociateEipAddressBatchResponse {
	s.Body = v
	return s
}

func (s *AssociateEipAddressBatchResponse) Validate() error {
	return dara.Validate(s)
}
