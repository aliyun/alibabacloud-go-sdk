// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSet2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSet2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSet2Response
	GetStatusCode() *int32
	SetBody(v *ListDataSet2ResponseBody) *ListDataSet2Response
	GetBody() *ListDataSet2ResponseBody
}

type ListDataSet2Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSet2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSet2Response) String() string {
	return dara.Prettify(s)
}

func (s ListDataSet2Response) GoString() string {
	return s.String()
}

func (s *ListDataSet2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSet2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSet2Response) GetBody() *ListDataSet2ResponseBody {
	return s.Body
}

func (s *ListDataSet2Response) SetHeaders(v map[string]*string) *ListDataSet2Response {
	s.Headers = v
	return s
}

func (s *ListDataSet2Response) SetStatusCode(v int32) *ListDataSet2Response {
	s.StatusCode = &v
	return s
}

func (s *ListDataSet2Response) SetBody(v *ListDataSet2ResponseBody) *ListDataSet2Response {
	s.Body = v
	return s
}

func (s *ListDataSet2Response) Validate() error {
	return dara.Validate(s)
}
