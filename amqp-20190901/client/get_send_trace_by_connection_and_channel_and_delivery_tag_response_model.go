// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendTraceByConnectionAndChannelAndDeliveryTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse
	GetStatusCode() *int32
	SetBody(v *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse
	GetBody() *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody
}

type GetSendTraceByConnectionAndChannelAndDeliveryTagResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSendTraceByConnectionAndChannelAndDeliveryTagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByConnectionAndChannelAndDeliveryTagResponse) GoString() string {
	return s.String()
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse) GetBody() *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody {
	return s.Body
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse) SetHeaders(v map[string]*string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse {
	s.Headers = v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse) SetStatusCode(v int32) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse) SetBody(v *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse {
	s.Body = v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
