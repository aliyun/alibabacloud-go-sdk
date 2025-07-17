// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReDeployLhDagVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReDeployLhDagVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReDeployLhDagVersionResponse
	GetStatusCode() *int32
	SetBody(v *ReDeployLhDagVersionResponseBody) *ReDeployLhDagVersionResponse
	GetBody() *ReDeployLhDagVersionResponseBody
}

type ReDeployLhDagVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReDeployLhDagVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReDeployLhDagVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReDeployLhDagVersionResponse) GoString() string {
	return s.String()
}

func (s *ReDeployLhDagVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReDeployLhDagVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReDeployLhDagVersionResponse) GetBody() *ReDeployLhDagVersionResponseBody {
	return s.Body
}

func (s *ReDeployLhDagVersionResponse) SetHeaders(v map[string]*string) *ReDeployLhDagVersionResponse {
	s.Headers = v
	return s
}

func (s *ReDeployLhDagVersionResponse) SetStatusCode(v int32) *ReDeployLhDagVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReDeployLhDagVersionResponse) SetBody(v *ReDeployLhDagVersionResponseBody) *ReDeployLhDagVersionResponse {
	s.Body = v
	return s
}

func (s *ReDeployLhDagVersionResponse) Validate() error {
	return dara.Validate(s)
}
