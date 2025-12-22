// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardForBase64Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MultiModalGuardForBase64Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MultiModalGuardForBase64Response
	GetStatusCode() *int32
	SetBody(v *MultiModalGuardForBase64ResponseBody) *MultiModalGuardForBase64Response
	GetBody() *MultiModalGuardForBase64ResponseBody
}

type MultiModalGuardForBase64Response struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultiModalGuardForBase64ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultiModalGuardForBase64Response) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardForBase64Response) GoString() string {
	return s.String()
}

func (s *MultiModalGuardForBase64Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MultiModalGuardForBase64Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MultiModalGuardForBase64Response) GetBody() *MultiModalGuardForBase64ResponseBody {
	return s.Body
}

func (s *MultiModalGuardForBase64Response) SetHeaders(v map[string]*string) *MultiModalGuardForBase64Response {
	s.Headers = v
	return s
}

func (s *MultiModalGuardForBase64Response) SetStatusCode(v int32) *MultiModalGuardForBase64Response {
	s.StatusCode = &v
	return s
}

func (s *MultiModalGuardForBase64Response) SetBody(v *MultiModalGuardForBase64ResponseBody) *MultiModalGuardForBase64Response {
	s.Body = v
	return s
}

func (s *MultiModalGuardForBase64Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
