// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaClusterLabelListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpaClusterLabelListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpaClusterLabelListResponse
	GetStatusCode() *int32
	SetBody(v *GetOpaClusterLabelListResponseBody) *GetOpaClusterLabelListResponse
	GetBody() *GetOpaClusterLabelListResponseBody
}

type GetOpaClusterLabelListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpaClusterLabelListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpaClusterLabelListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterLabelListResponse) GoString() string {
	return s.String()
}

func (s *GetOpaClusterLabelListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpaClusterLabelListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpaClusterLabelListResponse) GetBody() *GetOpaClusterLabelListResponseBody {
	return s.Body
}

func (s *GetOpaClusterLabelListResponse) SetHeaders(v map[string]*string) *GetOpaClusterLabelListResponse {
	s.Headers = v
	return s
}

func (s *GetOpaClusterLabelListResponse) SetStatusCode(v int32) *GetOpaClusterLabelListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpaClusterLabelListResponse) SetBody(v *GetOpaClusterLabelListResponseBody) *GetOpaClusterLabelListResponse {
	s.Body = v
	return s
}

func (s *GetOpaClusterLabelListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
