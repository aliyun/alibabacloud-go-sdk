// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSelectionProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSelectionProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSelectionProductResponse
	GetStatusCode() *int32
	SetBody(v *Product) *GetSelectionProductResponse
	GetBody() *Product
}

type GetSelectionProductResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Product           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSelectionProductResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSelectionProductResponse) GoString() string {
	return s.String()
}

func (s *GetSelectionProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSelectionProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSelectionProductResponse) GetBody() *Product {
	return s.Body
}

func (s *GetSelectionProductResponse) SetHeaders(v map[string]*string) *GetSelectionProductResponse {
	s.Headers = v
	return s
}

func (s *GetSelectionProductResponse) SetStatusCode(v int32) *GetSelectionProductResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSelectionProductResponse) SetBody(v *Product) *GetSelectionProductResponse {
	s.Body = v
	return s
}

func (s *GetSelectionProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
