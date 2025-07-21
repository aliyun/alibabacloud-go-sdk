// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowNodePrototypeV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlowNodePrototypeV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlowNodePrototypeV2Response
	GetStatusCode() *int32
	SetBody(v *ListFlowNodePrototypeV2ResponseBody) *ListFlowNodePrototypeV2Response
	GetBody() *ListFlowNodePrototypeV2ResponseBody
}

type ListFlowNodePrototypeV2Response struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowNodePrototypeV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlowNodePrototypeV2Response) String() string {
	return dara.Prettify(s)
}

func (s ListFlowNodePrototypeV2Response) GoString() string {
	return s.String()
}

func (s *ListFlowNodePrototypeV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlowNodePrototypeV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlowNodePrototypeV2Response) GetBody() *ListFlowNodePrototypeV2ResponseBody {
	return s.Body
}

func (s *ListFlowNodePrototypeV2Response) SetHeaders(v map[string]*string) *ListFlowNodePrototypeV2Response {
	s.Headers = v
	return s
}

func (s *ListFlowNodePrototypeV2Response) SetStatusCode(v int32) *ListFlowNodePrototypeV2Response {
	s.StatusCode = &v
	return s
}

func (s *ListFlowNodePrototypeV2Response) SetBody(v *ListFlowNodePrototypeV2ResponseBody) *ListFlowNodePrototypeV2Response {
	s.Body = v
	return s
}

func (s *ListFlowNodePrototypeV2Response) Validate() error {
	return dara.Validate(s)
}
