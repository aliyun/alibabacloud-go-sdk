// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDepGroupTreeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDepGroupTreeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDepGroupTreeDataResponse
	GetStatusCode() *int32
	SetBody(v *GetDepGroupTreeDataResponseBody) *GetDepGroupTreeDataResponse
	GetBody() *GetDepGroupTreeDataResponseBody
}

type GetDepGroupTreeDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDepGroupTreeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDepGroupTreeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDepGroupTreeDataResponse) GoString() string {
	return s.String()
}

func (s *GetDepGroupTreeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDepGroupTreeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDepGroupTreeDataResponse) GetBody() *GetDepGroupTreeDataResponseBody {
	return s.Body
}

func (s *GetDepGroupTreeDataResponse) SetHeaders(v map[string]*string) *GetDepGroupTreeDataResponse {
	s.Headers = v
	return s
}

func (s *GetDepGroupTreeDataResponse) SetStatusCode(v int32) *GetDepGroupTreeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDepGroupTreeDataResponse) SetBody(v *GetDepGroupTreeDataResponseBody) *GetDepGroupTreeDataResponse {
	s.Body = v
	return s
}

func (s *GetDepGroupTreeDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
