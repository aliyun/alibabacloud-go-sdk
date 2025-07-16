// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceMirrorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateServiceMirrorResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceMirrorResponseBody
	GetRequestId() *string
}

type UpdateServiceMirrorResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Traffic mirroring is updating for service [foo] in region [cn-shanghia], ratio [70%]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceMirrorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceMirrorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceMirrorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceMirrorResponseBody) SetMessage(v string) *UpdateServiceMirrorResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceMirrorResponseBody) SetRequestId(v string) *UpdateServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceMirrorResponseBody) Validate() error {
	return dara.Validate(s)
}
