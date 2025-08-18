// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeContainerAppRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEdgeContainerAppRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEdgeContainerAppRecordResponse
	GetStatusCode() *int32
	SetBody(v *CreateEdgeContainerAppRecordResponseBody) *CreateEdgeContainerAppRecordResponse
	GetBody() *CreateEdgeContainerAppRecordResponseBody
}

type CreateEdgeContainerAppRecordResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEdgeContainerAppRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEdgeContainerAppRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEdgeContainerAppRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEdgeContainerAppRecordResponse) GetBody() *CreateEdgeContainerAppRecordResponseBody {
	return s.Body
}

func (s *CreateEdgeContainerAppRecordResponse) SetHeaders(v map[string]*string) *CreateEdgeContainerAppRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateEdgeContainerAppRecordResponse) SetStatusCode(v int32) *CreateEdgeContainerAppRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEdgeContainerAppRecordResponse) SetBody(v *CreateEdgeContainerAppRecordResponseBody) *CreateEdgeContainerAppRecordResponse {
	s.Body = v
	return s
}

func (s *CreateEdgeContainerAppRecordResponse) Validate() error {
	return dara.Validate(s)
}
