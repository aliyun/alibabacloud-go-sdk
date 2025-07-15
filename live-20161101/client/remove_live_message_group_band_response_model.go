// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveLiveMessageGroupBandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveLiveMessageGroupBandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveLiveMessageGroupBandResponse
	GetStatusCode() *int32
	SetBody(v *RemoveLiveMessageGroupBandResponseBody) *RemoveLiveMessageGroupBandResponse
	GetBody() *RemoveLiveMessageGroupBandResponseBody
}

type RemoveLiveMessageGroupBandResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveLiveMessageGroupBandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveLiveMessageGroupBandResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveLiveMessageGroupBandResponse) GoString() string {
	return s.String()
}

func (s *RemoveLiveMessageGroupBandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveLiveMessageGroupBandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveLiveMessageGroupBandResponse) GetBody() *RemoveLiveMessageGroupBandResponseBody {
	return s.Body
}

func (s *RemoveLiveMessageGroupBandResponse) SetHeaders(v map[string]*string) *RemoveLiveMessageGroupBandResponse {
	s.Headers = v
	return s
}

func (s *RemoveLiveMessageGroupBandResponse) SetStatusCode(v int32) *RemoveLiveMessageGroupBandResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveLiveMessageGroupBandResponse) SetBody(v *RemoveLiveMessageGroupBandResponseBody) *RemoveLiveMessageGroupBandResponse {
	s.Body = v
	return s
}

func (s *RemoveLiveMessageGroupBandResponse) Validate() error {
	return dara.Validate(s)
}
