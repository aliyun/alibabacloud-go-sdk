// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancesV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancesV2Response
	GetStatusCode() *int32
	SetBody(v *ListInstancesV2ResponseBody) *ListInstancesV2Response
	GetBody() *ListInstancesV2ResponseBody
}

type ListInstancesV2Response struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesV2Response) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesV2Response) GoString() string {
	return s.String()
}

func (s *ListInstancesV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancesV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancesV2Response) GetBody() *ListInstancesV2ResponseBody {
	return s.Body
}

func (s *ListInstancesV2Response) SetHeaders(v map[string]*string) *ListInstancesV2Response {
	s.Headers = v
	return s
}

func (s *ListInstancesV2Response) SetStatusCode(v int32) *ListInstancesV2Response {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesV2Response) SetBody(v *ListInstancesV2ResponseBody) *ListInstancesV2Response {
	s.Body = v
	return s
}

func (s *ListInstancesV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
