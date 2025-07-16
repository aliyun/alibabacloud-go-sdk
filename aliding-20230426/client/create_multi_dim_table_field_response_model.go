// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiDimTableFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMultiDimTableFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMultiDimTableFieldResponse
	GetStatusCode() *int32
	SetBody(v *CreateMultiDimTableFieldResponseBody) *CreateMultiDimTableFieldResponse
	GetBody() *CreateMultiDimTableFieldResponseBody
}

type CreateMultiDimTableFieldResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultiDimTableFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultiDimTableFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiDimTableFieldResponse) GoString() string {
	return s.String()
}

func (s *CreateMultiDimTableFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMultiDimTableFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMultiDimTableFieldResponse) GetBody() *CreateMultiDimTableFieldResponseBody {
	return s.Body
}

func (s *CreateMultiDimTableFieldResponse) SetHeaders(v map[string]*string) *CreateMultiDimTableFieldResponse {
	s.Headers = v
	return s
}

func (s *CreateMultiDimTableFieldResponse) SetStatusCode(v int32) *CreateMultiDimTableFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultiDimTableFieldResponse) SetBody(v *CreateMultiDimTableFieldResponseBody) *CreateMultiDimTableFieldResponse {
	s.Body = v
	return s
}

func (s *CreateMultiDimTableFieldResponse) Validate() error {
	return dara.Validate(s)
}
