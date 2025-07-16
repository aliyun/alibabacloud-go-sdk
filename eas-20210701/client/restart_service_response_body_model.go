// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *RestartServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartServiceResponseBody
	GetRequestId() *string
}

type RestartServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Service is restarting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartServiceResponseBody) SetMessage(v string) *RestartServiceResponseBody {
	s.Message = &v
	return s
}

func (s *RestartServiceResponseBody) SetRequestId(v string) *RestartServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
