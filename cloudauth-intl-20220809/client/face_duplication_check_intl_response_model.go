// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceDuplicationCheckIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FaceDuplicationCheckIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FaceDuplicationCheckIntlResponse
	GetStatusCode() *int32
	SetBody(v *FaceDuplicationCheckIntlResponseBody) *FaceDuplicationCheckIntlResponse
	GetBody() *FaceDuplicationCheckIntlResponseBody
}

type FaceDuplicationCheckIntlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceDuplicationCheckIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceDuplicationCheckIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s FaceDuplicationCheckIntlResponse) GoString() string {
	return s.String()
}

func (s *FaceDuplicationCheckIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FaceDuplicationCheckIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FaceDuplicationCheckIntlResponse) GetBody() *FaceDuplicationCheckIntlResponseBody {
	return s.Body
}

func (s *FaceDuplicationCheckIntlResponse) SetHeaders(v map[string]*string) *FaceDuplicationCheckIntlResponse {
	s.Headers = v
	return s
}

func (s *FaceDuplicationCheckIntlResponse) SetStatusCode(v int32) *FaceDuplicationCheckIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponse) SetBody(v *FaceDuplicationCheckIntlResponseBody) *FaceDuplicationCheckIntlResponse {
	s.Body = v
	return s
}

func (s *FaceDuplicationCheckIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
