// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveMessageGroupBandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveMessageGroupBandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveMessageGroupBandResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveMessageGroupBandResponseBody) *AddLiveMessageGroupBandResponse
	GetBody() *AddLiveMessageGroupBandResponseBody
}

type AddLiveMessageGroupBandResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveMessageGroupBandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveMessageGroupBandResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveMessageGroupBandResponse) GoString() string {
	return s.String()
}

func (s *AddLiveMessageGroupBandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveMessageGroupBandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveMessageGroupBandResponse) GetBody() *AddLiveMessageGroupBandResponseBody {
	return s.Body
}

func (s *AddLiveMessageGroupBandResponse) SetHeaders(v map[string]*string) *AddLiveMessageGroupBandResponse {
	s.Headers = v
	return s
}

func (s *AddLiveMessageGroupBandResponse) SetStatusCode(v int32) *AddLiveMessageGroupBandResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveMessageGroupBandResponse) SetBody(v *AddLiveMessageGroupBandResponseBody) *AddLiveMessageGroupBandResponse {
	s.Body = v
	return s
}

func (s *AddLiveMessageGroupBandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
