// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceMirrorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateServiceMirrorResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceMirrorResponseBody
	GetRequestId() *string
}

type CreateServiceMirrorResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Traffic mirroring is ON for service [foo] in region [cn-shanghia], ratio [70%]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceMirrorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceMirrorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceMirrorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceMirrorResponseBody) SetMessage(v string) *CreateServiceMirrorResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceMirrorResponseBody) SetRequestId(v string) *CreateServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceMirrorResponseBody) Validate() error {
	return dara.Validate(s)
}
