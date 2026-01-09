// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataSet2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDataSet2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDataSet2Response
	GetStatusCode() *int32
	SetBody(v *RemoveDataSet2ResponseBody) *RemoveDataSet2Response
	GetBody() *RemoveDataSet2ResponseBody
}

type RemoveDataSet2Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDataSet2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDataSet2Response) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataSet2Response) GoString() string {
	return s.String()
}

func (s *RemoveDataSet2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDataSet2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDataSet2Response) GetBody() *RemoveDataSet2ResponseBody {
	return s.Body
}

func (s *RemoveDataSet2Response) SetHeaders(v map[string]*string) *RemoveDataSet2Response {
	s.Headers = v
	return s
}

func (s *RemoveDataSet2Response) SetStatusCode(v int32) *RemoveDataSet2Response {
	s.StatusCode = &v
	return s
}

func (s *RemoveDataSet2Response) SetBody(v *RemoveDataSet2ResponseBody) *RemoveDataSet2Response {
	s.Body = v
	return s
}

func (s *RemoveDataSet2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
