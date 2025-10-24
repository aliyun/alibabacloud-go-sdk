// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudGroupShrinkServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridCloudGroupShrinkServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridCloudGroupShrinkServerResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridCloudGroupShrinkServerResponseBody) *ModifyHybridCloudGroupShrinkServerResponse
	GetBody() *ModifyHybridCloudGroupShrinkServerResponseBody
}

type ModifyHybridCloudGroupShrinkServerResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridCloudGroupShrinkServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridCloudGroupShrinkServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudGroupShrinkServerResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudGroupShrinkServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridCloudGroupShrinkServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridCloudGroupShrinkServerResponse) GetBody() *ModifyHybridCloudGroupShrinkServerResponseBody {
	return s.Body
}

func (s *ModifyHybridCloudGroupShrinkServerResponse) SetHeaders(v map[string]*string) *ModifyHybridCloudGroupShrinkServerResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridCloudGroupShrinkServerResponse) SetStatusCode(v int32) *ModifyHybridCloudGroupShrinkServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridCloudGroupShrinkServerResponse) SetBody(v *ModifyHybridCloudGroupShrinkServerResponseBody) *ModifyHybridCloudGroupShrinkServerResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridCloudGroupShrinkServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
