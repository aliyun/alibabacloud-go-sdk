// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageGroupBandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLiveMessageGroupBandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLiveMessageGroupBandResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLiveMessageGroupBandResponseBody) *ModifyLiveMessageGroupBandResponse
	GetBody() *ModifyLiveMessageGroupBandResponseBody
}

type ModifyLiveMessageGroupBandResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLiveMessageGroupBandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLiveMessageGroupBandResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageGroupBandResponse) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageGroupBandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLiveMessageGroupBandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLiveMessageGroupBandResponse) GetBody() *ModifyLiveMessageGroupBandResponseBody {
	return s.Body
}

func (s *ModifyLiveMessageGroupBandResponse) SetHeaders(v map[string]*string) *ModifyLiveMessageGroupBandResponse {
	s.Headers = v
	return s
}

func (s *ModifyLiveMessageGroupBandResponse) SetStatusCode(v int32) *ModifyLiveMessageGroupBandResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLiveMessageGroupBandResponse) SetBody(v *ModifyLiveMessageGroupBandResponseBody) *ModifyLiveMessageGroupBandResponse {
	s.Body = v
	return s
}

func (s *ModifyLiveMessageGroupBandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
