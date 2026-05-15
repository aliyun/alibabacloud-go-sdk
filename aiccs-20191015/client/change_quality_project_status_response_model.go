// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeQualityProjectStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeQualityProjectStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeQualityProjectStatusResponse
	GetStatusCode() *int32
	SetBody(v *ChangeQualityProjectStatusResponseBody) *ChangeQualityProjectStatusResponse
	GetBody() *ChangeQualityProjectStatusResponseBody
}

type ChangeQualityProjectStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeQualityProjectStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeQualityProjectStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeQualityProjectStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeQualityProjectStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeQualityProjectStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeQualityProjectStatusResponse) GetBody() *ChangeQualityProjectStatusResponseBody {
	return s.Body
}

func (s *ChangeQualityProjectStatusResponse) SetHeaders(v map[string]*string) *ChangeQualityProjectStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeQualityProjectStatusResponse) SetStatusCode(v int32) *ChangeQualityProjectStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeQualityProjectStatusResponse) SetBody(v *ChangeQualityProjectStatusResponseBody) *ChangeQualityProjectStatusResponse {
	s.Body = v
	return s
}

func (s *ChangeQualityProjectStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
