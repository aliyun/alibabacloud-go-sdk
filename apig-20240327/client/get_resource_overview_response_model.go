// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceOverviewResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceOverviewResponseBody) *GetResourceOverviewResponse
	GetBody() *GetResourceOverviewResponseBody
}

type GetResourceOverviewResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceOverviewResponse) GetBody() *GetResourceOverviewResponseBody {
	return s.Body
}

func (s *GetResourceOverviewResponse) SetHeaders(v map[string]*string) *GetResourceOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetResourceOverviewResponse) SetStatusCode(v int32) *GetResourceOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceOverviewResponse) SetBody(v *GetResourceOverviewResponseBody) *GetResourceOverviewResponse {
	s.Body = v
	return s
}

func (s *GetResourceOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
