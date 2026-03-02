// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePbcInvokeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePbcInvokeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePbcInvokeResponse
	GetStatusCode() *int32
	SetBody(v *CatalogCommonResult) *DeletePbcInvokeResponse
	GetBody() *CatalogCommonResult
}

type DeletePbcInvokeResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogCommonResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePbcInvokeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePbcInvokeResponse) GoString() string {
	return s.String()
}

func (s *DeletePbcInvokeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePbcInvokeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePbcInvokeResponse) GetBody() *CatalogCommonResult {
	return s.Body
}

func (s *DeletePbcInvokeResponse) SetHeaders(v map[string]*string) *DeletePbcInvokeResponse {
	s.Headers = v
	return s
}

func (s *DeletePbcInvokeResponse) SetStatusCode(v int32) *DeletePbcInvokeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePbcInvokeResponse) SetBody(v *CatalogCommonResult) *DeletePbcInvokeResponse {
	s.Body = v
	return s
}

func (s *DeletePbcInvokeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
