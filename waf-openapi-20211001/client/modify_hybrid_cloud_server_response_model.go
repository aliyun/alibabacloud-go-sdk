// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridCloudServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridCloudServerResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridCloudServerResponseBody) *ModifyHybridCloudServerResponse
	GetBody() *ModifyHybridCloudServerResponseBody
}

type ModifyHybridCloudServerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridCloudServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridCloudServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudServerResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridCloudServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridCloudServerResponse) GetBody() *ModifyHybridCloudServerResponseBody {
	return s.Body
}

func (s *ModifyHybridCloudServerResponse) SetHeaders(v map[string]*string) *ModifyHybridCloudServerResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridCloudServerResponse) SetStatusCode(v int32) *ModifyHybridCloudServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridCloudServerResponse) SetBody(v *ModifyHybridCloudServerResponseBody) *ModifyHybridCloudServerResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridCloudServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
