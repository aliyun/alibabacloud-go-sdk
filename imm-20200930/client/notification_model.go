// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotification interface {
	dara.Model
	String() string
	GoString() string
	SetExtendedMessageURI(v string) *Notification
	GetExtendedMessageURI() *string
	SetMNS(v *MNS) *Notification
	GetMNS() *MNS
	SetRocketMQ(v *RocketMQ) *Notification
	GetRocketMQ() *RocketMQ
}

type Notification struct {
	// The Object Storage Service (OSS) URI of the object that stores task notifications. Task information is written to the object in the JSON format. In most cases, you can receive notifications only by using [EventBridge](https://help.aliyun.com/document_detail/161886.html), [Simple Message Queue](https://help.aliyun.com/document_detail/27412.html), or [ApsaraMQ for RocketMQ](https://help.aliyun.com/document_detail/29530.html). For tasks that have a large amount of task information, such as archive file inspection tasks and decompression tasks, you can use an OSS object to store detailed task information.
	//
	// The OSS URI follows the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// >  The object is not a messaging method. It serves only as a container for detailed task information. Task status information is sent as a message, whereas the object stores detailed task information.
	//
	// example:
	//
	// oss://test-bucket/test-object.json
	ExtendedMessageURI *string `json:"ExtendedMessageURI,omitempty" xml:"ExtendedMessageURI,omitempty"`
	// The SMQ notification settings.
	MNS *MNS `json:"MNS,omitempty" xml:"MNS,omitempty"`
	// The ApsaraMQ for RocketMQ notification settings.
	RocketMQ *RocketMQ `json:"RocketMQ,omitempty" xml:"RocketMQ,omitempty"`
}

func (s Notification) String() string {
	return dara.Prettify(s)
}

func (s Notification) GoString() string {
	return s.String()
}

func (s *Notification) GetExtendedMessageURI() *string {
	return s.ExtendedMessageURI
}

func (s *Notification) GetMNS() *MNS {
	return s.MNS
}

func (s *Notification) GetRocketMQ() *RocketMQ {
	return s.RocketMQ
}

func (s *Notification) SetExtendedMessageURI(v string) *Notification {
	s.ExtendedMessageURI = &v
	return s
}

func (s *Notification) SetMNS(v *MNS) *Notification {
	s.MNS = v
	return s
}

func (s *Notification) SetRocketMQ(v *RocketMQ) *Notification {
	s.RocketMQ = v
	return s
}

func (s *Notification) Validate() error {
	if s.MNS != nil {
		if err := s.MNS.Validate(); err != nil {
			return err
		}
	}
	if s.RocketMQ != nil {
		if err := s.RocketMQ.Validate(); err != nil {
			return err
		}
	}
	return nil
}
