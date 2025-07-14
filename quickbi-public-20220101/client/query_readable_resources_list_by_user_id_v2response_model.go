// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReadableResourcesListByUserIdV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryReadableResourcesListByUserIdV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryReadableResourcesListByUserIdV2Response
	GetStatusCode() *int32
	SetBody(v *QueryReadableResourcesListByUserIdV2ResponseBody) *QueryReadableResourcesListByUserIdV2Response
	GetBody() *QueryReadableResourcesListByUserIdV2ResponseBody
}

type QueryReadableResourcesListByUserIdV2Response struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReadableResourcesListByUserIdV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReadableResourcesListByUserIdV2Response) String() string {
	return dara.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdV2Response) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryReadableResourcesListByUserIdV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryReadableResourcesListByUserIdV2Response) GetBody() *QueryReadableResourcesListByUserIdV2ResponseBody {
	return s.Body
}

func (s *QueryReadableResourcesListByUserIdV2Response) SetHeaders(v map[string]*string) *QueryReadableResourcesListByUserIdV2Response {
	s.Headers = v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2Response) SetStatusCode(v int32) *QueryReadableResourcesListByUserIdV2Response {
	s.StatusCode = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2Response) SetBody(v *QueryReadableResourcesListByUserIdV2ResponseBody) *QueryReadableResourcesListByUserIdV2Response {
	s.Body = v
	return s
}

func (s *QueryReadableResourcesListByUserIdV2Response) Validate() error {
	return dara.Validate(s)
}
