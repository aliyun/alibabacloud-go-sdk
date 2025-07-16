// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ReleaseServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReleaseServiceResponseBody
	GetRequestId() *string
}

type ReleaseServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Releasing service [foo] in region [cn-shanghai] with weight [40], service status: [Running]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReleaseServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseServiceResponseBody) SetMessage(v string) *ReleaseServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseServiceResponseBody) SetRequestId(v string) *ReleaseServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
