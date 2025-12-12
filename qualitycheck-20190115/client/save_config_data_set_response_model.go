// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveConfigDataSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveConfigDataSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveConfigDataSetResponse
	GetStatusCode() *int32
	SetBody(v *SaveConfigDataSetResponseBody) *SaveConfigDataSetResponse
	GetBody() *SaveConfigDataSetResponseBody
}

type SaveConfigDataSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveConfigDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveConfigDataSetResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveConfigDataSetResponse) GoString() string {
	return s.String()
}

func (s *SaveConfigDataSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveConfigDataSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveConfigDataSetResponse) GetBody() *SaveConfigDataSetResponseBody {
	return s.Body
}

func (s *SaveConfigDataSetResponse) SetHeaders(v map[string]*string) *SaveConfigDataSetResponse {
	s.Headers = v
	return s
}

func (s *SaveConfigDataSetResponse) SetStatusCode(v int32) *SaveConfigDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveConfigDataSetResponse) SetBody(v *SaveConfigDataSetResponseBody) *SaveConfigDataSetResponse {
	s.Body = v
	return s
}

func (s *SaveConfigDataSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
