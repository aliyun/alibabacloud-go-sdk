// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubSceneSeqResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSubSceneSeqResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSubSceneSeqResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSubSceneSeqResponseBody) *UpdateSubSceneSeqResponse
	GetBody() *UpdateSubSceneSeqResponseBody
}

type UpdateSubSceneSeqResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSubSceneSeqResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSubSceneSeqResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubSceneSeqResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneSeqResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSubSceneSeqResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSubSceneSeqResponse) GetBody() *UpdateSubSceneSeqResponseBody {
	return s.Body
}

func (s *UpdateSubSceneSeqResponse) SetHeaders(v map[string]*string) *UpdateSubSceneSeqResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubSceneSeqResponse) SetStatusCode(v int32) *UpdateSubSceneSeqResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubSceneSeqResponse) SetBody(v *UpdateSubSceneSeqResponseBody) *UpdateSubSceneSeqResponse {
	s.Body = v
	return s
}

func (s *UpdateSubSceneSeqResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
