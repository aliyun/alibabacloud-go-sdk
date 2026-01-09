// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectDataSet2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SelectDataSet2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SelectDataSet2Response
	GetStatusCode() *int32
	SetBody(v *SelectDataSet2ResponseBody) *SelectDataSet2Response
	GetBody() *SelectDataSet2ResponseBody
}

type SelectDataSet2Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectDataSet2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectDataSet2Response) String() string {
	return dara.Prettify(s)
}

func (s SelectDataSet2Response) GoString() string {
	return s.String()
}

func (s *SelectDataSet2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SelectDataSet2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SelectDataSet2Response) GetBody() *SelectDataSet2ResponseBody {
	return s.Body
}

func (s *SelectDataSet2Response) SetHeaders(v map[string]*string) *SelectDataSet2Response {
	s.Headers = v
	return s
}

func (s *SelectDataSet2Response) SetStatusCode(v int32) *SelectDataSet2Response {
	s.StatusCode = &v
	return s
}

func (s *SelectDataSet2Response) SetBody(v *SelectDataSet2ResponseBody) *SelectDataSet2Response {
	s.Body = v
	return s
}

func (s *SelectDataSet2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
