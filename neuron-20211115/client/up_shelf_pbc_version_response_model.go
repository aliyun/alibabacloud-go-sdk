// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpShelfPbcVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpShelfPbcVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpShelfPbcVersionResponse
	GetStatusCode() *int32
	SetBody(v *CatalogCommonResult) *UpShelfPbcVersionResponse
	GetBody() *CatalogCommonResult
}

type UpShelfPbcVersionResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogCommonResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpShelfPbcVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpShelfPbcVersionResponse) GoString() string {
	return s.String()
}

func (s *UpShelfPbcVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpShelfPbcVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpShelfPbcVersionResponse) GetBody() *CatalogCommonResult {
	return s.Body
}

func (s *UpShelfPbcVersionResponse) SetHeaders(v map[string]*string) *UpShelfPbcVersionResponse {
	s.Headers = v
	return s
}

func (s *UpShelfPbcVersionResponse) SetStatusCode(v int32) *UpShelfPbcVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpShelfPbcVersionResponse) SetBody(v *CatalogCommonResult) *UpShelfPbcVersionResponse {
	s.Body = v
	return s
}

func (s *UpShelfPbcVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
