// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildStsAK2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BuildStsAK2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BuildStsAK2Response
	GetStatusCode() *int32
	SetBody(v *BuildStsAK2ResponseBody) *BuildStsAK2Response
	GetBody() *BuildStsAK2ResponseBody
}

type BuildStsAK2Response struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuildStsAK2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuildStsAK2Response) String() string {
	return dara.Prettify(s)
}

func (s BuildStsAK2Response) GoString() string {
	return s.String()
}

func (s *BuildStsAK2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BuildStsAK2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BuildStsAK2Response) GetBody() *BuildStsAK2ResponseBody {
	return s.Body
}

func (s *BuildStsAK2Response) SetHeaders(v map[string]*string) *BuildStsAK2Response {
	s.Headers = v
	return s
}

func (s *BuildStsAK2Response) SetStatusCode(v int32) *BuildStsAK2Response {
	s.StatusCode = &v
	return s
}

func (s *BuildStsAK2Response) SetBody(v *BuildStsAK2ResponseBody) *BuildStsAK2Response {
	s.Body = v
	return s
}

func (s *BuildStsAK2Response) Validate() error {
	return dara.Validate(s)
}
