// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceMirrorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteServiceMirrorResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceMirrorResponseBody
	GetRequestId() *string
}

type DeleteServiceMirrorResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Traffic mirroring is OFF for service [foo] in region [cn-shanghia], ratio [70%]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceMirrorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceMirrorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceMirrorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceMirrorResponseBody) SetMessage(v string) *DeleteServiceMirrorResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceMirrorResponseBody) SetRequestId(v string) *DeleteServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceMirrorResponseBody) Validate() error {
	return dara.Validate(s)
}
