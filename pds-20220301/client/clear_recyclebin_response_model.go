// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearRecyclebinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearRecyclebinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearRecyclebinResponse
	GetStatusCode() *int32
	SetBody(v *ClearRecyclebinResponseBody) *ClearRecyclebinResponse
	GetBody() *ClearRecyclebinResponseBody
}

type ClearRecyclebinResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearRecyclebinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearRecyclebinResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearRecyclebinResponse) GoString() string {
	return s.String()
}

func (s *ClearRecyclebinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearRecyclebinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearRecyclebinResponse) GetBody() *ClearRecyclebinResponseBody {
	return s.Body
}

func (s *ClearRecyclebinResponse) SetHeaders(v map[string]*string) *ClearRecyclebinResponse {
	s.Headers = v
	return s
}

func (s *ClearRecyclebinResponse) SetStatusCode(v int32) *ClearRecyclebinResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearRecyclebinResponse) SetBody(v *ClearRecyclebinResponseBody) *ClearRecyclebinResponse {
	s.Body = v
	return s
}

func (s *ClearRecyclebinResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
