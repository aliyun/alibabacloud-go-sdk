// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgScenedDeleteSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgScenedDeleteSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgScenedDeleteSceneResponse
	GetStatusCode() *int32
	SetBody(v *DsgScenedDeleteSceneResponseBody) *DsgScenedDeleteSceneResponse
	GetBody() *DsgScenedDeleteSceneResponseBody
}

type DsgScenedDeleteSceneResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgScenedDeleteSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgScenedDeleteSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgScenedDeleteSceneResponse) GoString() string {
	return s.String()
}

func (s *DsgScenedDeleteSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgScenedDeleteSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgScenedDeleteSceneResponse) GetBody() *DsgScenedDeleteSceneResponseBody {
	return s.Body
}

func (s *DsgScenedDeleteSceneResponse) SetHeaders(v map[string]*string) *DsgScenedDeleteSceneResponse {
	s.Headers = v
	return s
}

func (s *DsgScenedDeleteSceneResponse) SetStatusCode(v int32) *DsgScenedDeleteSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgScenedDeleteSceneResponse) SetBody(v *DsgScenedDeleteSceneResponseBody) *DsgScenedDeleteSceneResponse {
	s.Body = v
	return s
}

func (s *DsgScenedDeleteSceneResponse) Validate() error {
	return dara.Validate(s)
}
