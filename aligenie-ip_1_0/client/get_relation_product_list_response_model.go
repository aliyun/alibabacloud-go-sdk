// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRelationProductListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRelationProductListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRelationProductListResponse
	GetStatusCode() *int32
	SetBody(v *GetRelationProductListResponseBody) *GetRelationProductListResponse
	GetBody() *GetRelationProductListResponseBody
}

type GetRelationProductListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRelationProductListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRelationProductListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRelationProductListResponse) GoString() string {
	return s.String()
}

func (s *GetRelationProductListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRelationProductListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRelationProductListResponse) GetBody() *GetRelationProductListResponseBody {
	return s.Body
}

func (s *GetRelationProductListResponse) SetHeaders(v map[string]*string) *GetRelationProductListResponse {
	s.Headers = v
	return s
}

func (s *GetRelationProductListResponse) SetStatusCode(v int32) *GetRelationProductListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRelationProductListResponse) SetBody(v *GetRelationProductListResponseBody) *GetRelationProductListResponse {
	s.Body = v
	return s
}

func (s *GetRelationProductListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
