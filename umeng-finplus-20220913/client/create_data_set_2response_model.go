// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSet2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataSet2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataSet2Response
	GetStatusCode() *int32
	SetBody(v *CreateDataSet2ResponseBody) *CreateDataSet2Response
	GetBody() *CreateDataSet2ResponseBody
}

type CreateDataSet2Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataSet2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataSet2Response) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSet2Response) GoString() string {
	return s.String()
}

func (s *CreateDataSet2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataSet2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataSet2Response) GetBody() *CreateDataSet2ResponseBody {
	return s.Body
}

func (s *CreateDataSet2Response) SetHeaders(v map[string]*string) *CreateDataSet2Response {
	s.Headers = v
	return s
}

func (s *CreateDataSet2Response) SetStatusCode(v int32) *CreateDataSet2Response {
	s.StatusCode = &v
	return s
}

func (s *CreateDataSet2Response) SetBody(v *CreateDataSet2ResponseBody) *CreateDataSet2Response {
	s.Body = v
	return s
}

func (s *CreateDataSet2Response) Validate() error {
	return dara.Validate(s)
}
