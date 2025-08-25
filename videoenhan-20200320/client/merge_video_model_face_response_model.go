// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeVideoModelFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MergeVideoModelFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MergeVideoModelFaceResponse
	GetStatusCode() *int32
	SetBody(v *MergeVideoModelFaceResponseBody) *MergeVideoModelFaceResponse
	GetBody() *MergeVideoModelFaceResponseBody
}

type MergeVideoModelFaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MergeVideoModelFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MergeVideoModelFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoModelFaceResponse) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MergeVideoModelFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MergeVideoModelFaceResponse) GetBody() *MergeVideoModelFaceResponseBody {
	return s.Body
}

func (s *MergeVideoModelFaceResponse) SetHeaders(v map[string]*string) *MergeVideoModelFaceResponse {
	s.Headers = v
	return s
}

func (s *MergeVideoModelFaceResponse) SetStatusCode(v int32) *MergeVideoModelFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeVideoModelFaceResponse) SetBody(v *MergeVideoModelFaceResponseBody) *MergeVideoModelFaceResponse {
	s.Body = v
	return s
}

func (s *MergeVideoModelFaceResponse) Validate() error {
	return dara.Validate(s)
}
