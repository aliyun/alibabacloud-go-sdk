// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneLaboratoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneLaboratoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneLaboratoryResponse
	GetStatusCode() *int32
	SetBody(v *CloneLaboratoryResponseBody) *CloneLaboratoryResponse
	GetBody() *CloneLaboratoryResponseBody
}

type CloneLaboratoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneLaboratoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *CloneLaboratoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneLaboratoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneLaboratoryResponse) GetBody() *CloneLaboratoryResponseBody {
	return s.Body
}

func (s *CloneLaboratoryResponse) SetHeaders(v map[string]*string) *CloneLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *CloneLaboratoryResponse) SetStatusCode(v int32) *CloneLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneLaboratoryResponse) SetBody(v *CloneLaboratoryResponseBody) *CloneLaboratoryResponse {
	s.Body = v
	return s
}

func (s *CloneLaboratoryResponse) Validate() error {
	return dara.Validate(s)
}
