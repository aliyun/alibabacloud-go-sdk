// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeVideoFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MergeVideoFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MergeVideoFaceResponse
	GetStatusCode() *int32
	SetBody(v *MergeVideoFaceResponseBody) *MergeVideoFaceResponse
	GetBody() *MergeVideoFaceResponseBody
}

type MergeVideoFaceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MergeVideoFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MergeVideoFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoFaceResponse) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MergeVideoFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MergeVideoFaceResponse) GetBody() *MergeVideoFaceResponseBody {
	return s.Body
}

func (s *MergeVideoFaceResponse) SetHeaders(v map[string]*string) *MergeVideoFaceResponse {
	s.Headers = v
	return s
}

func (s *MergeVideoFaceResponse) SetStatusCode(v int32) *MergeVideoFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeVideoFaceResponse) SetBody(v *MergeVideoFaceResponseBody) *MergeVideoFaceResponse {
	s.Body = v
	return s
}

func (s *MergeVideoFaceResponse) Validate() error {
	return dara.Validate(s)
}
