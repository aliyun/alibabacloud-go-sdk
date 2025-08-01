// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSendMessageToGlobeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSendMessageToGlobeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSendMessageToGlobeResponse
	GetStatusCode() *int32
	SetBody(v *BatchSendMessageToGlobeResponseBody) *BatchSendMessageToGlobeResponse
	GetBody() *BatchSendMessageToGlobeResponseBody
}

type BatchSendMessageToGlobeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSendMessageToGlobeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSendMessageToGlobeResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMessageToGlobeResponse) GoString() string {
	return s.String()
}

func (s *BatchSendMessageToGlobeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSendMessageToGlobeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSendMessageToGlobeResponse) GetBody() *BatchSendMessageToGlobeResponseBody {
	return s.Body
}

func (s *BatchSendMessageToGlobeResponse) SetHeaders(v map[string]*string) *BatchSendMessageToGlobeResponse {
	s.Headers = v
	return s
}

func (s *BatchSendMessageToGlobeResponse) SetStatusCode(v int32) *BatchSendMessageToGlobeResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSendMessageToGlobeResponse) SetBody(v *BatchSendMessageToGlobeResponseBody) *BatchSendMessageToGlobeResponse {
	s.Body = v
	return s
}

func (s *BatchSendMessageToGlobeResponse) Validate() error {
	return dara.Validate(s)
}
