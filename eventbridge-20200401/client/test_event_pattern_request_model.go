// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestEventPatternRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEvent(v string) *TestEventPatternRequest
	GetEvent() *string
	SetEventPattern(v string) *TestEventPatternRequest
	GetEventPattern() *string
}

type TestEventPatternRequest struct {
	// The event.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "datacontenttype": "application/json;charset=utf-8",
	//
	//     "aliyunaccountid": "*****",
	//
	//     "aliyunpublishtime": "2023-04-****:54:57.939Z",
	//
	//     "data": {
	//
	//         "resourceEventType": "****",
	//
	//         "resourceCreateTime": "****",
	//
	//         "resourceId": "sls-code-***-debug",
	//
	//         "captureTime": "***"
	//
	//     },
	//
	//     "aliyunoriginalaccountid": "****",
	//
	//     "specversion": "1.0",
	//
	//     "aliyuneventbusname": "****",
	//
	//     "id": "295e6bd2-bb72-4f70-****-204a0680ee41",
	//
	//     "source": "acs.sls",
	//
	//     "time": "2023-04-***:37:56Z",
	//
	//     "aliyunregionid": "cn-***",
	//
	//     "type": "sls:Config:****"
	//
	// }
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The event pattern.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"key1": "value1"}
	EventPattern *string `json:"EventPattern,omitempty" xml:"EventPattern,omitempty"`
}

func (s TestEventPatternRequest) String() string {
	return dara.Prettify(s)
}

func (s TestEventPatternRequest) GoString() string {
	return s.String()
}

func (s *TestEventPatternRequest) GetEvent() *string {
	return s.Event
}

func (s *TestEventPatternRequest) GetEventPattern() *string {
	return s.EventPattern
}

func (s *TestEventPatternRequest) SetEvent(v string) *TestEventPatternRequest {
	s.Event = &v
	return s
}

func (s *TestEventPatternRequest) SetEventPattern(v string) *TestEventPatternRequest {
	s.EventPattern = &v
	return s
}

func (s *TestEventPatternRequest) Validate() error {
	return dara.Validate(s)
}
