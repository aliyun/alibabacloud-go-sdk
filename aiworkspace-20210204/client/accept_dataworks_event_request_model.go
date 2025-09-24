// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptDataworksEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *AcceptDataworksEventRequest
	GetData() map[string]interface{}
	SetMessageId(v string) *AcceptDataworksEventRequest
	GetMessageId() *string
}

type AcceptDataworksEventRequest struct {
	// The event content in the message.
	//
	// example:
	//
	// {"eventCode":"d****ct","projectId":"8***6","tenantId":4*******8,"operator":"115*****901"}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message ID. You can obtain the ID from the message received when an extension point event is triggered. For more information about the message format, see [Message formats](https://help.aliyun.com/document_detail/436911.html).
	//
	// example:
	//
	// 539306ba-*****-41a0-****-6dc81060985c
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s AcceptDataworksEventRequest) String() string {
	return dara.Prettify(s)
}

func (s AcceptDataworksEventRequest) GoString() string {
	return s.String()
}

func (s *AcceptDataworksEventRequest) GetData() map[string]interface{} {
	return s.Data
}

func (s *AcceptDataworksEventRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *AcceptDataworksEventRequest) SetData(v map[string]interface{}) *AcceptDataworksEventRequest {
	s.Data = v
	return s
}

func (s *AcceptDataworksEventRequest) SetMessageId(v string) *AcceptDataworksEventRequest {
	s.MessageId = &v
	return s
}

func (s *AcceptDataworksEventRequest) Validate() error {
	return dara.Validate(s)
}
