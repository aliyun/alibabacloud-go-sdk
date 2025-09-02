// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgSceneAddOrUpdateSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgSceneAddOrUpdateSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgSceneAddOrUpdateSceneResponse
	GetStatusCode() *int32
	SetBody(v *DsgSceneAddOrUpdateSceneResponseBody) *DsgSceneAddOrUpdateSceneResponse
	GetBody() *DsgSceneAddOrUpdateSceneResponseBody
}

type DsgSceneAddOrUpdateSceneResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgSceneAddOrUpdateSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgSceneAddOrUpdateSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgSceneAddOrUpdateSceneResponse) GoString() string {
	return s.String()
}

func (s *DsgSceneAddOrUpdateSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgSceneAddOrUpdateSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgSceneAddOrUpdateSceneResponse) GetBody() *DsgSceneAddOrUpdateSceneResponseBody {
	return s.Body
}

func (s *DsgSceneAddOrUpdateSceneResponse) SetHeaders(v map[string]*string) *DsgSceneAddOrUpdateSceneResponse {
	s.Headers = v
	return s
}

func (s *DsgSceneAddOrUpdateSceneResponse) SetStatusCode(v int32) *DsgSceneAddOrUpdateSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneResponse) SetBody(v *DsgSceneAddOrUpdateSceneResponseBody) *DsgSceneAddOrUpdateSceneResponse {
	s.Body = v
	return s
}

func (s *DsgSceneAddOrUpdateSceneResponse) Validate() error {
	return dara.Validate(s)
}
