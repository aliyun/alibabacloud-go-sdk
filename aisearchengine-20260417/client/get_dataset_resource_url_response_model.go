// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetResourceUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatasetResourceUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatasetResourceUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetDatasetResourceUrlResponseBody) *GetDatasetResourceUrlResponse
	GetBody() *GetDatasetResourceUrlResponseBody
}

type GetDatasetResourceUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasetResourceUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasetResourceUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResourceUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetResourceUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatasetResourceUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatasetResourceUrlResponse) GetBody() *GetDatasetResourceUrlResponseBody {
	return s.Body
}

func (s *GetDatasetResourceUrlResponse) SetHeaders(v map[string]*string) *GetDatasetResourceUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetResourceUrlResponse) SetStatusCode(v int32) *GetDatasetResourceUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetResourceUrlResponse) SetBody(v *GetDatasetResourceUrlResponseBody) *GetDatasetResourceUrlResponse {
	s.Body = v
	return s
}

func (s *GetDatasetResourceUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
