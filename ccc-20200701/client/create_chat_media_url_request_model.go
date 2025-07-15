// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatMediaUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateChatMediaUrlRequest
	GetInstanceId() *string
	SetMimeType(v string) *CreateChatMediaUrlRequest
	GetMimeType() *string
	SetRequestId(v string) *CreateChatMediaUrlRequest
	GetRequestId() *string
}

type CreateChatMediaUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// media id
	//
	// This parameter is required.
	//
	// example:
	//
	// jpg
	MimeType *string `json:"MimeType,omitempty" xml:"MimeType,omitempty"`
	// example:
	//
	// 9F766284-F103-4298-8EC5-19F9F9BE5522
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChatMediaUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatMediaUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateChatMediaUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateChatMediaUrlRequest) GetMimeType() *string {
	return s.MimeType
}

func (s *CreateChatMediaUrlRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatMediaUrlRequest) SetInstanceId(v string) *CreateChatMediaUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateChatMediaUrlRequest) SetMimeType(v string) *CreateChatMediaUrlRequest {
	s.MimeType = &v
	return s
}

func (s *CreateChatMediaUrlRequest) SetRequestId(v string) *CreateChatMediaUrlRequest {
	s.RequestId = &v
	return s
}

func (s *CreateChatMediaUrlRequest) Validate() error {
	return dara.Validate(s)
}
