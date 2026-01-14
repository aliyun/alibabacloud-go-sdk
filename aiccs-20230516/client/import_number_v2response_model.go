// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNumberV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportNumberV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportNumberV2Response
	GetStatusCode() *int32
	SetBody(v *ImportNumberV2ResponseBody) *ImportNumberV2Response
	GetBody() *ImportNumberV2ResponseBody
}

type ImportNumberV2Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportNumberV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportNumberV2Response) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberV2Response) GoString() string {
	return s.String()
}

func (s *ImportNumberV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportNumberV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportNumberV2Response) GetBody() *ImportNumberV2ResponseBody {
	return s.Body
}

func (s *ImportNumberV2Response) SetHeaders(v map[string]*string) *ImportNumberV2Response {
	s.Headers = v
	return s
}

func (s *ImportNumberV2Response) SetStatusCode(v int32) *ImportNumberV2Response {
	s.StatusCode = &v
	return s
}

func (s *ImportNumberV2Response) SetBody(v *ImportNumberV2ResponseBody) *ImportNumberV2Response {
	s.Body = v
	return s
}

func (s *ImportNumberV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
