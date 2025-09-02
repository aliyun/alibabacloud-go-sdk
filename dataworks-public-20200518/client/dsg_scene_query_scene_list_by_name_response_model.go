// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgSceneQuerySceneListByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgSceneQuerySceneListByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgSceneQuerySceneListByNameResponse
	GetStatusCode() *int32
	SetBody(v *DsgSceneQuerySceneListByNameResponseBody) *DsgSceneQuerySceneListByNameResponse
	GetBody() *DsgSceneQuerySceneListByNameResponseBody
}

type DsgSceneQuerySceneListByNameResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgSceneQuerySceneListByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgSceneQuerySceneListByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgSceneQuerySceneListByNameResponse) GoString() string {
	return s.String()
}

func (s *DsgSceneQuerySceneListByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgSceneQuerySceneListByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgSceneQuerySceneListByNameResponse) GetBody() *DsgSceneQuerySceneListByNameResponseBody {
	return s.Body
}

func (s *DsgSceneQuerySceneListByNameResponse) SetHeaders(v map[string]*string) *DsgSceneQuerySceneListByNameResponse {
	s.Headers = v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponse) SetStatusCode(v int32) *DsgSceneQuerySceneListByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponse) SetBody(v *DsgSceneQuerySceneListByNameResponseBody) *DsgSceneQuerySceneListByNameResponse {
	s.Body = v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponse) Validate() error {
	return dara.Validate(s)
}
